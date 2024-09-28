package commands

import (
	"context"
	"fmt"
	"github.com/maliciousbucket/plumage/pkg/kplus"
	"github.com/spf13/cobra"
	"log"
	"time"
)

var (
	watchDuration time.Duration
)

func WatchCmd(template string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "watch",
		Short: "Watch Kubernetes resources",
		Args: func(cmd *cobra.Command, args []string) error {
			if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
				return err
			}
			if args[0] == "" {
				return fmt.Errorf("resource name is required")
			}
			return nil
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return newKubeClient()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.ValidateRequiredFlags(); err != nil {
				return err
			}
			if err := cmd.ValidateFlagGroups(); err != nil {
				return err
			}
			name := args[0]
			ns, _ := cmd.Flags().GetString("namespace")
			svc, _ := cmd.Flags().GetBool("service")
			pod, _ := cmd.Flags().GetBool("pod")
			deployment, _ := cmd.Flags().GetBool("deployment")
			allPods, _ := cmd.Flags().GetBool("allPods")

			ctx := context.Background()
			if deployment {

				if err := kubernetesClient.WatchDeployment(ctx, ns, name); err != nil {
					log.Fatalf("failed to watch deployment %s/%s: %v", ns, name, err)
				}
			}

			if svc {

				if err := kubernetesClient.WaitPodNameRunning(ctx, ns, name); err != nil {
					log.Fatalf("failed to watch service %s/%s: %v", ns, name, err)
				}
			}

			if pod {
				if err := kubernetesClient.WaitPodInstanceRunning(ctx, ns, name); err != nil {
					log.Fatalf("failed to watch pod instance %s/%s: %v", ns, name, err)
				}
			}
			if allPods {

				if err := kubernetesClient.WaitAppPods(ctx, ns, name, watchDuration); err != nil {
					log.Fatalf("failed to watch all pods %s/%s: %v", ns, name, err)
				}
			}

			return nil
		},
	}
	cmd.Flags().BoolP("service", "s", false, "Watch a service")
	cmd.Flags().BoolP("deployment", "d", false, "Watch a deployment")
	cmd.Flags().BoolP("pod", "p", false, "Watch a pod instance")
	cmd.Flags().BoolP("allPods", "a", false, "Watch all pods for a service")
	cmd.Flags().StringP("namespace", "n", "default", "Namespace of the resource")
	cmd.PersistentFlags().DurationVarP(&watchDuration, "timeout", "t", time.Second*30, "Timeout in seconds")

	cmd.MarkFlagsMutuallyExclusive("service", "deployment", "pod", "allPods")
	cmd.MarkFlagsOneRequired("service", "deployment", "pod", "allPods")
	_ = cmd.MarkFlagRequired("name")
	cmd.AddCommand(watchTemplateDeploymentCmd(template))
	return cmd
}

func watchTemplateDeploymentCmd(file string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "template",
		Short: "Watch the template's deployment",
		Run: func(cmd *cobra.Command, args []string) {
			template, _ := cmd.Flags().GetString("file")
			if template == "" {
				template = file
			}
			services, ns, appName, err := kplus.GetServices(template)
			if err != nil {
				log.Fatal(err)
			}
			ctx := context.Background()
			if err = kubernetesClient.WatchAppDeployment(ctx, ns, services); err != nil {
				log.Fatalf("failed to watch deployment %s/%s: %v", ns, appName, err)
			}
		},
	}
	cmd.Flags().StringP("file", "f", "", "Template file")
	return cmd
}
