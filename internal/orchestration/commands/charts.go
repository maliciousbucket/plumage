package commands

import (
	"context"
	"github.com/maliciousbucket/plumage/internal/orchestration"
	"github.com/maliciousbucket/plumage/pkg/config"
	"github.com/spf13/cobra"
	"log"
)

func ChartsCmd(cfg *config.ChartConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "charts",
		Short: "Manage Helm Charts",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return newHelmClient()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	cmd.AddCommand(installChartCmd(cfg))
	return cmd
}

func installChartCmd(cfg *config.ChartConfig) *cobra.Command {
	var (
		installArgo                bool
		valuesFile                 string
		installPromOperatorCrds    bool
		installCertManager         bool
		installKubeStateMetrics    bool
		installKubeMetricsServer   bool
		installK6Operator          bool
		installKubePrometheusStack bool
		replace                    bool
		version                    string
	)
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Install Helm Charts",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.ValidateRequiredFlags(); err != nil {
				return err
			}

			if err := cmd.ValidateFlagGroups(); err != nil {
				return err
			}
			opts := cfg.ToBaseOpts()
			ctx := context.Background()

			if installArgo {
				argoVersion := opts.ArgoCD
				if version != "" {
					argoVersion = version
				}
				values := opts.ArgoValues
				if valuesFile != "" {
					values = valuesFile
				}
				if err := helmClient.InstallArgoChart(ctx, argoVersion, values); err != nil {
					log.Fatal(err)
				}
				if err := newKubeClient(); err != nil {
					log.Fatal(err)
				}
				if err := orchestration.WatchArgoDeployment(ctx, kubernetesClient); err != nil {
					log.Fatal(err)
				}
				return nil
			}

			if installPromOperatorCrds {
				promVersion := opts.PromOperatorCRDs
				if version != "" {
					promVersion = version
				}
				values := opts.PromOperatorValues
				if valuesFile != "" {
					values = valuesFile
				}
				if err := helmClient.InstallPromOperatorCRDs(ctx, promVersion, values, replace); err != nil {
					log.Fatal(err)
				}
				return nil
			}

			if installCertManager {
				certVersion := opts.CertManager
				if version != "" {
					certVersion = version
				}
				values := opts.CertManagerValues
				if valuesFile != "" {
					values = valuesFile
				}
				if err := helmClient.InstallCertManagerChart(ctx, certVersion, values, replace); err != nil {
					log.Fatal(err)
				}
				return nil
			}

			if installKubeStateMetrics {
				stateMetricsVersion := opts.KubeStateMetrics
				if version != "" {
					stateMetricsVersion = version
				}
				values := opts.StateMetricsValues
				if valuesFile != "" {
					values = valuesFile
				}

				if err := helmClient.InstallKubeStateMetricsServerChart(ctx, stateMetricsVersion, values, replace); err != nil {
					log.Fatal(err)
				}
				return nil
			}

			if installKubeMetricsServer {
				metricsVersion := opts.MetricsServer
				if version != "" {
					metricsVersion = version
				}

				if err := helmClient.InstallKubeMetricsServerChart(ctx, metricsVersion, valuesFile, replace); err != nil {
					log.Fatal(err)
				}
				return nil
			}

			if installK6Operator {
				k6Version := opts.K6Operator
				if version != "" {
					k6Version = version
				}
				values := opts.K6Values
				if valuesFile != "" {
					values = valuesFile
				}
				if err := helmClient.InstallK6(ctx, k6Version, values, replace); err != nil {
					log.Fatal(err)
				}
			}

			if installKubePrometheusStack {
				stackVersion := opts.KubePrometheusStack
				if version != "" {
					stackVersion = version
				}
				values := opts.KubePrometheusStackValues
				if valuesFile != "" {
					values = valuesFile
				}
				if err := helmClient.InstallKubePrometheusStack(ctx, stackVersion, values, replace); err != nil {
					log.Fatal(err)
				}
			}

			return nil
		},
	}
	cmd.Flags().BoolVar(&installArgo, "argo", false, "Install Argo")
	cmd.Flags().StringVar(&valuesFile, "values-file", "", "Chart values.yaml file")
	cmd.Flags().BoolVar(&installPromOperatorCrds, "prom-operator-crds", false, "Install Prometheus Operator CRDs")
	cmd.Flags().BoolVar(&installCertManager, "cert-manager", false, "Install Certificate Manager")
	cmd.Flags().BoolVar(&installKubeStateMetrics, "kube-state-metrics", false, "Install KubeState Metrics")
	cmd.Flags().BoolVar(&installKubeMetricsServer, "kube-metrics-server", false, "Install Metrics Server")
	cmd.Flags().BoolVar(&installK6Operator, "k6-operator", false, "Install K6 Operator")
	cmd.Flags().BoolVar(&installKubePrometheusStack, "kube-prometheus-stack", false, "Install Prometheus Stack")
	cmd.Flags().BoolVarP(&replace, "replace", "r", false, "Replace existing Helm Charts")
	cmd.Flags().StringVarP(&version, "version", "v", "", "Version")

	cmd.MarkFlagsMutuallyExclusive("argo", "prom-operator-crds", "cert-manager", "kube-state-metrics",
		"kube-metrics-server", "k6-operator", "kube-prometheus-stack")
	_ = cmd.MarkFlagFilename("values-file", "yaml", "yml")

	return cmd
}

func installBaseChartsCmd(cfg *config.ChartConfig) *cobra.Command {
	var (
		argoVersion               string
		certManagerVersion        string
		kubeStateMetricsVersion   string
		kubeMetricsServerVersion  string
		prometheusOperatorVersion string
		replace                   bool
	)
	cmd := &cobra.Command{
		Use:   "base-install",
		Short: "Install Base Helm Charts",
		Long: `Install the following charts:
- Prometheus Operator CRDs
- ArgoCD
- Kube State Metrics
- Kubernetes Metrics Server
- Cert Manager`,
		Run: func(cmd *cobra.Command, args []string) {
			opts := cfg.ToBaseOpts()
			if argoVersion != "" {
				opts.ArgoCD = argoVersion
			}
			if certManagerVersion != "" {
				opts.CertManager = certManagerVersion
			}
			if kubeStateMetricsVersion != "" {
				opts.KubeStateMetrics = kubeStateMetricsVersion
			}
			if kubeMetricsServerVersion != "" {
				opts.MetricsServer = kubeMetricsServerVersion
			}
			if prometheusOperatorVersion != "" {
				opts.PromOperatorCRDs = prometheusOperatorVersion
			}
			ctx := context.Background()
			log.Println("Installing base helm charts")
			err := helmClient.InstallBaseCharts(ctx, opts, replace)
			if err != nil {
				log.Fatal(err)
			}

		},
	}
	cmd.Flags().StringVar(&argoVersion, "argo-version", "", "Argo Version")
	cmd.Flags().StringVar(&certManagerVersion, "cert-manager-version", "", "Cert-Manager Version")
	cmd.Flags().StringVar(&kubeStateMetricsVersion, "kube-state-metrics-version", "", "KubeState Metrics Version")
	cmd.Flags().StringVar(&kubeMetricsServerVersion, "metrics-server-version", "", "KubeState Metrics Server")
	cmd.Flags().StringVar(&prometheusOperatorVersion, "prom-crds-version", "", "Prometheus Operator Version")
	cmd.Flags().BoolVarP(&replace, "replace", "r", false, "Replace existing Helm Charts")
	return cmd
}
