package cmd

import (
	"fmt"
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
)

type rootCommand struct {
	cmd *cobra.Command
}

func newRootCommand() *rootCommand {
	c := &rootCommand{}
	rt := rootCmd

	rt.AddCommand(testComand())

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
func Execute() {
	command := newRootCommand()
	command.execute()
	//err := rootCmd.Execute()
	//if err != nil {
	//	os.Exit(1)
	//}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.plumage.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
