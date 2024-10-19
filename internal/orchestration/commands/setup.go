package commands

import (
	"context"
	"github.com/maliciousbucket/plumage/internal/orchestration"
	"github.com/spf13/cobra"
	"log"
)

func SetupCmd(ns, argoVersion, promVersion, valuesFile string) *cobra.Command {
	var namespace string
	var envFile string
	var version string
	var values string
	cmd := &cobra.Command{
		Use:   "setup",
		Short: "Setup ArgoCD and Namespaces",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := newHelmClient(); err != nil {
				return err
			}
			if err := newKubeClient(); err != nil {
				return err
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			if namespace == "" {
				namespace = ns
			}
			if version == "" {
				version = argoVersion
			}
			if values == "" {
				values = valuesFile
			}
			ctx := context.Background()
			if err := orchestration.Setup(ctx, helmClient, kubernetesClient, namespace, version, promVersion, values, envFile); err != nil {
				log.Fatalln(err)
			}
		},
	}
	cmd.Flags().StringVar(&namespace, "namespace", "", "the namespace to deploy services to")
	cmd.Flags().StringVar(&envFile, "env", "", "path to file with environment variables")
	cmd.Flags().StringVar(&values, "values", "", "path to argocd values file")
	cmd.Flags().StringVar(&version, "version", version, "ArgoCD version to use")
	return cmd
}
