package helm

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
)

func InstallArgoCmd(argoVersion string) *cobra.Command {
	var client *helmClient
	var chartVersion string
	cmd := &cobra.Command{
		Use:   "install-argo",
		Short: "Install Argo CD",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			cfg := &ClientCfg{}
			helm, err := newClient(cfg)
			if err != nil {
				return err
			}
			client = helm
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.ParseFlags(args)
			if err != nil {
				return err
			}
			file := cmd.Flag("file").Value.String()
			chartVersion = argoVersion
			version := cmd.Flag("version").Value.String()
			if version != "" {
				chartVersion = version
			}

			ctx := context.Background()

			err = client.InstallArgoChart(ctx, chartVersion, file)
			if err != nil {
				log.Fatal(err)
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

	return cmd
}

func InstallChartFromConfigCmd(cfg *ChartsConfig, configDir string) *cobra.Command {
	var appConfig bool
	var helmConfig string
	var dir string

	cmd := &cobra.Command{
		Use:   "install-chart-cfg",
		Short: "Install Helm Charts",
		Long:  "Install helm charts from config files, local helm charts, or remote repositories",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.ValidateFlagGroups(); err != nil {
				return err
			}
			ctx := context.Background()

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
	cmd.Flags().BoolVarP(&appConfig, "appConfig", "a", false, "Install charts from the AppConfig")
	cmd.Flags().StringVarP(&helmConfig, "helmConfig", "f", "", "Specify a file with Helm Chart Configs")
	cmd.Flags().StringVarP(&dir, "dir", "d", "", "Specify a directory with Helm Configs")

	cmd.MarkFlagsRequiredTogether("helmConfig", "dir")
	cmd.MarkFlagsMutuallyExclusive("helmConfig", "appConfig")
	_ = cmd.MarkPersistentFlagFilename("helmConfig", ".yaml", ".yml")
	_ = cmd.MarkFlagDirname("dir")
	cmd.MarkFlagsOneRequired("appConfig", "helmConfig")
	return cmd
}
