package commands

import (
	"context"
	"fmt"
	"github.com/maliciousbucket/plumage/internal/orchestration"
	"github.com/maliciousbucket/plumage/pkg/config"
	"github.com/spf13/cobra"
	"log"
)

var (
	targetDir   = ""
	targetFiles = ""
	envFiles    = ""
	message     = ""
)

func CommitPushCmd(configDir, fileName string) *cobra.Command {
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
			fmt.Println("YEET")
			fmt.Println(response)

		},
	}
	cmd.Flags().StringVarP(&targetDir, "target-dir", "t", "", "target directory")
	cmd.Flags().StringVarP(&targetFiles, "target-files", "f", "", "comma seperated target files")
	cmd.Flags().StringVarP(&envFiles, "env-files", "e", "", "comma seperated environment files")
	cmd.Flags().StringVarP(&message, "message", "m", "", "commit message")

	err := cmd.MarkFlagRequired("message")
	if err != nil {
		fmt.Printf("Commit Message must be set")
	}

	return cmd
}
