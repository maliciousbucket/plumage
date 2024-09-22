package commands

import "github.com/spf13/cobra"

var (
	get    = false
	create = false
)

func ClusterCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cluster",
		Short: "manage argo's clusters",
		Run: func(cmd *cobra.Command, args []string) {
			if get {

			}
		},
	}
	cmd.Flags().BoolVarP(&get, "get", "g", false, "get cluster")
	cmd.Flags().BoolVarP(&create, "create", "c", false, "create cluster")
	cmd.MarkFlagsMutuallyExclusive("get", "create")
	cmd.MarkFlagsOneRequired("get", "create")
	return cmd
}
