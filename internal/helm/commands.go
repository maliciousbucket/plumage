package helm

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
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
			//namespace := cmd.Flag("namespace").Value.String()

			chart := cmd.Flag("chart").Value.String()
			version := cmd.Flag("version").Value.String()
			repo := cmd.Flag("remote").Value.String()

			err = installArgo("argocd", file, version, chart, repo)
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

func installArgo(ns, file, ver, chart, remote string) error {
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
	err := InstallArgo(ctx, clientCfg, ns, opts...)
	if err != nil {
		return err
	}
	return nil
}

func InstallChartCmd(cfg *ChartsConfig, configDir string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install-chart",
		Short: "Install Helm Charts",
		Long:  "Install helm charts from config files, local helm charts, or remote repositories",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.ValidateFlagGroups(); err != nil {
				return err
			}
			ctx := context.Background()
			appConfig, _ := cmd.Flags().GetBool("appConfig")
			if appConfig {
				if cfg != nil {
					err := InstallCRDChartsFromConfig(ctx, nil, cfg)
					if err != nil {
						log.Fatal(err)
					}
					return nil
				}
				path := filepath.Join(configDir, "helm.yaml")
				configFile, err := os.Stat(path)
				if err != nil {
					return err
				}
				if configFile.IsDir() {
					return fmt.Errorf("%s is a directory, not a file", path)
				}
				return fmt.Errorf("unable to build helm config from given app config")
			}
			dir, _ := cmd.Flags().GetString("dir")
			helmConfig, _ := cmd.Flags().GetString("helmConfig")
			if helmConfig != "" {
				if dir != "" {

					err := InstallCRDChartsFromFile(ctx, nil, dir, helmConfig)
					if err != nil {
						log.Fatal(err)
					}
					return nil
				}
				return fmt.Errorf("must specify both a dir and file together")
			}

			return nil
		},
	}
	cmd.Flags().BoolP("appConfig", "a", false, "Install charts from the AppConfig")
	cmd.Flags().StringP("helmConfig", "h", "", "Specify a file with Helm Chart Configs")
	cmd.Flags().StringP("dir", "d", "", "Specify a directory with Helm Configs")

	cmd.MarkFlagsRequiredTogether("helmConfig", "dir")
	cmd.MarkFlagsMutuallyExclusive("helmConfig", "appConfig")
	_ = cmd.MarkPersistentFlagFilename("helmConfig", ".yaml", ".yml")
	_ = cmd.MarkFlagDirname("dir")
	return cmd
}
