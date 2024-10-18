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
	externalLbAnnotation       = "service.beta.kubernetes.io/aws-load-balancer-type"
	externalLbTypeAnnotation   = "service.beta.kubernetes.io/aws-load-balancer-nlb-target-type"
	externalLbSchemeAnnotation = "service.beta.kubernetes.io/aws-load-balancer-scheme"
)

//service.beta.kubernetes.io/aws-load-balancer-type: "external"
//service.beta.kubernetes.io/aws-load-balancer-nlb-target-type: "ip"
//service.beta.kubernetes.io/aws-load-balancer-scheme: internet-facing

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
