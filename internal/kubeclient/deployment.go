package kubeclient

import (
	"context"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *k8sClient) ListDeployments(ctx context.Context, ns string) ([]v1.Deployment, error) {
	deployments, err := k.kubeClient.AppsV1().Deployments(ns).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return deployments.Items, nil
}
