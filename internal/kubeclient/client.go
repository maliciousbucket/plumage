package kubeclient

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"os"
	"path/filepath"
)

type Client struct {
	kubeClient *kubernetes.Clientset
	helmClient *helmClient
}

func New() (*kubernetes.Clientset, error) {
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
