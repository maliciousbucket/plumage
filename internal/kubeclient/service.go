package kubeclient

import (
	"context"
	"errors"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sync"
)

func (k *k8sClient) getService(ctx context.Context, ns string, name string) (*v1.Service, error) {
	res, err := k.kubeClient.CoreV1().Services(ns).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	//a, err := k.kubeClient.AppsV1().Deployments(ns).Watch(ctx, metav1.ListOptions{})
	//
	//informer := informers.NewSharedInformerFactoryWithOptions(k.kubeClient, 10*time.Second, informers.WithNamespace(ns))
	//
	//depInfromer := informer.Apps().V1().Deployments().Informer()

	return res, nil
}

func (k *k8sClient) getServicePods(ctx context.Context, ns string, name string) (*v1.PodList, error) {
	labelSelector := fmt.Sprintf("app.kubernetes.io/name=%s", name)
	res, err := k.kubeClient.CoreV1().Pods(ns).List(ctx, metav1.ListOptions{
		LabelSelector: labelSelector,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (k *k8sClient) WaitServicePods(ctx context.Context, ns string, name string) error {
	pods, err := k.getServicePods(ctx, ns, name)
	if err != nil {
		return err
	}
	if pods == nil || len(pods.Items) == 0 {
		return fmt.Errorf("no pods found for service %s/%s", ns, name)
	}

	return k.waitServicePods(ctx, ns, pods)
}

func (k *k8sClient) waitServicePods(ctx context.Context, ns string, pods *v1.PodList) error {

	var instances []string

	for _, pod := range pods.Items {
		//instance, ok := pod.Labels["app.kubernetes.io/instance"]
		//if !ok {
		//	return fmt.Errorf("pod %s/%s does not have instance label", pod.Namespace, pod.Name)
		//}
		instance := pod.Name
		instances = append(instances, instance)
	}
	errChan := make(chan error, len(pods.Items))
	var wg sync.WaitGroup
	for _, instance := range instances {
		wg.Add(1)
		go func(instance string) {
			defer wg.Done()
			//err := k.waitPodRunning(ctx, ns, instance)
			err := k.WaitPodNameRunning(ctx, ns, instance)
			select {
			case errChan <- err:
			case <-ctx.Done():
				errChan <- ctx.Err()
			}
		}(instance)
	}
	wg.Wait()
	close(errChan)
	var podErr error
	for errs := range errChan {
		if errs != nil {
			podErr = errors.Join(podErr, errs)
		}
	}
	return podErr
}
func (k *k8sClient) getRelatedPods(ctx context.Context, ns string, name string) (*v1.PodList, error) {
	labelSelector := fmt.Sprintf("app.kubernetes.io/part-of=%s", name)
	res, err := k.kubeClient.CoreV1().Pods(ns).List(ctx, metav1.ListOptions{
		LabelSelector: labelSelector,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (k *k8sClient) GetServiceAddress(ctx context.Context, ns string, name string) (string, error) {
	service, err := k.getService(ctx, ns, name)
	if err != nil {
		return "", err
	}
	if service == nil {
		return "", fmt.Errorf("service %s/%s not found", ns, name)
	}

	if service.Status.LoadBalancer.Ingress != nil {
		if len(service.Status.LoadBalancer.Ingress[0].Hostname) != 0 {
			return service.Status.LoadBalancer.Ingress[0].Hostname, nil
		}
	}

	if len(service.Spec.ClusterIPs) != 0 {
		return service.Spec.ClusterIPs[0], nil
	}
	return "", fmt.Errorf("no LoadBalancer or ClusterIP found for service %s/%s", ns, name)

}
