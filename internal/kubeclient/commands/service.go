package commands

import (
	"context"
	"fmt"
	"github.com/maliciousbucket/plumage/internal/kubeclient"
	"github.com/spf13/cobra"
	"log"
)

func ServiceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "service",
		Short: "Watch Kubernetes services",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.ValidateRequiredFlags(); err != nil {
				return err
			}
			service, _ := cmd.Flags().GetString("service")
			if service == "" {
				return fmt.Errorf("please specify a service")
			}
			namespace, _ := cmd.Flags().GetString("namespace")
			if namespace == "" {
				return fmt.Errorf("please specify a namespace")
			}
			client, err := kubeclient.NewClient()
			if err != nil {
				return err
			}
			ctx := context.Background()
			err = client.WaitServicePods(ctx, namespace, service)
			if err != nil {
				return err
			}
			fmt.Println("service is ready")
			return nil
		},
	}
	cmd.Flags().StringP("service", "s", "", "Service name")
	cmd.Flags().StringP("namespace", "n", "", "Namespace to watch")
	_ = cmd.MarkFlagRequired("namespace")
	err := cmd.MarkFlagRequired("service")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return cmd
}

// WaitRelatedPodsCmd TODO: FIx
func WaitRelatedPodsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "wait-related-pods",
		Short: "Wait for related pods to be ready",
		Run: func(cmd *cobra.Command, args []string) {
			namespace, _ := cmd.Flags().GetString("namespace")
			if namespace == "" {
				log.Fatalf("please specify a namespace")

			}
			client, err := kubeclient.NewClient()
			if err != nil {
				log.Fatal(err)
			}
			ctx := context.Background()
			err = client.WaitAllArgoPods(ctx, namespace)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
	cmd.Flags().StringP("namespace", "n", "", "namespace")
	err := cmd.MarkFlagRequired("namespace")
	if err != nil {
		fmt.Println(err)
	}
	return cmd
}
