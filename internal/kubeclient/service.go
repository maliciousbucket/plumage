package kubeclient

import (
	"context"
	"errors"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/util/retry"
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

func (k *k8sClient) ExposeService(ctx context.Context, ns string, name string, port int, nodePort int) error {
	return k.exposeService(ctx, ns, name, port, nodePort)
}

func (k *k8sClient) exposeService(ctx context.Context, ns string, name string, port int, nodePort int) error {

	serviceClient := k.kubeClient.CoreV1().Services(ns)
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		service, err := serviceClient.Get(ctx, name, metav1.GetOptions{})
		if err != nil {
			return err
		}
		if service == nil {
			return fmt.Errorf("service %s/%s not found", ns, name)
		}

		foundNodePort := false
		for i, p := range service.Spec.Ports {
			if p.Port == int32(port) {
				service.Spec.Ports[i].NodePort = int32(nodePort)
				//service.Spec.Ports[i].Name = "nodeport"
				foundNodePort = true
				break
			}
		}
		if !foundNodePort {
			exposedPort := v1.ServicePort{
				NodePort: int32(nodePort),
				Name:     "nodeport",
				Port:     int32(port),
				Protocol: v1.ProtocolTCP,
			}
			service.Spec.Ports = append(service.Spec.Ports, exposedPort)
		}

		foundDisruptorPort := false
		for _, p := range service.Spec.Ports {
			if p.Port == 8000 {
				foundDisruptorPort = true
				break
			}
		}
		if !foundDisruptorPort {
			disruptorPort := v1.ServicePort{
				Port:       int32(8000),
				TargetPort: intstr.FromInt32(int32(port)),
				Name:       "disruptor",
				Protocol:   v1.ProtocolTCP,
			}
			service.Spec.Ports = append(service.Spec.Ports, disruptorPort)
		}

		_, updateErr := serviceClient.Update(ctx, service, metav1.UpdateOptions{})
		return updateErr
	})
	if retryErr != nil {
		return fmt.Errorf("could not expose service %s/%s to node port: %v", ns, name, retryErr)
	}

	log.Printf("Exposed %s/%s on Port %d - NodePort: %d, Disruptor Port: 8000", ns, name, port, nodePort)
	return nil
}
