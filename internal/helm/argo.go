package helm

import (
	"context"
	"fmt"
	helmc "github.com/mittwald/go-helm-client"
	"github.com/mittwald/go-helm-client/values"
	"os"
)

const (
	argoHelmRepo = "https://github.com/argoproj/argo-helm/releases/download/argo-cd-7.5.2/argo-cd-7.5.2.tgz"
	argoVersion  = "argo-cd-7.5.2"
)

type ArgoConfig struct {
	Namespace   string
	ChartName   string
	Version     string
	ValuesFiles []string
}

func InstallArgo(ctx context.Context, clientCfg *ClientCfg, ns string, opts ...ArgoOpts) error {
	argo := defaultArgoConfig(ns)

	client, err := New(clientCfg)
	if err != nil {
		return err
	}

	res, err := client.InstallArgo(ctx, argo, opts...)
	if err != nil {
		return err
	}
	fmt.Printf("Successfully installed helm chart %s in namespace %s\n", res.Name, res.Namespace)
	return nil
}

func (c *ArgoConfig) ChartSpec() *helmc.ChartSpec {
	valuesOpts := values.Options{
		ValueFiles: c.ValuesFiles,
	}

	return &helmc.ChartSpec{
		Namespace:       "argocd",
		ChartName:       c.ChartName,
		ReleaseName:     "argocd-helm",
		Version:         c.Version,
		ValuesOptions:   valuesOpts,
		UpgradeCRDs:     true,
		CreateNamespace: true,
	}
}

func defaultArgoConfig(namespace string) *ArgoConfig {
	return &ArgoConfig{
		Namespace: namespace,
		ChartName: argoHelmRepo,
		Version:   argoVersion,
	}
}

type ArgoOpts func(*ArgoConfig) error

func WithNamespace(namespace string) ArgoOpts {
	return func(config *ArgoConfig) error {
		config.Namespace = namespace
		return nil
	}
}

func WithLocalChart(filePath string) ArgoOpts {
	return func(config *ArgoConfig) error {
		path, err := os.Stat(filePath)
		if err != nil {
			return err
		}
		if path.IsDir() {
			return fmt.Errorf("%s is a directory", filePath)
		}
		config.ChartName = filePath
		return nil
	}
}

func FromRemote(url string) ArgoOpts {
	return func(config *ArgoConfig) error {
		config.ChartName = url
		return nil
	}
}

func WithValuesFiles(files []string) ArgoOpts {
	return func(config *ArgoConfig) error {
		config.ValuesFiles = append(config.ValuesFiles, files...)
		return nil
	}
}

func WithValuesFile(file string) ArgoOpts {
	return func(config *ArgoConfig) error {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			return fmt.Errorf("%s does not exist", file)
		}
		config.ValuesFiles = append(config.ValuesFiles, file)
		return nil
	}
}

func WithVersion(v string) ArgoOpts {
	return func(config *ArgoConfig) error {
		config.Version = v
		return nil
	}
}
