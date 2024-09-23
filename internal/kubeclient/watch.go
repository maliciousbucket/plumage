package kubeclient

import (
	"context"
	"fmt"
	v1 "k8s.io/api/apps/v1"
	v2 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/klog"
	"log"
)

func (k *k8sClient) WatchDeployment(ctx context.Context, ns, name string) error {
	return k.waitDeploymentReady(ctx, ns, name)
}

func (k *k8sClient) createDeploymentWatcher(ctx context.Context, ns, name string) (watch.Interface, error) {
	labelSelector := fmt.Sprintf("app=%s", name)

	opts := metav1.ListOptions{
		TypeMeta:      metav1.TypeMeta{},
		LabelSelector: labelSelector,
		FieldSelector: "",
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
	labelSelector := fmt.Sprintf("app.kubernetes.io/instance=%s", name)

	opts := metav1.ListOptions{
		TypeMeta:      metav1.TypeMeta{},
		LabelSelector: labelSelector,
		FieldSelector: "",
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
	filedSelector := fmt.Sprintf("metadata.name=%s", name)
	opts := metav1.ListOptions{
		TypeMeta:      metav1.TypeMeta{},
		LabelSelector: "",
		FieldSelector: filedSelector,
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

			if pod.Status.Phase == "Running" {
				fmt.Printf("pod %s is running\n", pod.Name)
				return nil
			}

			if pod.Status.Phase == "Pending" {
				fmt.Printf("pod %s is pending\n", pod.Name)
			}

			if pod.Status.Phase == "Failed" {
				return fmt.Errorf("pod %s failed", pod.Name)
			}
		case <-ctx.Done():
			fmt.Printf("Exiting from wait pod runnning becuase context stopped\n")
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
			fmt.Printf("Deployment Status : %s", deployment.Status.String())
			if deployment.Status.String() == "Ready" {
				fmt.Printf("deployment %s is ready\n", deployment.Name)
				klog.Info("deployment %s is ready\n", deployment.Name)
			}

			for _, cond := range deployment.Status.Conditions {

				klog.Info("Type: %s Status: %s Condition: %s\n", cond.Type, cond.Status, cond.Reason)
			}

			if deployment.Status.ReadyReplicas == deployment.Status.Replicas {
				fmt.Printf("Deployment %s is ready\n", deployment.Name)
				return nil
			}

		case <-ctx.Done():
			fmt.Printf("Context done: %s\n", name)
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
