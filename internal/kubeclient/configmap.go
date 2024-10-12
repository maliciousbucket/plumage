package kubeclient

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/util/retry"
)

func (k *k8sClient) listConfigMaps(ctx context.Context, namespace string) (*v1.ConfigMapList, error) {
	configMaps, err := k.kubeClient.CoreV1().ConfigMaps(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return configMaps, nil
}

func (k *k8sClient) getConfigMapByName(ctx context.Context, ns, name string) (*v1.ConfigMap, error) {
	configMap, err := k.kubeClient.CoreV1().ConfigMaps(ns).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return configMap, nil
}

func (k *k8sClient) CreateGalahArgoAccount(ctx context.Context, ns string) error {
	return k.addGalahAccount(ctx, ns)
}

func (k *k8sClient) addGalahAccount(ctx context.Context, ns string) error {
	configMapClient := k.kubeClient.CoreV1().ConfigMaps(ns)

	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		result, getErr := configMapClient.Get(ctx, "argocd-cm", metav1.GetOptions{})
		if getErr != nil {
			return fmt.Errorf("error getting argocd-cm ConfigMap: %v", getErr)
		}
		result.Data = map[string]string{
			"accounts.galah":         "apiKey, login",
			"kustomize.buildOptions": "--enable-helm",
		}
		_, updateErr := configMapClient.Update(ctx, result, metav1.UpdateOptions{})
		return updateErr
	})

	if retryErr != nil {
		return fmt.Errorf("error updating argocd-cm ConfigMap: %v", retryErr)
	}

	rbacRetryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		result, getErr := configMapClient.Get(ctx, "argocd-rbac-cm", metav1.GetOptions{})
		if getErr != nil {
			return fmt.Errorf("error getting argocd-rbac-cm ConfigMap: %v", getErr)
		}
		result.Data = map[string]string{
			"policy.csv": `g, galah, role:admin`,
		}
		_, updateErr := configMapClient.Update(ctx, result, metav1.UpdateOptions{})
		return updateErr
	})
	if rbacRetryErr != nil {
		return fmt.Errorf("error updating argocd-rbac-cm ConfigMap: %v", retryErr)
	}
	return nil

}

func (k *k8sClient) setKustomizeBuildOptions(ctx context.Context, ns string) error {
	configMapClient := k.kubeClient.CoreV1().ConfigMaps(ns)

	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		result, getErr := configMapClient.Get(ctx, "argocd-cm", metav1.GetOptions{})
		if getErr != nil {
			return fmt.Errorf("error getting argocd-cm ConfigMap: %v", getErr)
		}
		result.Data["kustomize.buildOptions"] = "--enable-helm"

		_, updateErr := configMapClient.Update(ctx, result, metav1.UpdateOptions{})
		return updateErr
	})
	if retryErr != nil {
		return fmt.Errorf("error updating argocd-cm ConfigMap: %v", retryErr)
	}
	return nil
}
