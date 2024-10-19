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
	"log"
	"sync"
	"time"
)

func (k *k8sClient) WatchDeployment(ctx context.Context, ns, name string, meta bool) error {
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

	return k.waitDeploymentReady(ctx, ns, name, meta)
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
				err := k.WaitAppPods(ctx, ns, service, 1, time.Duration(2*time.Minute))
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

func (k *k8sClient) createDeploymentMetaNameWatcher(ctx context.Context, ns, name string) (watch.Interface, error) {
	fieldSelector := fmt.Sprintf("metadata.name=%s", name)
	opts := metav1.ListOptions{
		FieldSelector: fieldSelector,
	}
	return k.kubeClient.AppsV1().Deployments(ns).Watch(ctx, opts)
}

func (k *k8sClient) waitDeploymentReady(ctx context.Context, ns, name string, meta bool) error {

	var watcher watch.Interface
	if meta {
		nameWatcher, err := k.createDeploymentMetaNameWatcher(ctx, ns, name)
		if err != nil {
			return err
		}
		watcher = nameWatcher
	} else {
		nameWatcher, err := k.createDeploymentWatcher(ctx, ns, name)
		if err != nil {
			return err
		}
		watcher = nameWatcher
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

				log.Printf("deployment %s is ready\n", deployment.Name)
				return nil
			}

			for _, cond := range deployment.Status.Conditions {
				if cond.Type == v1.DeploymentAvailable {
					log.Printf("deployment %s is available\n", deployment.Name)
					return nil
				}

				if cond.Type == v1.DeploymentReplicaFailure {
					log.Printf("deployment %s is failed\n", deployment.Name)
					return fmt.Errorf("deployment %s replicas have failed", deployment.Name)
				}

			}

		case <-ctx.Done():
			log.Printf("Exiting from wait deployment ready becuase context stopped\n")
			return ctx.Err()
		}
	}
}
func (k *k8sClient) WaitPodInstanceRunning(ctx context.Context, ns, name string) error {
	//"app.kubernetes.io/instance"
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

func (k *k8sClient) WaitPodNameLabelRunning(ctx context.Context, ns, name string) error {
	//app.kubernetes.io/name
	watcher, err := k.createPodNameLabelWatcher(ctx, ns, name)
	if err != nil {
		return err
	}
	return k.waitPodRunning(ctx, watcher)
}

func (k *k8sClient) WaitPodNameRunning(ctx context.Context, ns, name string) error {
	watcher, err := k.createPodFieldWatcher(ctx, ns, name)
	if err != nil {
		return err
	}
	return k.waitPodRunning(ctx, watcher)
}

func (k *k8sClient) createPodNameLabelWatcher(ctx context.Context, ns, name string) (watch.Interface, error) {
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

func (k *k8sClient) createPodFieldWatcher(ctx context.Context, ns, name string) (watch.Interface, error) {
	fieldSelector := fmt.Sprintf("metadata.name=%s", name)
	opts := metav1.ListOptions{
		FieldSelector: fieldSelector,
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

			return ctx.Err()
		}
	}
}

func (k *k8sClient) WaitAppPods(ctx context.Context, ns, name string, expected int, timeout time.Duration) error {
	return k.waitAppPods(ctx, ns, name, expected, timeout)
}

func (k *k8sClient) waitAppPods(ctx context.Context, ns, name string, expected int, timeout time.Duration) error {
	tctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	//fieldSelector := fmt.Sprintf("metadata.name=%s", name)
	labelSelector, err := labels.ValidatedSelectorFromSet(map[string]string{"app.kubernetes.io/name": name})
	if err != nil {
		return fmt.Errorf("failed to parse label selector: %v", err)
	}

	pods, err := k.kubeClient.CoreV1().Pods(ns).List(tctx, metav1.ListOptions{
		LabelSelector: labelSelector.String(),
	})

	if err != nil {
		return err
	}

	if len(pods.Items) == 0 {
		nameSelector, err := labels.ValidatedSelectorFromSet(map[string]string{"app": name})
		if err != nil {
			return fmt.Errorf("failed to parse label selector: %v", err)
		}
		namePods, err := k.kubeClient.CoreV1().Pods(ns).List(tctx, metav1.ListOptions{
			LabelSelector: nameSelector.String(),
		})
		if err != nil {
			return err
		}
		if len(namePods.Items) == 0 {
			return fmt.Errorf("no pods found for %s/%s\n", ns, name)
		}
		pods = namePods

	}

	podsRunning := 0
	for _, pod := range pods.Items {
		if pod.Status.Phase == v2.PodRunning {
			podsRunning++
		}
	}
	if podsRunning == expected {
		return nil
	}
	if podsRunning <= 0 {
		return nil
	}

	newExpected := 0
	for _, pod := range pods.Items {
		if pod.Status.Phase == v2.PodPending || pod.Status.Phase == v2.PodUnknown {
			newExpected++
		}
	}

	watcher, err := k.createPodInstanceWatcher(tctx, ns, name)
	if err != nil {
		return err
	}
	defer watcher.Stop()
	for event := range watcher.ResultChan() {
		pod, ok := event.Object.(*v2.Pod)
		if !ok {
			return fmt.Errorf("unexpected type: %v", event.Object)
		}
		switch pod.Status.Phase {
		case v2.PodRunning, v2.PodSucceeded:
			podsRunning++
		case v2.PodFailed:
			return fmt.Errorf("pod %s failed", pod.Name)
		case v2.PodPending, v2.PodUnknown:
			log.Printf("pod %s is pending or in unknown state\n", pod.Name)
		}
		if podsRunning >= expected {
			return nil
		}
		if podsRunning >= newExpected {
			return nil
		}
		select {
		case <-tctx.Done():
			return fmt.Errorf("timed out waiting for pod %s to be running: %v", pod.Name, tctx.Err())
		default:
		}

	}

	log.Printf("Waiting for %d/%d pods running\n", expected-podsRunning, len(pods.Items))
	return fmt.Errorf("timedout waiting for %d pods running", expected-podsRunning)

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
