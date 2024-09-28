package cmd

import (
	argoCmds "github.com/maliciousbucket/plumage/internal/argocd/commands"
	"github.com/maliciousbucket/plumage/internal/helm"
	orchestrationCmds "github.com/maliciousbucket/plumage/internal/orchestration/commands"
	"github.com/maliciousbucket/plumage/pkg/config"
	"github.com/spf13/cobra"
	"log"
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
	DeployCmd.AddCommand(argoCmds.ClusterCommand())
	DeployCmd.AddCommand(orchestrationCmds.DeployAppCmd("testdata/chirp/template.yaml"))
	DeployCmd.AddCommand(orchestrationCmds.ArgoAuthCmd())
	if appCfg == nil {
		cfg, err := config.NewAppConfig("./config/compose", "galah-testbed", "", "")
		if err != nil {
			log.Fatal(err)
		}
		appCfg = cfg
	}
	DeployCmd.AddCommand(orchestrationCmds.DeployTemplateCommand(appCfg))
	DeployCmd.AddCommand(orchestrationCmds.DeployGatewayCommand(appCfg.ConfigDir, appCfg.OutputDir, appCfg.Namespace))
	DeployCmd.AddCommand(orchestrationCmds.DeployMonitoringCommand())
	rootCmd.AddCommand(DeployCmd)
}
