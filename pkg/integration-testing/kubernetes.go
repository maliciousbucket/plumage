package integration_testing

import (
	"context"
	"github.com/maliciousbucket/plumage/internal/kubeclient"
	appsV1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"time"
)

type kubeClient interface {
	WatchDeployment(ctx context.Context, ns string, name string, meta bool) error
	WatchAppDeployment(ctx context.Context, ns string, services []string) error
	WaitPodInstanceRunning(ctx context.Context, ns string, name string) error
	WaitPodNameRunning(ctx context.Context, ns string, name string) error
	CreateNamespace(ctx context.Context, ns string) (*kubeclient.NameSpaceInfo, error)
	CheckArgoExists(ctx context.Context, ns string) (*kubeclient.ServiceInfo, error)
	WaitAllArgoPods(ctx context.Context, ns string) error
	WaitServicePods(ctx context.Context, ns string, name string) error
	WaitAppPods(ctx context.Context, ns, name string, expected int, timeout time.Duration) error

	GetServiceAddress(ctx context.Context, ns string, name string) (string, error)
	CreateGalahArgoAccount(ctx context.Context, ns string) error
	GetArgoPassword(ctx context.Context, ns string) (string, error)
	CheckServiceExists(ctx context.Context, ns string, name string) (bool, error)
	ExposeService(ctx context.Context, ns string, name string, port int, nodePort int, serviceType kubeclient.ExposeServiceType) error
	ListDeployments(ctx context.Context, ns string) ([]appsV1.Deployment, error)
	ListPods(ctx context.Context, namespace string) (*v1.PodList, error)
}
