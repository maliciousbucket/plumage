package helm

import (
	"context"
	"fmt"
	helmc "github.com/mittwald/go-helm-client"
	"helm.sh/helm/v3/pkg/release"
	"os"
	"path/filepath"
)

type Client interface {
	InstallArgo(ctx context.Context, argo *ArgoConfig, opts ...ArgoOpts) (*release.Release, error)
	InstallBaseCharts(ctx context.Context, opts *BaseChartOpts, replace bool) error
	InstallArgoChart(ctx context.Context, version, valuesFile string) error
	InstallKubeMetricsServerChart(ctx context.Context, version string, replace bool) error
	InstallCertManagerChart(ctx context.Context, version string, replace bool) error
	InstallPromOperatorCRDs(ctx context.Context, version string, replace bool) error
	InstallKubeStateMetricsServerChart(ctx context.Context, version string, replace bool) error
	InstallK6(ctx context.Context, version string, replace bool) error
	InstallKubePrometheusStack(ctx context.Context, version, valuesFile string, replace bool) error
	InstallChart(ctx context.Context, chart *ChartConfig) error
	GetRelease(name string) (*ChartRelease, error)
	ListReleases() ([]*ReleaseMeta, error)
	GetReleaseValues(name string, allValues bool) (map[string]interface{}, error)
	UninstallRelease(name string) error
}

func NewClient(cfg *ClientCfg) (Client, error) {
	return newClient(cfg)
}

func NewClientFromConfig(cfg *ClientCfg, kubeConfig []byte) (Client, error) {
	return newClientFromConfig(cfg, kubeConfig)
}

type ClientCfg struct {
	Namespace        string
	RepositoryConfig string
	RepositoryCache  string
	RegistryConfig   string
	KubeCfgPath      string
	username         string
	password         string
}

type helmClient struct {
	Client helmc.Client
}

func newClient(cfg *ClientCfg) (*helmClient, error) {
	opts, err := cfg.kubeClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := helmc.NewClientFromKubeConf(opts)
	if err != nil {
		return nil, err
	}
	return &helmClient{Client: client}, nil
}

func (cfg *ClientCfg) kubeClientOptions() (*helmc.KubeConfClientOptions, error) {
	kubeConfig, err := getKubeConfig(cfg.KubeCfgPath)
	if err != nil {
		return nil, err
	}

	return &helmc.KubeConfClientOptions{
		Options: &helmc.Options{
			Namespace:        cfg.Namespace,
			RepositoryConfig: cfg.RepositoryConfig,
			RepositoryCache:  cfg.RepositoryCache,
			Debug:            true,
			Linting:          false,
			DebugLog:         nil,
			RegistryConfig:   cfg.RegistryConfig,
			Output:           nil,
		},
		KubeContext: "",
		KubeConfig:  kubeConfig,
	}, nil
}

func newClientFromConfig(cfg *ClientCfg, kubeConfig []byte) (*helmClient, error) {
	if kubeConfig == nil || len(kubeConfig) == 0 {
		return nil, fmt.Errorf("empty kube config provided")
	}
	opts := cfg.kubeClientOptionsFromConfig(kubeConfig)
	client, err := helmc.NewClientFromKubeConf(opts)
	if err != nil {
		return nil, err
	}
	return &helmClient{Client: client}, nil
}

func (cfg *ClientCfg) kubeClientOptionsFromConfig(kubeConfig []byte) *helmc.KubeConfClientOptions {
	return &helmc.KubeConfClientOptions{
		Options: &helmc.Options{
			Namespace:        cfg.Namespace,
			RepositoryConfig: cfg.RepositoryConfig,
			RepositoryCache:  cfg.RepositoryCache,
			Debug:            true,
			Linting:          false,
			DebugLog:         nil,
			RegistryConfig:   cfg.RegistryConfig,
			Output:           nil,
		},
		KubeContext: "",
		KubeConfig:  kubeConfig,
	}
}

func getKubeConfig(configPath string) ([]byte, error) {
	if configPath != "" {
		kubeConfig, err := os.ReadFile(configPath)
		if err != nil {
			return nil, err
		}
		return kubeConfig, nil
	}
	kubeConfig := os.Getenv("KUBECONFIG")
	if kubeConfig != "" {
		data, err := os.ReadFile(kubeConfig)
		if err != nil {
			return nil, err
		}
		return data, nil

	}
	dir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	path := filepath.Join(dir, ".kube", "config")
	config, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return config, err
}

func (c *helmClient) setNamespace(namespace string) {
	c.Client.GetSettings().SetNamespace(namespace)
}

func (c *helmClient) InstallArgo(ctx context.Context, argo *ArgoConfig, opts ...ArgoOpts) (*release.Release, error) {

	for _, opt := range opts {
		err := opt(argo)
		if err != nil {
			return nil, err
		}
	}

	chartSpec := argo.ChartSpec()

	helmOpts := &helmc.GenericHelmOptions{
		RollBack: c.Client,
	}

	res, err := c.Client.InstallOrUpgradeChart(ctx, chartSpec, helmOpts)
	if err != nil {
		return nil, err
	}

	return res, nil
}
