package integration_testing

import (
	"context"
	"github.com/maliciousbucket/plumage/internal/helm"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	argoVersion             = "argo-cd-7.6.1"
	promOperatorVersion     = "14.0.0"
	kubeStateMetricsVersion = "5.25.1"
	certManagerVersion      = "v1.15.3"
	k6OperatorVersion       = "3.9.0"
)

type HelmClient interface {
	InstallArgoChart(ctx context.Context, version, valuesFile string) error
	InstallKubeMetricsServerChart(ctx context.Context, version string, replace bool) error
	InstallCertManagerChart(ctx context.Context, version string, replace bool) error
	InstallPromOperatorCRDs(ctx context.Context, version string, replace bool) error
	InstallBaseCharts(ctx context.Context, opts *helm.BaseChartOpts, replace bool) error
	InstallK6(ctx context.Context, version string, replace bool) error
	InstallChart(ctx context.Context, chart *helm.ChartConfig) error
}

func newHelmClient(kubeConfig []byte) (HelmClient, error) {
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

	t.Run("Install CertManager Chart", func(t *testing.T) {
		err = client.InstallCertManagerChart(ctx, certManagerVersion, false)
		assert.NoError(t, err)
	})

}
