package commands

import (
	"context"
	"fmt"
	"github.com/maliciousbucket/plumage/internal/orchestration"
	"github.com/maliciousbucket/plumage/pkg/config"
	"github.com/maliciousbucket/plumage/pkg/kplus"
	"github.com/spf13/cobra"
	"log"
)

func DeployTemplateCommand(cfg *config.AppConfig) *cobra.Command {
	var gateway bool
	var monitoring bool
	var synthApp bool
	var synthGateway bool
	var synthAllManifests bool
	var template string
	cmd := &cobra.Command{
		Use:   "template",
		Short: "Deploy a template",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := newArgoClient(); err != nil {
				return err
			}

			if err := newKubeClient(); err != nil {
				return err
			}

			return nil
		},
		Args: func(cmd *cobra.Command, args []string) error {
			if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
				return err
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if args[0] == "" {
				return fmt.Errorf("app name is required")
			}
			//TODO: Change to orchestration method
			ctx := context.Background()
			if err := kubernetesClient.WatchDeployment(ctx, "argocd", "argo-helm-argocd-server", false); err != nil {
				log.Fatal(fmt.Errorf("failed to watch argocd deployment: %w", err))
			}

			err := orchestration.CreateNamespace(ctx, kubernetesClient, cfg.Namespace)
			if err != nil {
				log.Fatal(fmt.Errorf("failed to create namespace: %w", err))
			}

			if err = orchestration.AddRepoCredentials(ctx, argoClient, ".env"); err != nil {
				log.Fatal(err)
			}
			if cfg == nil {
				log.Fatal("config is required")
			}

			templateFile := cfg.UserConfig.TemplateConfig.TemplateFile
			if template != "" {
				templateFile = template
			}
			m := cfg.MonitoringConfig.Collectors.ToStringMap()

			synthOpts := &orchestration.SynthOpts{
				SynthTemplate: synthApp,
				SynthGateway:  synthGateway,
				SynthTests:    false,
				TemplateFile:  templateFile,
				OutputDir:     cfg.OutputDir,
				Namespace:     cfg.Namespace,
				Monitoring:    m,
			}

			if err = orchestration.SynthDeployment(synthOpts); err != nil {
				log.Fatal(err)
			}

			services, ns, appName, err := kplus.GetServices(templateFile)
			if err != nil {
				log.Fatal(err)
			}

			if err = commitAndPushAll(ctx, cfg, appName); err != nil {
				log.Fatal(err)
			}

			if monitoring {
				if err = orchestration.DeployMonitoring(ctx, argoClient, kubernetesClient); err != nil {
					log.Fatal(err)
				}

				if err = orchestration.WaitForMonitoringDeployment(ctx, kubernetesClient); err != nil {
					log.Fatal(err)
				}
			}

			if gateway {
				if err = orchestration.DeployGateway(ctx, argoClient, kubernetesClient, ns); err != nil {
					log.Fatal(err)
				}

				if err = orchestration.WaitForGatewayDeployment(ctx, kubernetesClient, ns); err != nil {
					log.Fatal(err)
				}

			}
			if appProject, _ := argoClient.GetProject(ctx, appName); appProject != nil {
				if _, err = argoClient.CreateProject(ctx, appName); err != nil {
					log.Fatal(err)
				}

			}

			if err = argoClient.CreateApplicationProject(ctx, appName); err != nil {
				log.Fatal(err)
			}

			if err = orchestration.DeployAndWaitForApp(ctx, argoClient, kubernetesClient, ns, appName, services); err != nil {
				log.Fatal(err)
			}

			log.Printf("\n Successfully Deployed %s in %s", appName, ns)

			return nil
		},
	}
	cmd.Flags().StringVar(&template, "template", "", "Path to a file containing the template definition")
	cmd.Flags().BoolVar(&gateway, "gateway", false, "Deploy with gateway")
	cmd.Flags().BoolVar(&monitoring, "monitoring", false, "Deploy with galah-observability monitoring")
	cmd.Flags().BoolVar(&synthApp, "synth-app", false, "Synth app manifests")
	cmd.Flags().BoolVar(&synthGateway, "synth-gateway", false, "Synth gateway manifests")
	cmd.Flags().BoolVar(&synthAllManifests, "synth-all", false, "Synth all manifests")

	cmd.MarkFlagsMutuallyExclusive("synth-all", "synth-gateway")
	cmd.MarkFlagsMutuallyExclusive("synth-all", "synth-app")

	return cmd
}

func commitAndPushAll(ctx context.Context, cfg *config.AppConfig, app string) error {
	templateCommit, gatewayCommit, err := orchestration.CommitAndPushTestBed(ctx, cfg, app)
	if err != nil {
		return err
	}

	if err = prettyPrint(gatewayCommit); err != nil {
		return err
	}

	if err = prettyPrint(templateCommit); err != nil {
		return err
	}

	return nil
}

func DeployMonitoringCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "monitoring",
		Short: "Deploy monitoring infrastructure",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := newArgoClient(); err != nil {
				return err
			}
			if err := newKubeClient(); err != nil {
				return err
			}
			return nil

		},
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()

			if err := orchestration.AddRepoCredentials(ctx, argoClient, ".env"); err != nil {
				log.Fatal(err)
			}
			if err := orchestration.DeployMonitoring(ctx, argoClient, kubernetesClient); err != nil {
				log.Fatal(err)
			}
		},
	}
	return cmd
}

func DeployGatewayCommand(configDir, outDir, ns string) *cobra.Command {
	var synthGateway bool
	cmd := &cobra.Command{
		Use:   "gateway",
		Short: "Deploy traefik gateway",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := newArgoClient(); err != nil {
				return err
			}
			if err := newKubeClient(); err != nil {
				return err
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {

			if synthGateway {
				if err := kplus.SynthGateway(outDir, ns); err != nil {
					log.Fatal(err)
				}
			}

			ctx := context.Background()

			err := orchestration.CreateNamespace(ctx, kubernetesClient, ns)
			if err != nil {
				log.Fatal("Error creating namespace", err)
			}

			ghCfg, err := config.NewGithubConfig(configDir, "github.yaml")
			if err != nil {
				log.Fatal("creating GitHub Config", err)
			}

			if err = orchestration.AddRepoCredentials(ctx, argoClient, ".env"); err != nil {
				log.Fatal("Adding Repo Credentials", err)
			}
			if synthGateway {
				res, err := orchestration.CommitAndPushGateway(ctx, ghCfg, outDir)

				if err != nil {
					log.Fatal(err)
				}

				if err = prettyPrint(res); err != nil {
					log.Fatal(err)
				}
			}

			if err = orchestration.DeployGateway(ctx, argoClient, kubernetesClient, ns); err != nil {
				log.Fatal(err)
			}

			log.Println("Gateway has been deployed")
		},
	}
	cmd.Flags().BoolVar(&synthGateway, "synth", false, "synth gateway manifests")
	return cmd
}
