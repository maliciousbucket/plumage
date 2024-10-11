package integration_testing

import (
	"context"
	"fmt"
	"github.com/maliciousbucket/plumage/internal/helm"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	argoVersion         = "argo-cd-7.6.1"
	promOperatorVersion = "14.0.0"

	kubeStateMetricsVersion = "5.25.1"
	certManagerVersion      = "v1.15.3"
	k6OperatorVersion       = "3.9.0"
	nginxVersion            = "4.11.3"
	helmTestData            = "./testdata/helm"
)

type helmClient interface {
	InstallArgoChart(ctx context.Context, version, valuesFile string) error
	InstallKubeMetricsServerChart(ctx context.Context, version string, replace bool) error
	InstallCertManagerChart(ctx context.Context, version string, replace bool) error
	InstallPromOperatorCRDs(ctx context.Context, version string, replace bool) error
	InstallBaseCharts(ctx context.Context, opts *helm.BaseChartOpts, replace bool) error
	InstallK6(ctx context.Context, version string, replace bool) error
	InstallChart(ctx context.Context, chart *helm.ChartConfig) error
	UninstallRelease(name string) error
	GetRelease(name string) (*helm.ChartRelease, error)
}

func newHelmClient(kubeConfig []byte) (helmClient, error) {
	cfg := &helm.ClientCfg{}
	return helm.NewClientFromConfig(cfg, kubeConfig)
}

func TestInstallBaseCharts(t *testing.T) {

	kubeContainer := NewKubernetesContainer()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	err := SetupKubernetesContainer(ctx, kubeContainer)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err = kubeContainer.Container.Terminate(ctx); err != nil {
			t.Fatal(err)
		}
	}()

	config, err := kubeContainer.KubeConfig(ctx)
	if err != nil {
		t.Fatal(err)
	}
	client, err := newHelmClient(config)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Install ArgoCD Chart", func(t *testing.T) {
		err = client.InstallArgoChart(ctx, argoVersion, "")
		assert.NoError(t, err)
	})

	t.Run("Install Prometheus Operator CRDs", func(t *testing.T) {
		err = client.InstallPromOperatorCRDs(ctx, promOperatorVersion, false)
		assert.NoError(t, err)
	})

	t.Run("Install K6 Operator", func(t *testing.T) {
		err = client.InstallK6(ctx, k6OperatorVersion, false)
	})

}

//Get chart

func TestInstallChart(t *testing.T) {
	kubeContainer := NewKubernetesContainer()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	err := SetupKubernetesContainer(ctx, kubeContainer)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err = kubeContainer.Container.Terminate(ctx); err != nil {
			t.Fatal(err)
		}
	}()

	config, err := kubeContainer.KubeConfig(ctx)
	if err != nil {
		t.Fatal(err)
	}
	client, err := newHelmClient(config)
	if err != nil {
		t.Fatal(err)
	}
	t.Run("Install chart from file", func(t *testing.T) {
		repo := fmt.Sprintf("%s/ingress-nginx-%s.tgz", helmTestData, nginxVersion)
		cfg := &helm.ChartConfig{
			Repository:  repo,
			Namespace:   "default",
			Name:        "nginx",
			ReleaseName: "ingress-nginx",
			Version:     "4.11.3",
			Replace:     false,
			LocalFile:   true,
			SkipCRDs:    true,
			UpgradeCRDs: false,
			Lint:        false,
		}

		err = client.InstallChart(ctx, cfg)
		assert.NoError(t, err)
		err = client.UninstallRelease(cfg.ReleaseName)
		assert.NoError(t, err)
	})

	t.Run("Install chart from remote repository", func(t *testing.T) {
		cfg := &helm.ChartConfig{
			Repository:  "https://prometheus-community.github.io/helm-charts",
			Namespace:   "default",
			Name:        "prometheus-community",
			ReleaseName: "prometheus-operator-crds",
			Version:     promOperatorVersion,
			Replace:     false,
			LocalFile:   false,
			SkipCRDs:    false,
			UpgradeCRDs: true,
			Lint:        false,
		}
		err = client.InstallChart(ctx, cfg)
		assert.NoError(t, err)
		err = client.UninstallRelease(cfg.ReleaseName)
		assert.NoError(t, err)
	})

	t.Run("Install chart from remote file", func(t *testing.T) {
		path := fmt.Sprintf("https://github.com/argoproj/argo-helm/releases/download/%s/%s.tgz", argoVersion, argoVersion)

		cfg := &helm.ChartConfig{
			Repository:  path,
			Namespace:   "argocd",
			Name:        "argocd",
			ReleaseName: "argo-helm",
			Version:     argoVersion,
			Replace:     false,
			LocalFile:   false,
			SkipCRDs:    false,
			UpgradeCRDs: true,
			Labels:      nil,
			Lint:        false,
		}
		err = client.InstallChart(ctx, cfg)
		assert.NoError(t, err)
		err = client.UninstallRelease(cfg.ReleaseName)
		assert.NoError(t, err)
	})

	t.Run("Install chart with values file", func(t *testing.T) {

	})
}

func TestInstallInvalidCharts(t *testing.T) {
	kubeContainer := NewKubernetesContainer()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	err := SetupKubernetesContainer(ctx, kubeContainer)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err = kubeContainer.Container.Terminate(ctx); err != nil {
			t.Fatal(err)
		}
	}()

	config, err := kubeContainer.KubeConfig(ctx)
	if err != nil {
		t.Fatal(err)
	}
	client, err := newHelmClient(config)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Non-existent repository", func(t *testing.T) {
		cfg := &helm.ChartConfig{
			Repository:  "https://github.com/maliciousbucket/helm-charts/releases/download/chart/1.0/tgz",
			Namespace:   "default",
			Name:        "chart",
			ReleaseName: "chart",
			Version:     "1.0",
			Replace:     false,
			LocalFile:   false,
			SkipCRDs:    false,
			UpgradeCRDs: false,
			Labels:      nil,
			Lint:        false,
		}
		err = client.InstallChart(ctx, cfg)
		assert.Error(t, err)
	})

	t.Run("Local file with incorrect extension", func(t *testing.T) {
		repo := fmt.Sprintf("%s/ingress-nginx-%s.yaml", helmTestData, nginxVersion)
		cfg := &helm.ChartConfig{
			Repository:  repo,
			Namespace:   "default",
			Name:        "nginx",
			ReleaseName: "ingress-nginx",
			Version:     "4.11.3",
			Replace:     false,
			LocalFile:   true,
			SkipCRDs:    true,
			UpgradeCRDs: false,
			Lint:        false,
		}
		err = client.InstallChart(ctx, cfg)
		assert.Error(t, err)
	})

	t.Run("Remote chart with incorrect name", func(t *testing.T) {
		cfg := &helm.ChartConfig{
			Repository:  "https://prometheus-community.github.io/helm-charts",
			Namespace:   "default",
			Name:        "foo",
			ReleaseName: "bar",
			Version:     promOperatorVersion,
			Replace:     false,
			LocalFile:   false,
			SkipCRDs:    false,
			UpgradeCRDs: true,
			Lint:        false,
		}
		err = client.InstallChart(ctx, cfg)
		assert.Error(t, err)
	})

}
