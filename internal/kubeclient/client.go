package kubeclient

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"os"
	"path/filepath"
)

// Client TODO: Change to interface
type Client struct {
	Client k8sClient
}

type k8sClient struct {
	kubeClient *kubernetes.Clientset
}

func NewClient() (*Client, error) {
	kubeClient, err := newClientset()
	if err != nil {
		return nil, err
	}
	client := k8sClient{
		kubeClient: kubeClient,
	}
	return &Client{Client: client}, nil
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
