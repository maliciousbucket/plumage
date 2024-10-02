package helm

import (
	"context"
	"fmt"
)

type BaseChartOpts struct {
	KubeStateMetrics string
	CertManager      string
	MetricsServer    string
	PromOperatorCRDs string
	ArgoCD           string
	ArgoValues       string
}

func (c *helmClient) InstallBaseCharts(ctx context.Context, opts *BaseChartOpts, replace bool) error {
	metricsErr := c.installKubeMetricsServer(ctx, opts.MetricsServer, replace)
	if metricsErr != nil {
		return fmt.Errorf("installing metrics server: %w", metricsErr)
	}

	certManagerErr := c.installCertManager(ctx, opts.CertManager, replace)
	if certManagerErr != nil {
		return fmt.Errorf("installing cert manager: %w", certManagerErr)
	}

	kubeStateErr := c.installKubeStateMetrics(ctx, opts.KubeStateMetrics, replace)
	if kubeStateErr != nil {
		return fmt.Errorf("installing kube-state: %w", kubeStateErr)
	}
	promOperatorErr := c.installPromOperatorCrds(ctx, opts.PromOperatorCRDs, replace)
	if promOperatorErr != nil {
		return fmt.Errorf("installing prometheus operator: %w", promOperatorErr)
	}

	argoErr := c.installArgo(ctx, opts.ArgoCD, opts.ArgoValues, replace)
	if argoErr != nil {
		return fmt.Errorf("installing ArgoCD: %w", argoErr)
	}

	return nil
}

func (c *helmClient) InstallArgoChart(ctx context.Context, version, valuesFile string) error {
	err := c.installArgo(ctx, version, valuesFile, false)
	if err != nil {
		return fmt.Errorf("installing argo chart: %w", err)
	}
	return nil
}

func (c *helmClient) InstallKubeMetricsServerChart(ctx context.Context, version string, replace bool) error {
	err := c.installKubeMetricsServer(ctx, version, replace)
	if err != nil {
		return fmt.Errorf("installing kube-metrics server chart: %w", err)
	}
	return nil
}

func (c *helmClient) InstallKubeStateMetricsServerChart(ctx context.Context, version string, replace bool) error {
	err := c.installKubeStateMetrics(ctx, version, replace)
	if err != nil {
		return fmt.Errorf("installing kube-state-metrics server chart: %w", err)
	}
	return nil
}

func (c *helmClient) InstallCertManagerChart(ctx context.Context, version string, replace bool) error {
	err := c.installCertManager(ctx, version, replace)
	if err != nil {
		return fmt.Errorf("installing cert manager chart: %w", err)
	}
	return nil
}

func (c *helmClient) InstallPromOperatorCRDs(ctx context.Context, version string, replace bool) error {
	err := c.installPromOperatorCrds(ctx, version, replace)
	if err != nil {
		return fmt.Errorf("installing prometheus operator crds: %w", err)
	}
	return nil
}

func (c *helmClient) installKubeMetricsServer(ctx context.Context, version string, replace bool) error {
	config := &ChartConfig{
		Repository:  "https://kubernetes-sigs.github.io/metrics-server/",
		Namespace:   "",
		Name:        "metrics-server",
		ReleaseName: "metrics-server",
		Version:     version,
		Replace:     replace,
		ValuesFiles: nil,
		Local:       false,
		SkipCRDs:    false,
		UpgradeCRDs: true,
		Labels:      nil,
		Lint:        false,
	}

	c.setNamespace(config.Namespace)

	if err := c.installChart(ctx, config); err != nil {
		return err
	}

	return nil
}

func (c *helmClient) installCertManager(ctx context.Context, version string, replace bool) error {
	config := &ChartConfig{
		Repository:  "https://charts.jetstack.io",
		Namespace:   "cert-manager",
		Name:        "cert-manager",
		ReleaseName: "cert-manager",
		Version:     version,
		Replace:     replace,
		ValuesFiles: nil,
		Local:       false,
		SkipCRDs:    true,
		UpgradeCRDs: false,
		Labels:      nil,
		Lint:        false,
	}
	c.setNamespace(config.Namespace)

	if err := c.installChart(ctx, config); err != nil {
		return err
	}
	return nil
}

func (c *helmClient) installKubeStateMetrics(ctx context.Context, version string, replace bool) error {
	config := &ChartConfig{
		Repository:  "https://prometheus-community.github.io/helm-charts",
		Namespace:   "galah-monitoring",
		Name:        "prometheus-community",
		ReleaseName: "kube-state-metrics",
		Version:     version,
		Replace:     replace,
		ValuesFiles: nil,
		Local:       false,
		SkipCRDs:    false,
		UpgradeCRDs: true,
		Labels:      nil,
		Lint:        false,
	}

	c.setNamespace(config.Namespace)

	if err := c.installChart(ctx, config); err != nil {
		return err
	}

	return nil
}

func (c *helmClient) installPromOperatorCrds(ctx context.Context, version string, replace bool) error {
	config := &ChartConfig{
		Repository:  "https://prometheus-community.github.io/helm-charts",
		Namespace:   "default",
		Name:        "prometheus-community",
		ReleaseName: "prometheus-operator-crds",
		Version:     version,
		Replace:     true,
		ValuesFiles: nil,
		Local:       false,
		SkipCRDs:    false,
		UpgradeCRDs: true,
		Labels:      nil,
		Lint:        false,
	}

	c.setNamespace(config.Namespace)

	if err := c.installChart(ctx, config); err != nil {
		return err
	}

	return nil
}

func (c *helmClient) installArgo(ctx context.Context, version, valuesFile string, replace bool) error {
	valuesFiles := []string{valuesFile}
	values := map[string]interface{}{
		"server": map[string]interface{}{
			"ingress": map[string]interface{}{
				"enabled": true,
			},
			"service": map[string]interface{}{
				"type": "NodePort",
			},
		},
	}

	path := fmt.Sprintf("https://github.com/argoproj/argo-helm/releases/download/%s/%s.tgz", version, version)

	config := &ChartConfig{
		Repository:  path,
		Namespace:   "argocd",
		Name:        "argocd",
		ReleaseName: "argo-helm",
		Version:     version,
		Replace:     replace,
		ValuesFiles: valuesFiles,
		Values:      values,
		Local:       false,
		SkipCRDs:    false,
		UpgradeCRDs: true,
		Labels:      nil,
		Lint:        false,
	}
	return c.installChart(ctx, config)
}
