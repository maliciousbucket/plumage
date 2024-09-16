package orchestration

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/maliciousbucket/plumage/internal/kubeclient"
	"github.com/maliciousbucket/plumage/pkg/config"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os/exec"
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
			response, err := CommitAndPush(ctx, cfg, cfg.TargetDir, message)
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

func SetArgoTokenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-argo-token",
		Short: "set argo token",
		Run: func(cmd *cobra.Command, args []string) {
			k8sClient, err := kubeclient.NewClient()
			if err != nil {
				log.Fatal(err)
			}
			ctx := context.Background()
			err = k8sClient.Client.PatchArgoToLB(ctx, "argocd")
			if err != nil {
				log.Fatal(err)
			}
			err = k8sClient.Client.CreateGalahArgoAccount(ctx, "argocd")
			if err != nil {
				log.Fatal(err)
			}
			secret, err := k8sClient.Client.GetArgoPassword(ctx, "argocd")
			if err != nil {
				log.Fatal(err)
			}

			pss := fmt.Sprintf("--password %s", secret)
			loginCommand := exec.Command("argocd", "login", "localhost:8081", "--insecure", "--username", "galah", pss)
			err = loginCommand.Start()
			if err != nil {
				log.Fatal(err)
			}

			tokenCmd := exec.Command("argocd", "account", "generate-token", "--account", "galah")
			stdout, err := tokenCmd.StdoutPipe()

			if err = tokenCmd.Start(); err != nil {
				log.Fatal(err)
			}
			data, err := io.ReadAll(stdout)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Token: %s\n", string(data))

			myEnv, err := godotenv.Read(".env")
			if err != nil {
				log.Fatal(err)
			}
			myEnv["ARGOCD_TOKEN"] = string(data)
			godotenv.Write(myEnv, ".env")

		},
	}
	return cmd
}
