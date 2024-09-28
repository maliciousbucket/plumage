package kubeclient

import (
	"context"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"os"
	"path/filepath"
	"time"
)

// Client TODO: Change to interface
type Client interface {
	WatchDeployment(ctx context.Context, ns string, name string) error
	WatchAppDeployment(ctx context.Context, ns string, services []string) error
	WaitPodInstanceRunning(ctx context.Context, ns string, name string) error
	WaitPodNameRunning(ctx context.Context, ns string, name string) error
	CreateNamespace(ctx context.Context, ns string) (*NameSpaceInfo, error)
	CheckArgoExists(ctx context.Context, ns string) (*ServiceInfo, error)
	WaitAllArgoPods(ctx context.Context, ns string) error
	PatchArgoToLB(ctx context.Context, ns string) error
	WaitServicePods(ctx context.Context, ns string, name string) error
	WaitAppPods(ctx context.Context, ns, name string, timeout time.Duration) error

	GetServiceAddress(ctx context.Context, ns string, name string) (string, error)
	CreateGalahArgoAccount(ctx context.Context, ns string) error
	GetArgoPassword(ctx context.Context, ns string) (string, error)
	CheckServiceExists(ctx context.Context, ns string, name string) (bool, error)
	ExposeService(ctx context.Context, ns string, name string, port int, nodePort int) error
}

type k8sClient struct {
	kubeClient *kubernetes.Clientset
}

func NewClient() (Client, error) {
	kubeClient, err := newClientset()
	if err != nil {
		return nil, err
	}
	client := k8sClient{
		kubeClient: kubeClient,
	}
	return &client, nil
}

func newClientset() (*kubernetes.Clientset, error) {
	kubeCfg := os.Getenv("KUBECONFIG")

	if kubeCfg == "" {
		home := homedir.HomeDir()
		kubeCfg = filepath.Join(home, ".kube", "config")
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeCfg)
	if err != nil {
		return nil, err
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return client, nil
}
