package cmd

import (
	"fmt"
	"github.com/maliciousbucket/plumage/internal/argocd/commands"
	kubeCmd "github.com/maliciousbucket/plumage/internal/kubeclient/commands"
	orchestrationCmds "github.com/maliciousbucket/plumage/internal/orchestration/commands"
	"github.com/maliciousbucket/plumage/pkg/config"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "plumage",
		Short: "A tool for parsing templates for the Galah Testing Suite",
		Long: `Plumage is a tool for parsing templates for the Galah Testing Suite.
It generates the necessary Kubernetes manifests and chaos test configurations.`,
		// Run: func(cmd *cobra.Command, args []string) { },
	}
	appCfg *config.AppConfig
)

type rootCommand struct {
	cmd *cobra.Command
}

func newRootCommand(cfg *config.AppConfig) *rootCommand {
	appCfg = cfg
	c := &rootCommand{}
	rt := rootCmd

	rt.AddCommand(testComand(cfg))
	rt.AddCommand(configCmd(cfg))
	rt.AddCommand(loadCmd())
	rt.AddCommand(synthCommand())
	rt.AddCommand(orchestrationCmds.CommitPushCmd(cfg.ConfigDir, "github.yaml", cfg))
	//rt.AddCommand(commands.ArgoProjectCmd())
	rt.AddCommand(orchestrationCmds.ProjectCmd())
	rt.AddCommand(commands.ArgoApplicationCmd())

	rt.AddCommand(kubeCmd.ServiceCmd())
	rt.AddCommand(kubeCmd.WaitRelatedPodsCmd())
	rt.AddCommand(orchestrationCmds.SetArgoTokenCmd())
	rt.AddCommand(orchestrationCmds.WatchCmd(cfg.UserConfig.TemplateConfig.TemplateFile))
	rt.AddCommand(orchestrationCmds.ExposeCmd())

	c.cmd = rt
	return c
}

func (c *rootCommand) execute() {
	if err := c.cmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(cfg *config.AppConfig) {
	command := newRootCommand(cfg)
	command.execute()
	//err := rootCmd.Execute()
	//if err != nil {
	//	os.Exit(1)
	//}
}

func init() {

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.plumage.yaml)")

}
