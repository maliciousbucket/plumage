package kubeclient

import (
	"context"
	"errors"
	"fmt"
	v1 "k8s.io/api/apps/v1"
	v2 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/klog"
	"log"
	"sync"
	"time"
)

func (k *k8sClient) WatchDeployment(ctx context.Context, ns, name string) error {
	dep, err := k.kubeClient.AppsV1().Deployments(ns).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return err
	}
	if dep == nil {
		return fmt.Errorf("deployment not found")
	}

	for _, cond := range dep.Status.Conditions {
		if cond.Type == v1.DeploymentAvailable {
			log.Printf("deployment %s is available", dep.Name)
			return nil
		}
	}

	return k.waitDeploymentReady(ctx, ns, name)
}

func (k *k8sClient) WatchAppDeployment(ctx context.Context, ns string, services []string) error {
	if len(services) == 0 {
		return fmt.Errorf("no services found")
	}
	var watchErr error

	errChan := make(chan error)
	resChan := make(chan string)

	var wg sync.WaitGroup

	go func() {
		wg.Wait()
		close(errChan)
		close(resChan)
	}()

	expected := len(services)
	completed := 0

	for _, service := range services {
		wg.Add(1)
		go func(service string) {
			defer wg.Done()
			select {
			case <-ctx.Done():
				errChan <- fmt.Errorf("context cancelled while watching service %s", service)
				return
			default:
				err := k.WaitAppPods(ctx, ns, service, time.Duration(2*time.Minute))
				if err != nil {
					errChan <- err
				} else {
					resChan <- service
				}
			}

		}(service)
	}
	for completed < expected {
		select {
		case err := <-errChan:
			if err != nil {
				watchErr = errors.Join(watchErr, err)
			}
			completed++
		case res := <-resChan:
			log.Printf("Pods for service %s/%s are ready", ns, res)
			completed++
		case <-ctx.Done():
			fmt.Printf("context cancelled, stopping watch for deployment in %s", ns)
			return ctx.Err()

		}
	}
	return watchErr
}

func (k *k8sClient) createDeploymentWatcher(ctx context.Context, ns, name string) (watch.Interface, error) {
	depLabels := map[string]string{
		"app": name,
	}

	labelSelector, err := labels.ValidatedSelectorFromSet(depLabels)
	if err != nil {
		return nil, err
	}

	opts := metav1.ListOptions{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
		LabelSelector: labelSelector.String(),
	}
	return k.kubeClient.AppsV1().Deployments(ns).Watch(ctx, opts)
}

func (k *k8sClient) WaitPodInstanceRunning(ctx context.Context, ns, name string) error {
	watcher, err := k.createPodInstanceWatcher(ctx, ns, name)
	if err != nil {
		return err
	}
	return k.waitPodRunning(ctx, watcher)
}

func (k *k8sClient) createPodInstanceWatcher(ctx context.Context, ns, name string) (watch.Interface, error) {
	podsLabels := map[string]string{
		"app.kubernetes.io/instance": name,
	}

	//labelSelector := fmt.Sprintf("app.kubernetes.io/instance=%s", name)
	labelSelector, err := labels.ValidatedSelectorFromSet(podsLabels)
	if err != nil {
		return nil, err
	}

	opts := metav1.ListOptions{
		TypeMeta:      metav1.TypeMeta{},
		LabelSelector: labelSelector.String(),
	}
	return k.kubeClient.CoreV1().Pods(ns).Watch(ctx, opts)
}

func (k *k8sClient) WaitPodNameRunning(ctx context.Context, ns, name string) error {
	watcher, err := k.createPodNameWatcher(ctx, ns, name)
	if err != nil {
		return err
	}
	return k.waitPodRunning(ctx, watcher)
}

func (k *k8sClient) createPodNameWatcher(ctx context.Context, ns, name string) (watch.Interface, error) {
	//fieldSelector := fmt.Sprintf("metadata.name=%s", name)
	podLabels := map[string]string{
		"app.kubernetes.io/name": name,
	}
	labelSelector, err := labels.ValidatedSelectorFromSet(podLabels)
	if err != nil {
		return nil, err
	}

	opts := metav1.ListOptions{
		TypeMeta:      metav1.TypeMeta{},
		LabelSelector: labelSelector.String(),
		//FieldSelector: fieldSelecto,
	}
	return k.kubeClient.CoreV1().Pods(ns).Watch(ctx, opts)
}

func (k *k8sClient) createPodLabelWatcher(ctx context.Context, ns, name string) (watch.Interface, error) {
	podLabels := map[string]string{
		"app": name,
	}

	labelSelector, err := labels.ValidatedSelectorFromSet(podLabels)
	if err != nil {
		return nil, fmt.Errorf("failed to parse label selector: %v", err)
	}
	opts := metav1.ListOptions{
		TypeMeta:      metav1.TypeMeta{},
		LabelSelector: labelSelector.String(),
	}
	return k.kubeClient.CoreV1().Pods(ns).Watch(ctx, opts)
}

func (k *k8sClient) waitPodRunning(ctx context.Context, watcher watch.Interface) error {

	defer watcher.Stop()

	for {
		select {
		case event := <-watcher.ResultChan():
			pod, ok := event.Object.(*v2.Pod)
			if !ok {
				return fmt.Errorf("unexpected type: %v", event.Object)
			}
			if pod == nil {
				err := fmt.Errorf("failed to get the pod")
				log.Fatal(err)
				return err
			}

			if pod.Status.Phase == v2.PodRunning {
				log.Printf("pod %s is running\n", pod.Name)
				return nil
			}

			if pod.Status.Phase == v2.PodPending {
				log.Printf("pod %s is pending\n", pod.Name)
			}

			if pod.Status.Phase == v2.PodSucceeded {
				log.Printf("pod %s is succeeded\n", pod.Name)
				return nil
			}

			if pod.Status.Phase == "Failed" {
				return fmt.Errorf("pod %s failed", pod.Name)
			}
		case <-ctx.Done():
			log.Printf("Exiting from wait pod runnning becuase context stopped\n")
			return nil
		}
	}
}

