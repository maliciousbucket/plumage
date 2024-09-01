package cmd

import (
	"fmt"
	"github.com/maliciousbucket/plumage/pkg/config"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

var (
	targetService = ""
)

func configCmd(cfg *config.AppConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "configuration",
		Long:  "Gets and sets various configuration values",
		Run: func(cmd *cobra.Command, args []string) {
			app, _ := cmd.Flags().GetBool("app")
			monitoring, _ := cmd.Flags().GetBool("monitoring")
			compose, _ := cmd.Flags().GetBool("compose")
			service, _ := cmd.Flags().GetString("service")

			if app {
				cfgYml, err := yaml.Marshal(cfg)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(string(cfgYml))
			}

			if monitoring {
				yml, err := yaml.Marshal(*cfg.MonitoringConfig)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(string(yml))
			}

			if compose {
				if cfg.ProjectDir == "" {
					fmt.Println("project directory not set")

					os.Exit(1)
				}
				if len(cfg.UserConfig.ComposeConfig.ComposeFiles) == 0 {
					fmt.Println("No compose files found")
				}
				var data []byte
				for _, filePath := range cfg.UserConfig.ComposeConfig.ComposeFiles {

					path := filepath.Join(cfg.ProjectDir, filePath)

					content, err := os.ReadFile(path)
					if err != nil {
						fmt.Println(err)
						break
					}
					data = append(data, content...)
				}
				fmt.Println(string(data))
			}

			if service != "" {
				fmt.Println("Not implemented, sorry!")
			}

		},
	}

	cmd.Flags().BoolP("app", "a", false, "View application configuration")
	cmd.Flags().BoolP("monitoring", "m", false, "View monitoring configuration")
	cmd.Flags().BoolP("compose", "c", false, "View target compose configuration")
	cmd.Flags().StringVarP(&targetService, "service", "s", "", "View target service configuration")
	cmd.MarkFlagsMutuallyExclusive("app", "monitoring", "compose", "service")

	cmd.AddCommand(writeConfigCmd())

	return cmd
}

func writeConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "write",
		Short: "write config",
		Long:  "Writes configuration values to files in the config directory",
	}

	cmd.Flags().BoolP("app", "a", false, "Write application configuration")
	cmd.Flags().BoolP("monitoring", "m", false, "Write monitoring configuration")
	cmd.MarkFlagsOneRequired("app", "monitoring")
	cmd.MarkFlagsMutuallyExclusive("app", "monitoring")
	return cmd
}
