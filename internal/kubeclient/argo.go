package kubeclient

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sync"
	"time"
)

const (
	argoCDServerName = "argocd-server"
)

type ServiceInfo struct {
	Name        string
	Namespace   string
	CreatedAt   time.Time
	Labels      map[string]string
	Annotations map[string]string
	Status      string
}

func (k *k8sClient) CheckArgoExists(ctx context.Context, ns string) (*ServiceInfo, error) {
	res, err := k.getService(ctx, ns, argoCDServerName)

	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return &ServiceInfo{
		Name:        res.Name,
		Namespace:   res.Namespace,
		CreatedAt:   res.CreationTimestamp.Time,
		Labels:      res.Labels,
		Annotations: res.Annotations,
		Status:      res.Status.String(),
	}, nil
}

func (k *k8sClient) WaitAllArgoPods(ctx context.Context, ns string) error {
	pods, err := k.getRelatedPods(ctx, ns, "argocd")

	if err != nil {
		return err
	}
	if pods == nil || len(pods.Items) == 0 {
		return fmt.Errorf("no Argo CD pods found in namespace %s", ns)
	}
	return k.waitServicePods(ctx, ns, pods)
}

func (k *k8sClient) PatchArgoToLB(ctx context.Context, ns string) error {
	service, err := k.getService(ctx, ns, "argocd-helm-server")
	if err != nil {
		return err
	}
	if service == nil {
		return fmt.Errorf("service %s not found", argoCDServerName)
	}
	service.Spec.Type = v1.ServiceTypeLoadBalancer
	_, err = k.kubeClient.CoreV1().Services(ns).Update(ctx, service, metav1.UpdateOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (k *k8sClient) portForwardToArgoCD(ctx context.Context, ns string, port int, wg *sync.WaitGroup, stopCh chan struct{}) error {

	service, err := k.getService(ctx, ns, argoCDServerName)
	if err != nil {
		return err
	}
	if service == nil {
		return fmt.Errorf("argocd-server service not found")
	}
	var pfErr error
	go func(pfErr error) {
		defer wg.Done()
		err = PortForwardService(service, port, 8080, stopCh)
		if err != nil {
			pfErr = err
		}
		select {
		case <-stopCh:
			return
		case <-ctx.Done():
			return
		}
	}(pfErr)
	return pfErr

}
