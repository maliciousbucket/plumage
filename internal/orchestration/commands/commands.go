package commands

import (
	"context"
	"errors"
	"fmt"
	"github.com/maliciousbucket/plumage/internal/orchestration"
	"github.com/maliciousbucket/plumage/pkg/config"
	"github.com/spf13/cobra"
	"log"
	"time"
)

var (
	targetDir   = ""
	targetFiles = ""
	envFiles    = ""
	message     = ""
	chart       = false
	service     = ""
	resource    = ""
	app         = ""
)

func CommitPushCmd(configDir, fileName string, cfg *config.AppConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "commit",
		Short: "commit push",
		Run: func(cmd *cobra.Command, args []string) {
			ghCfg, err := config.NewGithubConfig(configDir, fileName)
			if err != nil {
				log.Fatal(err)
			}
			extraEnv, _ := cmd.Flags().GetString("extra-env")
			if extraEnv != "" {
				err = ghCfg.LoadExtraEnv(extraEnv)
				if err != nil {
					log.Fatal(err)
				}
			}
			message, _ = cmd.Flags().GetString("message")
			ctx := context.Background()
			response, err := orchestration.CommitAndPush(ctx, ghCfg, ghCfg.TargetDir, message)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(response)

		},
	}
	cmd.Flags().StringVarP(&targetDir, "target-dir", "t", "", "target directory")
	cmd.Flags().StringVarP(&targetFiles, "target-files", "f", "", "comma seperated target files")
	cmd.Flags().StringVarP(&envFiles, "env-files", "e", "", "comma seperated environment files")
	cmd.Flags().StringVarP(&message, "message", "m", "", "commit message")

	cmd.AddCommand(commitManifestsCmd(configDir, fileName, cfg))
	cmd.AddCommand(commitGatewayCommand(configDir, fileName, cfg))

	err := cmd.MarkFlagRequired("message")
	if err != nil {
		fmt.Printf("Commit Message must be set")
	}

	return cmd
}

func commitManifestsCmd(configDir, fileName string, cfg *config.AppConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "manifests",
		Short: "commit manifests",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.ValidateFlagGroups(); err != nil {
				return err
			}
			if err := cmd.ValidateRequiredFlags(); err != nil {
				return err
			}

			if app == "" {
				return errors.New("must specify app with -a")
			}

			ghCfg, err := config.NewGithubConfig(configDir, fileName)
			if err != nil {
				log.Fatal(err)
			}

			ctx := context.Background()
			if chart {
				path := fmt.Sprintf("dist/%s", app)
				cmtMsg := fmt.Sprintf("plumage manifests - chart: %s - %s", app, time.Now().String())
				chartCmt, commitErr := orchestration.CommitAndPush(ctx, ghCfg, path, cmtMsg)
				if commitErr != nil {
					return commitErr
				}
				fmt.Printf("Commit Created: %+v", chartCmt)
				return nil
			}
			if resource != "" && service == "" {
				log.Fatal("You need to specify a service for the resource")
			}

			if service != "" && resource != "" {
				cmtMsg := fmt.Sprintf("plumage manifests - resource: %s - %s", resource, time.Now().String())
				path := fmt.Sprintf("%s/%s", cfg.OutputDir, resource)
				resourceCmt, commitErr := orchestration.CommitAndPushResource(ctx, ghCfg, path, service, resource, cmtMsg)
				if commitErr != nil {
					return commitErr
				}
				fmt.Printf("Commit Created: %+v", resourceCmt)
				return nil
			}

			if service != "" {
				cmtMsg := fmt.Sprintf("plumage manifests - service: %s - %s", service, time.Now().String())
				path := fmt.Sprintf("%s/%s/%s", cfg.OutputDir, app, service)
				serviceCmt, commitErr := orchestration.CommitAndPushService(ctx, ghCfg, path, service, cmtMsg)
				if commitErr != nil {
					return commitErr
				}
				fmt.Printf("Commit Created: %+v", serviceCmt)
			}
			return nil

		},
	}

	cmd.Flags().BoolVarP(&chart, "chart", "c", false, "commit synthesised chart")
	cmd.Flags().StringVarP(&app, "app", "a", "", "app name")
	cmd.Flags().StringVarP(&service, "service", "s", "", "commit synthesised service")
	cmd.Flags().StringVarP(&resource, "resource", "r", "", "commit synthesised resource")
	cmd.Flags().StringVarP(&envFiles, "env-files", "e", "", "comma seperated environment files")
	cmd.Flags().StringVarP(&message, "message", "m", "", "commit message")

	cmd.MarkFlagsMutuallyExclusive("chart", "service")
	cmd.MarkFlagsMutuallyExclusive("chart", "resource")
	cmd.MarkFlagsOneRequired("chart", "service", "resource")
	_ = cmd.MarkFlagRequired("app")
	return cmd
}

