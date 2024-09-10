package helm

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
)

func InstallArgoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install-argo",
		Short: "Install Argo CD",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.ParseFlags(args)
			if err != nil {
				return err
			}
			file := cmd.Flag("file").Value.String()
			namespace := cmd.Flag("namespace").Value.String()

			chart := cmd.Flag("chart").Value.String()
			version := cmd.Flag("version").Value.String()
			repo := cmd.Flag("remote").Value.String()

			err = InstallArgo(namespace, file, version, chart, repo)
			if err != nil {
				return err
			}

			return nil
		},
	}
	cmd.Flags().StringP("file", "f", "", "Specify a values file for the ArgoCD helm chart")
	cmd.Flags().StringP("namespace", "n", "", "Specify the namespace for ArgoCD")
	cmd.Flags().BoolP("local", "l", false, "Use a local version of the ArgoCD chart")
	cmd.Flags().StringP("chart", "c", "", "The relative path to the local ArgoCD chart")
	cmd.MarkFlagsRequiredTogether("chart", "local")
	cmd.Flags().StringP("remote", "r", "", "Specfy the repo to pull the chart from")
	cmd.MarkFlagsMutuallyExclusive("remote", "local")
	cmd.Flags().StringP("version", "v", "", "Specify the version of the ArgoCD chart")

	//ht := cmd.Flags().Args()
	//cmd.Flags().Parse()
	return cmd
}

func InstallArgo(ns, file, ver, chart, remote string) error {
	argo := defaultArgoConfig(ns)
	var opts []ArgoOpts
	if ver != "" {
		opts = append(opts, WithVersion(ver))
	}
	if chart != "" {
		opts = append(opts, WithLocalChart(chart))
	}
	if remote != "" {
		opts = append(opts, FromRemote(remote))
	}
	if file != "" {
		opts = append(opts, WithValuesFile(file))
	}

	ctx := context.Background()

	clientCfg := &ClientCfg{
		Namespace:        "argocd",
		RepositoryConfig: "",
		RepositoryCache:  "",
		RegistryConfig:   "",
		KubeCfgPath:      "",
		username:         "123",
		password:         "abc",
	}
	client, err := New(clientCfg)
	if err != nil {
		return err
	}

	res, err := client.InstallArgo(ctx, argo, opts...)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil

}
