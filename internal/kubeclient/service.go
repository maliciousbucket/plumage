package kubeclient

import (
	"context"
	"errors"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"sync"
)

func (k *k8sClient) CheckServiceExists(ctx context.Context, namespace string, name string) (bool, error) {
	res, err := k.getService(ctx, namespace, name)
	if err != nil {
		return false, err
	}
	if res != nil {
		return true, nil
	}
	return false, nil
}

func (k *k8sClient) getService(ctx context.Context, ns string, name string) (*v1.Service, error) {
	log.Println("getting service", ns, name)
	res, err := k.kubeClient.CoreV1().Services(ns).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

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
	if len(res.Items) == 0 {
		return nil, fmt.Errorf("no pods found for service %s/%s", ns, name)
	}
	return res, nil
}

func (k *k8sClient) WaitServicePods(ctx context.Context, ns string, name string) error {
	pods, err := k.getServicePods(ctx, ns, name)
	log.Println(pods)
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
		instance := pod.Name
		instances = append(instances, instance)
	}
	errChan := make(chan error, len(pods.Items))
	var wg sync.WaitGroup
	for _, instance := range instances {
		wg.Add(1)
		go func(instance string) {
			defer wg.Done()
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

// GetServiceAddress TODO: Refactor to return lb ip
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
func (k *k8sClient) GetExternalAddress(ctx context.Context, ns string, name string) ([]string, error) {
	return k.getServiceExternalAddress(ctx, ns, name)
}

func (k *k8sClient) getServiceExternalAddress(ctx context.Context, ns string, name string) ([]string, error) {
	service, err := k.getService(ctx, ns, name)
	if err != nil {
		return nil, err
	}
	if service == nil {
		return nil, fmt.Errorf("service %s/%s not found", ns, name)
	}
	if service.Spec.Type != v1.ServiceTypeLoadBalancer {
		return nil, fmt.Errorf("service %s/%s is not LoadBalancer", ns, name)
	}
	if service.Spec.ExternalIPs == nil || len(service.Spec.ExternalIPs) == 0 {
		return nil, fmt.Errorf("service %s/%s has no external IPs", ns, name)
	}
	return service.Spec.ExternalIPs, nil
}

type LoadBalancerInfo struct {
	Name        string
	Namespace   string
	ExternalIPs []string
	Ports       []v1.ServicePort
}

func (k *k8sClient) GetLoadBalancersForNamespace(ctx context.Context, ns string) ([]*LoadBalancerInfo, error) {
	return k.getLoadBalancersForNamespace(ctx, ns)
}

func (k *k8sClient) getLoadBalancersForNamespace(ctx context.Context, ns string) ([]*LoadBalancerInfo, error) {
	serviceClient := k.kubeClient.CoreV1().Services(ns)
	fieldSelector := fmt.Sprintf("spec.type=LoadBalancer")
	loadBalancers, err := serviceClient.List(ctx, metav1.ListOptions{
		FieldSelector: fieldSelector,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list load balancers: %v", err)
	}

	if loadBalancers == nil || len(loadBalancers.Items) == 0 {
		return nil, fmt.Errorf("no load balancers found for namespace %s", ns)
	}
	result := make([]*LoadBalancerInfo, 0, len(loadBalancers.Items))
	for _, lb := range loadBalancers.Items {
		result = append(result, &LoadBalancerInfo{
			Name:        lb.Name,
			Namespace:   lb.Namespace,
			ExternalIPs: lb.Spec.ExternalIPs,
			Ports:       lb.Spec.Ports,
		})
	}
	return result, nil
}
