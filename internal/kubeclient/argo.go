package kubeclient

import (
	"context"
	"fmt"
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
