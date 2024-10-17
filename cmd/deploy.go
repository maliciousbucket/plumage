package cmd

import (
	"fmt"
	argoCmds "github.com/maliciousbucket/plumage/internal/argocd/commands"
	"github.com/maliciousbucket/plumage/internal/helm"
	orchestrationCmds "github.com/maliciousbucket/plumage/internal/orchestration/commands"
	"github.com/maliciousbucket/plumage/pkg/config"
	"github.com/spf13/cobra"
)

func deployCmd(cfg *config.AppConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deploy",
		Short: "Deployment commands",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(cmd.UsageString())
		},
	}
	cmd.AddCommand(helm.InstallArgoCmd(cfg.UserConfig.ChartConfig.ArgoVersion))
	cmd.AddCommand(orchestrationCmds.SyncCommand())
	cmd.AddCommand(argoCmds.ClusterCommand())
	cmd.AddCommand(orchestrationCmds.DeployAppCmd("testdata/chirp/template.yaml", cfg.Namespace))
	cmd.AddCommand(orchestrationCmds.ArgoAuthCmd())
	cmd.AddCommand(orchestrationCmds.DeployMonitoringCommand())
	cmd.AddCommand(orchestrationCmds.DeployTemplateCommand(cfg))
	cmd.AddCommand(orchestrationCmds.DeployGatewayCommand(cfg.ConfigDir, cfg.OutputDir, cfg.Namespace))
	cmd.AddCommand(orchestrationCmds.DeployChaosCommand(cfg.OutputDir, cfg.Namespace, cfg.UserConfig.ChartConfig.K6Version))
	//cmd.AddCommand(helm.InstallChartCmd(cfg.UserConfig.ChartConfig.Charts, cfg.ConfigDir))

	return cmd
}