func (k *k8sClient) waitDeploymentReady(ctx context.Context, ns, name string) error {
	watcher, err := k.createDeploymentWatcher(ctx, ns, name)
	if err != nil {
		return err
	}

	defer watcher.Stop()

	for {
		select {
		case event := <-watcher.ResultChan():
			deployment, ok := event.Object.(*v1.Deployment)
			if !ok {
				return fmt.Errorf("unexpected type of object: %v", event.Object)
			}
			if deployment.Status.String() == "Ready" {

				klog.Info("deployment %s is ready\n", deployment.Name)
				return nil
			}

			for _, cond := range deployment.Status.Conditions {
				if cond.Type == v1.DeploymentAvailable {
					klog.Info("deployment %s is available\n", deployment.Name)
					return nil
				}

				if cond.Type == v1.DeploymentReplicaFailure {
					klog.Errorf("deployment %s is failed\n", deployment.Name)
					return fmt.Errorf("deployment %s replicas have failed", deployment.Name)
				}

			}

		case <-ctx.Done():
			log.Printf("Exiting from wait deployment ready becuase context stopped\n")
		}
	}
}

func (k *k8sClient) WaitAppPods(ctx context.Context, ns, name string, timeout time.Duration) error {
	return k.waitAppPods(ctx, ns, name, timeout)
}

func (k *k8sClient) waitAppPods(ctx context.Context, ns, name string, timeout time.Duration) error {
	labelSelector, err := labels.ValidatedSelectorFromSet(map[string]string{"app.kubernetes.io/name": name})
	if err != nil {
		return fmt.Errorf("failed to parse label selector: %v", err)
	}

	pods, err := k.kubeClient.CoreV1().Pods(ns).List(ctx, metav1.ListOptions{
		TypeMeta:      metav1.TypeMeta{},
		LabelSelector: labelSelector.String(),
	})
	if err != nil {
		return fmt.Errorf("failed to list pods for %s/%s: %v", ns, name, err)
	}
	if len(pods.Items) == 0 {
		return fmt.Errorf("no pods found for %s/%s", ns, name)
	}
	expectedPods := len(pods.Items)
	fmt.Printf("Expected pods: %d\n", expectedPods)
	for _, pod := range pods.Items {
		fmt.Printf("Pod: %s\n", pod.Name)
	}

	watcher, err := k.createPodLabelWatcher(ctx, ns, name)
	if err != nil {
		return err
	}
	errChan := make(chan error)
	statusChan := make(chan int)
	runningPods := make(map[string]bool)
	defer watcher.Stop()
	go func() {
		for {
			select {
			case <-ctx.Done():
				errChan <- ctx.Err()
				return
			case event := <-watcher.ResultChan():
				pod, ok := event.Object.(*v2.Pod)
				if !ok {
					errChan <- fmt.Errorf("unexpected type: %v", event.Object)
					return
				}
				if pod.Status.Phase == v2.PodRunning && !runningPods[pod.Name] {
					runningPods[pod.Name] = true
					statusChan <- len(runningPods)
				}

				if pod.Status.Phase == v2.PodFailed {
					errChan <- fmt.Errorf("pod %s failed", pod.Name)
					return
				}

				if pod.Status.Phase == v2.PodSucceeded {

				}
			}
		}
	}()
	go func() {
		podList, podErr := k.kubeClient.CoreV1().Pods(ns).List(ctx, metav1.ListOptions{
			LabelSelector: labelSelector.String(),
		})
		if podErr != nil {
			errChan <- podErr
		}
		for _, pod := range podList.Items {
			if pod.Status.Phase == v2.PodRunning && !runningPods[pod.Name] {
				runningPods[pod.Name] = true
			}
		}
		statusChan <- len(runningPods)
	}()

	for {
		select {
		case podsRunning := <-statusChan:
			if podsRunning == expectedPods {
				log.Printf("all pods (%d) for %s/%s are running", podsRunning, ns, name)
				return nil
			} else {
				log.Printf("waiting for all pods for %s%s to be running. %d/%d", ns, name, podsRunning, expectedPods)
			}
		case podErr := <-errChan:
			if podErr != nil {
				return podErr
			}
		case <-time.After(timeout):
			return fmt.Errorf("timed out waiting for pods for %s%s to be running", ns, name)
		case <-ctx.Done():
			klog.Info("Exiting from wait deployment pods because context stopped\n")
		}

	}

}

//func (k *k8sClient) inform(ctx context.Context, ns, name string) error {
//	factory := informers.NewSharedInformerFactoryWithOptions(k.kubeClient, 10*time.Second, informers.WithNamespace(ns))
//
//	informer := factory.Apps().V1().Deployments().Informer()
//
//	defer runtime.HandleCrash()
//	stopper := make(chan struct{})
//	defer close(stopper)
//
//	go factory.Start(stopper)
//
//	if !cache.WaitForCacheSync(stopper, informer.HasSynced) {
//
//	}
//
//	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
//		AddFunc: func(obj interface{}) {
//
//		},
//		UpdateFunc: func(old, new interface{}) {},
//		DeleteFunc: func(obj interface{}) {},
//	})
//
//}
