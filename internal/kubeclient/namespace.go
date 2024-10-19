package kubeclient

import (
	"context"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	namespaceKind = "Namespace"
	appsV1        = "apps/v1"
)

func (k *k8sClient) CreateNamespace(ctx context.Context, ns string) (*NameSpaceInfo, error) {
	namespace, err := k.kubeClient.CoreV1().Namespaces().Get(ctx, ns, metav1.GetOptions{})
	if err != nil {
		if !errors.IsNotFound(err) {
			return nil, err
		}

	}
	if namespace != nil && namespace.Name != "" {
		return &NameSpaceInfo{
			Name:        namespace.Name,
			Namespace:   namespace.Namespace,
			Labels:      namespace.Labels,
			Annotations: namespace.Annotations,
		}, nil
	}
	res, err := k.kubeClient.CoreV1().Namespaces().Create(ctx, &v1.Namespace{
		TypeMeta: metav1.TypeMeta{
			Kind:       namespaceKind,
			APIVersion: appsV1,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: ns,
		},
	}, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return &NameSpaceInfo{
		Name:        res.Name,
		Namespace:   res.Namespace,
		Labels:      res.Labels,
		Annotations: res.Annotations,
	}, nil
}

type NameSpaceInfo struct {
	Name        string
	Namespace   string
	Labels      map[string]string
	Annotations map[string]string
}

func (k *k8sClient) listNamespaces(ctx context.Context) (*v1.NamespaceList, error) {
	namespaces, err := k.kubeClient.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return namespaces, nil
}
