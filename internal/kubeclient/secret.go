package kubeclient

import (
	"context"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const initialPasswordSecretName = "argocd-initial-admin-secret"

func (k *k8sClient) GetArgoPassword(ctx context.Context, ns string) (string, error) {
	return k.getArgoCDInitialPassword(ctx, ns)
}

func (k *k8sClient) getArgoCDInitialPassword(ctx context.Context, ns string) (string, error) {
	secretsClient := k.kubeClient.CoreV1().Secrets(ns)
	pass, err := secretsClient.Get(ctx, initialPasswordSecretName, v1.GetOptions{})
	if err != nil {
		return "", fmt.Errorf("error getting ArgoCD initial admin secret: %w", err)
	}
	result := pass.Data["password"]
	if len(result) == 0 {
		return "", fmt.Errorf("error getting ArgoCD initial admin secret: password not found in secret")
	}
	return string(result), nil

}
