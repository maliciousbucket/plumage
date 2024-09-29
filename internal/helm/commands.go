package helm

import (
	"context"
	"fmt"
	"github.com/maliciousbucket/plumage/pkg/config"
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

func InstallChartCmd(appCfg *config.AppConfig) *cobra.Command {
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
				if appCfg.UserConfig != nil && appCfg.UserConfig.ChartConfig != nil {
					err := InstallCRDChartsFromConfig(ctx, nil, appCfg.UserConfig.ChartConfig)
					if err != nil {
						log.Fatal(err)
					}
					return nil
				}
				path := filepath.Join(appCfg.ConfigDir, "helm.yaml")
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

			local, _ := cmd.Flags().GetBool("local")

			values, _ := cmd.Flags().GetString("values")
			namespace, _ := cmd.Flags().GetString("namespace")
			name, _ := cmd.Flags().GetString("name")
			repository, _ := cmd.Flags().GetString("repository")
			version, _ := cmd.Flags().GetString("version")
			crds, _ := cmd.Flags().GetBool("crds")
			replace, _ := cmd.Flags().GetBool("replace")
			lint, _ := cmd.Flags().GetBool("lint")

			skip := false
			upgrade := false
			if crds {
				upgrade = true
			} else {
				skip = true
			}

			cfg := &ChartConfig{
				Repository:  repository,
				Namespace:   namespace,
				Name:        name,
				ReleaseName: name,
				Version:     version,
				Replace:     replace,
				ValuesFiles: []string{
					values,
				},
				Local:       local,
				SkipCRDs:    upgrade,
				UpgradeCRDs: skip,
				Labels:      nil,
				Lint:        lint,
			}
			chartsConf := &ChartsConfig{Charts: []*ChartConfig{cfg}}

			err := InstallCRDChartsFromConfig(ctx, nil, chartsConf)
			if err != nil {
				log.Fatal(err)
			}

			return nil
		},
	}
	cmd.Flags().BoolP("appConfig", "a", false, "Install charts from the AppConfig")
	cmd.Flags().StringP("helmConfig", "h", "", "Specify a file with Helm Chart Configs")
	cmd.Flags().StringP("dir", "d", "", "Specify a directory with Helm Configs")
	cmd.Flags().BoolP("local", "l", false, "Use a local version of the Helm Chart")
	cmd.Flags().StringP("values", "v", "", "Specify a values file for the chart")
	cmd.Flags().StringP("namespace", "n", "", "Specify the namespace for Helm Chart")
	cmd.Flags().StringP("name", "n", "", "Specify the release name of the Helm Chart")
	cmd.Flags().StringP("repository", "r", "", "Specify the local file or repository to pull the helm chart from")
	cmd.Flags().String("version", "", "Specify the version of the Helm Chart")
	cmd.Flags().Bool("crds", false, "Install CRDs")
	cmd.Flags().Bool("replace", false, "Replace existing Helm Chart with the latest version")
	cmd.Flags().Bool("lint", false, "Lint the chart before installing")

	cmd.MarkFlagsRequiredTogether("helmConfig", "dir")
	cmd.MarkFlagsMutuallyExclusive("helmConfig", "appConfig")
	_ = cmd.MarkFlagFilename("values", "values.yaml", "values.yml", ".yaml", ".yml")
	_ = cmd.MarkFlagDirname("dir")
	cmd.MarkFlagsMutuallyExclusive("appConfig", "repository")
	cmd.MarkFlagsRequiredTogether("name", "repository", "local", "version")
	return cmd
}
