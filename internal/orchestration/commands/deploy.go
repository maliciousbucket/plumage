package commands

import (
	"context"
	"errors"
	"fmt"
	"github.com/maliciousbucket/plumage/internal/argocd"
	"github.com/maliciousbucket/plumage/internal/kubeclient"
	"github.com/maliciousbucket/plumage/internal/orchestration"
	"github.com/maliciousbucket/plumage/pkg/config"
	"github.com/maliciousbucket/plumage/pkg/kplus"
	"github.com/spf13/cobra"
	"log"
	"sync"
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
			if err := kubernetesClient.WatchDeployment(ctx, "argocd", "argocd-helm-server", false); err != nil {
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
				if err = handleGateway(ctx, ns); err != nil {
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
	log.Printf("\nCommits for App %s successful\n", app)
	gatewayCommit, err := orchestration.CommitAndPushGateway(ctx, ghCfg, cfg.OutputDir)
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

			if err := argoClient.AddRepoCredentials(ctx); err != nil {
				log.Fatal(err)
			}
			if err := handleMonitoring(ctx); err != nil {
				log.Fatal(err)
			}
		},
	}
	return cmd
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

	resources := map[string]string{
		"alloy":          "galah-monitoring",
		"nginx":          "gateway",
		"tempo":          "galah-tracing",
		"loki":           "galah-logging",
		"mimir":          "galah-monitoring",
		"grafana":        "galah-monitoring",
		"minio-operator": "minio-store",
		"minio-tenant":   "minio-store",
	}
	var watchErr error
	errChan := make(chan error)
	var wg sync.WaitGroup
	for res, namespace := range resources {
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer close(errChan)
			errChan <- watchInfrastructure(ctx, kubernetesClient, namespace, res)
		}()
	}
	wg.Wait()

	i := 0
	for i < len(resources) {
		select {
		case err := <-errChan:
			if err != nil {
				watchErr = errors.Join(watchErr, err)
			}
		case <-ctx.Done():
			watchErr = errors.Join(watchErr, ctx.Err())
			return watchErr
		}
	}

	return watchErr

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
					log.Fatal(err)
				}
			}

			ctx := context.Background()

			ghCfg, err := config.NewGithubConfig(configDir, "github.yaml")
			if err != nil {
				log.Fatal(err)
			}

			if err = argoClient.AddRepoCredentials(ctx); err != nil {
				log.Fatal(err)
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

			if err = handleGateway(ctx, ns); err != nil {
				log.Fatal(err)
			}
		},
	}
	cmd.Flags().BoolP("synth", "s", false, "synth gateway manifests")
	return cmd
}

func handleGateway(ctx context.Context, ns string) error {
	if gatewayProj, _ := argoClient.GetProject(ctx, "ingress"); gatewayProj == nil {
		if err := argoClient.CreateIngressProject(ctx); err != nil {
			return err
		}
	} else {
		params := &argocd.AppQueryParams{Options: []argocd.AppQueryFunc{
			argocd.WithProject("ingress"),
		}}
		apps, _ := argoClient.ListApplications(ctx, params)
		if apps == nil || len(apps.Items) == 0 {
			if err := argoClient.CreateIngressApp(ctx); err != nil {
				return err
			}
		}

		if err := argoClient.SyncProject(ctx, "ingress"); err != nil {
			return err
		}
	}
	time.Sleep(5 * time.Second)
	errChan := make(chan error)
	go func() {

		errChan <- watchInfrastructure(ctx, kubernetesClient, ns, "traefik")
		close(errChan)
	}()
	select {
	case err := <-errChan:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}

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

func watchInfrastructure(ctx context.Context, client kubeclient.Client, ns, name string) error {
	if err := client.WaitAppPods(ctx, ns, name, 2*time.Minute); err != nil {
		return err
	}
	return nil
}