func commitGatewayCommand(configDir, fileName string, cfg *config.AppConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gateway",
		Short: "Commit traefik gateway manifests",
		RunE: func(cmd *cobra.Command, args []string) error {

			ghCfg, err := config.NewGithubConfig(configDir, fileName)
			if err != nil {
				return err
			}
			ctx := context.Background()
			cmt, cmtErr := orchestration.CommitAndPushGateway(ctx, ghCfg, cfg.OutputDir)
			if cmtErr != nil {
				return cmtErr
			}
			log.Printf("Commit Created: %+v", cmt)
			return nil
		},
	}
	return cmd
}

func ArgoAuthCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "argo-auth",
		Short: "Add GitHub credentials to ArgoCD",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return newArgoClient()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.ValidateFlagGroups(); err != nil {
				return err
			}
			if err := cmd.ValidateRequiredFlags(); err != nil {
				return err
			}
			repo, _ := cmd.Flags().GetBool("repo")
			password, _ := cmd.Flags().GetBool("password")

			if repo {
				ctx := context.Background()
				if err := argoClient.AddRepoCredentials(ctx); err != nil {
					return fmt.Errorf("failed to add repo credentials: %w", err)
				}
				log.Println("Added repo credentials")
				return nil
			}
			if password {
				ctx := context.Background()
				if err := newKubeClient(); err != nil {
					return err
				}
				pass, err := kubernetesClient.GetArgoPassword(ctx, "argocd")
				if err != nil {
					return err
				}
				fmt.Println("I love security")
				fmt.Println(pass)
				return nil
			}
			return nil
		},
	}
	cmd.Flags().BoolP("repo", "r", false, "Add GitHub credentials to ArgoCD")
	cmd.Flags().BoolP("password", "p", false, "Get the ArgoCD Admin Password")
	cmd.MarkFlagsOneRequired("password", "repo")
	cmd.MarkFlagsMutuallyExclusive("password", "repo")
	return cmd
}

func ExposeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "expose",
		Short: "Expose as service for testing",
		Long:  "expose -p <port> -n <node port> <service> <namespace>",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return newKubeClient()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.ValidateRequiredFlags(); err != nil {
				return err
			}
			if err := cobra.MinimumNArgs(2)(cmd, args); err != nil {
				return err
			}
			nodePort, _ := cmd.Flags().GetInt("nodePort")
			port, _ := cmd.Flags().GetInt("port")
			if port == 0 {
				return fmt.Errorf("port must be specified")
			}
			if nodePort < 30000 || nodePort > 32767 {
				return fmt.Errorf("node port must be between 30000 and 32767")
			}
			if args[0] == "" {
				return fmt.Errorf("service name must be specified")
			}
			if args[1] == "" {
				return fmt.Errorf("namespace must be specified")
			}

			ctx := context.Background()
			if err := kubernetesClient.ExposeService(ctx, args[1], args[0], port, nodePort); err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().IntP("port", "p", 0, "Port to expose service on")
	cmd.Flags().IntP("nodePort", "n", 0, "Node port to expose service on")
	_ = cmd.MarkFlagRequired("port")
	_ = cmd.MarkFlagRequired("nodePort")
	return cmd
}
