package commands

import (
	"context"
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
)

func CommitPushCmd(configDir, fileName string, cfg *config.AppConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "commit",
		Short: "commit push",
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := config.NewGithubConfig(configDir, fileName)
			if err != nil {
				log.Fatal(err)
			}
			extraEnv, _ := cmd.Flags().GetString("extra-env")
			if extraEnv != "" {
				err = cfg.LoadExtraEnv(extraEnv)
				if err != nil {
					log.Fatal(err)
				}
			}
			message, _ = cmd.Flags().GetString("message")
			ctx := context.Background()
			response, err := orchestration.CommitAndPush(ctx, cfg, cfg.TargetDir, message)
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
		Run: func(cmd *cobra.Command, args []string) {
			ghCfg, err := config.NewGithubConfig(configDir, fileName)
			if err != nil {
				log.Fatal(err)
			}

			ctx := context.Background()
			if chart {
				chartCmt, commitErr := orchestration.CommitAndPush(ctx, ghCfg, "dist/kplus", "")
				if commitErr != nil {
					log.Fatal(commitErr)
				}
				fmt.Printf("Commit Created: %+v", chartCmt)
				return
			}
			if resource != "" && service == "" {
				log.Fatal("You need to specify a service for the resource")
			}

			if service != "" && resource != "" {
				cmtMsg := fmt.Sprintf("plumage manifests - %s", time.Now().String())
				resourceCmt, commitErr := orchestration.CommitAndPushResource(ctx, ghCfg, cfg.OutputDir, service, resource, cmtMsg)
				if commitErr != nil {
					log.Fatal(err)
				}
				fmt.Printf("Commit Created: %+v", resourceCmt)
				return
			}

			if service != "" {
				serviceCmt, commitErr := orchestration.CommitAndPushService(ctx, ghCfg, cfg.OutputDir, service, "")
				if commitErr != nil {
					log.Fatal(err)
				}
				fmt.Printf("Commit Created: %+v", serviceCmt)
			}

		},
	}

	cmd.Flags().BoolVarP(&chart, "chart", "c", false, "commit synthesised chart")
	cmd.Flags().StringVarP(&service, "service", "s", "", "commit synthesised service")
	cmd.Flags().StringVarP(&resource, "resource", "r", "", "commit synthesised resource")
	cmd.Flags().StringVarP(&envFiles, "env-files", "e", "", "comma seperated environment files")

	cmd.MarkFlagsMutuallyExclusive("chart", "service")
	cmd.MarkFlagsMutuallyExclusive("chart", "resource")
	return cmd
}
