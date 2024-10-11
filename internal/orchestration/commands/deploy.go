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

			file, _ := cmd.Flags().GetString("file")
			gateway, _ := cmd.Flags().GetBool("gateway")
			monitoring, _ := cmd.Flags().GetBool("monitoring")
			ctx := context.Background()
			if err := kubernetesClient.WatchDeployment(ctx, "argocd", "argo-helm-argocd-server", false); err != nil {
				log.Fatal(fmt.Errorf("failed to watch argocd deployment: %w", err))
			}

			err := orchestration.CreateNamespace(ctx, kubernetesClient, cfg.Namespace)
			if err != nil {
				log.Fatal(fmt.Errorf("failed to create namespace: %w", err))
			}

			if err = orchestration.AddRepoCredentials(ctx, argoClient); err != nil {
				log.Fatal(err)
			}
			if cfg == nil {
				log.Fatal("config is required")
			}
			fmt.Println(cfg.UserConfig)
			fmt.Println(cfg.UserConfig.TemplateConfig)

			templateFile := cfg.UserConfig.TemplateConfig.TemplateFile
			if file != "" {
				templateFile = file
			}
			m := cfg.MonitoringConfig.Collectors.ToStringMap()

			if err = synthAll(templateFile, cfg.OutputDir, m); err != nil {
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

			}

			if gateway {
				if err = orchestration.DeployGateway(ctx, argoClient, kubernetesClient, ns); err != nil {
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
	cmd.Flags().StringP("file", "f", "", "Path to a file containing the template definition")
	cmd.Flags().BoolP("gateway", "g", true, "deploy with gateway")
	cmd.Flags().BoolP("monitoring", "m", true, "deploy with monitoring")
	return cmd
}

func synthAll(file, outputDir string, monitoring map[string]string) error {
	if err := kplus.SynthTemplate(file, outputDir, monitoring); err != nil {
		log.Fatal(err)
	}

	if err := kplus.SynthGateway(outputDir, "galah-testbed"); err != nil {
		log.Fatal(err)
	}
	return nil
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

			if err := orchestration.AddRepoCredentials(ctx, argoClient); err != nil {
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
			synth, _ := cmd.Flags().GetBool("synth")

			if synth {
				if err := kplus.SynthGateway(outDir, ns); err != nil {
					fmt.Println("hmm")
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

			if err = orchestration.AddRepoCredentials(ctx, argoClient); err != nil {
				log.Fatal("Adding Repo Credentials", err)
			}
			if synth {
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
	cmd.Flags().BoolP("synth", "s", false, "synth gateway manifests")
	return cmd
}
