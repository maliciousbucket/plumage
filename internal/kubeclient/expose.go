package kubeclient

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/util/retry"
	"log"
)

var (
	externalLbAnnotation          = "service.beta.kubernetes.io/aws-load-balancer-type"
	externalLbTypeAnnotation      = "service.beta.kubernetes.io/aws-load-balancer-nlb-target-type"
	externalLbSchemeAnnotation    = "service.beta.kubernetes.io/aws-load-balancer-scheme"
	ExposeServiceTypeLoadBalancer = ExposeServiceType("LoadBalancer")
	ExposeServiceTypeNodePort     = ExposeServiceType("NodePort")
)

//service.beta.kubernetes.io/aws-load-balancer-type: "external"
//service.beta.kubernetes.io/aws-load-balancer-nlb-target-type: "ip"
//service.beta.kubernetes.io/aws-load-balancer-scheme: internet-facing

type ExposeServiceType string

func (k *k8sClient) ExposeService(ctx context.Context, ns string, name string, port int, nodePort int, serviceType ExposeServiceType) error {
	switch serviceType {
	case ExposeServiceTypeLoadBalancer:
		info, err := k.exposeServiceAsLoadBalancer(ctx, ns, name)
		if err != nil {
			return err
		}
		log.Printf("\n %+v \n", info)
		return nil
	case ExposeServiceTypeNodePort:
		return k.exposeServiceAsNodePort(ctx, ns, name, port, nodePort)
	default:
		return fmt.Errorf("unknown expose service type: %s", serviceType)

	}
}

func (k *k8sClient) exposeServiceAsNodePort(ctx context.Context, ns string, name string, port int, nodePort int) error {

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
		if service.Spec.Type != v1.ServiceTypeNodePort && service.Spec.Type != v1.ServiceTypeLoadBalancer {
			service.Spec.Type = v1.ServiceTypeNodePort
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

func (k *k8sClient) exposeServiceAsLoadBalancer(ctx context.Context, ns string, name string) (*LoadBalancerInfo, error) {
	serviceClient := k.kubeClient.CoreV1().Services(ns)
	foundSvc, err := serviceClient.Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	if foundSvc == nil {
		return nil, fmt.Errorf("service %s/%s not found", ns, name)
	}

	deleteRetryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		deleteErr := serviceClient.Delete(ctx, name, metav1.DeleteOptions{})
		if deleteErr != nil {
			return fmt.Errorf("could not delete service %s/%s: %v", ns, name, deleteErr)
		}
		return nil
	})

	if deleteRetryErr != nil {
		return nil, fmt.Errorf("error deleting service %s/%s: %v", ns, name, deleteRetryErr)
	}

	newService := foundSvc.DeepCopy()
	newService.Spec.Type = v1.ServiceTypeLoadBalancer
	newService.Annotations[externalLbAnnotation] = "external"
	newService.Annotations[externalLbTypeAnnotation] = "ip"
	newService.Annotations[externalLbSchemeAnnotation] = "internet-facing"

	createErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		_, err = serviceClient.Create(ctx, newService, metav1.CreateOptions{})
		if err != nil {
			return err
		}
		return nil
	})
	if createErr != nil {
		return nil, fmt.Errorf("could not create service %s/%s: %v", ns, name, createErr)
	}
	service, err := serviceClient.Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	if service == nil {
		return nil, fmt.Errorf("service %s/%s not found after update", ns, name)
	}
	log.Printf("Exposed %s/%s as LoadBalancer", ns, name)
	return &LoadBalancerInfo{
		Name:        service.Name,
		Namespace:   service.Namespace,
		ExternalIPs: service.Spec.ExternalIPs,
		Ports:       service.Spec.Ports,
	}, nil
}
