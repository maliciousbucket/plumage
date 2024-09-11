package commands

import "github.com/spf13/cobra"

func deploymentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deploy",
		Short: "deployment fun",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	return cmd
}
