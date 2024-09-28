package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/maliciousbucket/plumage/internal/orchestration"
	"github.com/maliciousbucket/plumage/pkg/config"
	"github.com/maliciousbucket/plumage/pkg/kplus"
	"github.com/spf13/cobra"
	"log"
	"time"
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
			if err := kubernetesClient.WatchDeployment(ctx, "argocd", "argocd-helm-server"); err != nil {
				log.Fatal(fmt.Errorf("failed to watch argocd deployment: %w", err))
			}

			if err := argoClient.AddRepoCredentials(ctx); err != nil {
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

			if err := synthAll(templateFile, cfg.OutputDir, m); err != nil {
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
				if err = handleMonitoring(ctx); err != nil {
					log.Fatal(err)
				}
			}

			if gateway {
				if err = handleGateway(ctx); err != nil {
					log.Fatal(err)
				}
			}
			if appProject, _ := argoClient.GetProject(ctx, appName); appProject != nil {
				if _, err = argoClient.CreateProject(ctx, appName); err != nil {
					log.Fatal(err)
				}

			}

			if err = argoClient.CreateServiceApplications(ctx, appName, services); err != nil {
				log.Fatal(err)
			}

			if err = deployAndWaitForApp(ctx, ns, appName, services); err != nil {
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
	ghCfg, err := config.NewGithubConfig(cfg.ConfigDir, "github.yaml")
	if err != nil {
		log.Fatal(err)
	}
	path := fmt.Sprintf("%s/%s", cfg.OutputDir, app)
	msg := fmt.Sprintf("Plumage Manifests - %s - %s", app, time.Now().String())
	templateCommit, err := orchestration.CommitAndPush(ctx, ghCfg, path, msg)
	if err != nil {
		return err
	}

	gatewayCommit, err := orchestration.CommitAndPushGateway(ctx, ghCfg, cfg.OutputDir)
	if err != nil {
		return err
	}

	templateJson, err := json.MarshalIndent(templateCommit, "", "    ")
	if err != nil {
		return err
	}
	gatewayJson, err := json.MarshalIndent(gatewayCommit, "", "    ")
	if err != nil {
		return err
	}
	log.Printf("\nCommits for App %s successful\n", app)
	log.Println(string(templateJson))
	log.Println(string(gatewayJson))
	return nil
}

func handleMonitoring(ctx context.Context) error {
	if monitoringProj, _ := argoClient.GetProject(ctx, "galah-monitoring"); monitoringProj == nil {
		if err := argoClient.CreateMonitoringProject(ctx); err != nil {
			return err
		}
	} else {
		if err := argoClient.SyncProject(ctx, "galah-monitoring"); err != nil {
			return err
		}
	}
	return nil

}

func handleGateway(ctx context.Context) error {
	if gatewayProj, _ := argoClient.GetProject(ctx, "ingress"); gatewayProj == nil {
		if err := argoClient.CreateIngressProject(ctx); err != nil {
			return err
		}
	} else {
		if err := argoClient.SyncProject(ctx, "ingress"); err != nil {
			return err
		}
	}
	return nil
}

func deployAndWaitForApp(ctx context.Context, ns, app string, services []string) error {
	if err := argoClient.CreateServiceApplications(ctx, app, services); err != nil {
		return err
	}

	if err := argoClient.SyncProject(ctx, app); err != nil {
		return err
	}
	if err := kubernetesClient.WatchAppDeployment(ctx, ns, services); err != nil {
		return err
	}
	return nil
}
