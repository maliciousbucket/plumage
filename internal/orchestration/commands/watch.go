package commands

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
)

func WatchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "watch",
		Short: "Watch Kubernetes resources",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return newKubeClient()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.ValidateRequiredFlags(); err != nil {
				return err
			}
			if err := cmd.ValidateFlagGroups(); err != nil {
				return err
			}
			name, _ := cmd.Flags().GetString("name")
			ns, _ := cmd.Flags().GetString("namespace")
			svc, _ := cmd.Flags().GetBool("service")
			pod, _ := cmd.Flags().GetBool("pod")
			deployment, _ := cmd.Flags().GetBool("deployment")

			if name == "" {
				return fmt.Errorf("must specify name")
			}
			ctx := context.Background()
			if deployment {

				return kubernetesClient.WatchDeployment(ctx, ns, name)
			}

			if svc {

				return kubernetesClient.WaitPodNameRunning(ctx, ns, name)
			}

			if pod {
				return kubernetesClient.WaitPodInstanceRunning(ctx, ns, name)
			}

			return nil
		},
	}
	cmd.Flags().BoolP("service", "s", false, "Watch a service")
	cmd.Flags().BoolP("deployment", "d", false, "Watch a deployment")
	cmd.Flags().BoolP("pod", "p", false, "Watch a pod instance")
	cmd.Flags().StringP("namespace", "n", "default", "Namespace of the resource")
	cmd.Flags().StringP("name", "r", "", "Name of the resource")

	cmd.MarkFlagsMutuallyExclusive("service", "deployment", "pod")
	cmd.MarkFlagsOneRequired("service", "deployment", "pod")
	_ = cmd.MarkFlagRequired("name")

	return cmd
}
