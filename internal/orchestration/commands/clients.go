package commands

import (
	"github.com/maliciousbucket/plumage/internal/argocd"
	"github.com/maliciousbucket/plumage/internal/helm"
	"github.com/maliciousbucket/plumage/internal/kubeclient"
)

var (
	argoClient       *argocd.Client
	kubernetesClient kubeclient.Client
	helmClient       helm.Client
)

func newArgoClient() error {
	if argoClient != nil {
		return nil
	}
	conn, err := argocd.GetConnection()
	if err != nil {
		return err
	}
	argo, err := argocd.NewClient(conn)
	if err != nil {
		return err
	}
	argoClient = argo
	return nil
}

func newKubeClient() error {
	if kubernetesClient != nil {
		return nil
	}
	kube, err := kubeclient.NewClient()
	if err != nil {
		return err
	}
	kubernetesClient = kube
	return nil
}

func newHelmClient() error {
	if helmClient != nil {
		return nil
	}
	opts := &helm.ClientCfg{}
	chartClient, err := helm.NewClient(opts)
	if err != nil {
		return err
	}
	helmClient = chartClient
	return nil
}
