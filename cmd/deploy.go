package cmd

import (
	"github.com/maliciousbucket/plumage/internal/helm"
	orchestrationCmds "github.com/maliciousbucket/plumage/internal/orchestration/commands"
	"github.com/spf13/cobra"
)

var (
	DeployCmd = &cobra.Command{
		Use:   "deploy",
		Short: "deployment commands",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
)

func init() {
	DeployCmd.AddCommand(createArgoTK())
	DeployCmd.AddCommand(helm.InstallArgoCmd())
	DeployCmd.AddCommand(orchestrationCmds.SyncCommand())
	rootCmd.AddCommand(DeployCmd)
}
