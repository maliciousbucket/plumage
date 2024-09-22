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
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			_, err := newClient()
			return err
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if get {

			}
			return nil
		},
	}
	cmd.Flags().BoolVarP(&get, "get", "g", false, "get cluster")
	cmd.Flags().BoolVarP(&create, "create", "c", false, "create cluster")
	cmd.MarkFlagsMutuallyExclusive("get", "create")
	cmd.MarkFlagsOneRequired("get", "create")
	return cmd
}
