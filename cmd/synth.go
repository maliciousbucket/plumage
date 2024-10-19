package cmd

import (
	"fmt"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/maliciousbucket/plumage/pkg/chaos"
	"github.com/maliciousbucket/plumage/pkg/config"
	"github.com/maliciousbucket/plumage/pkg/kplus"
	"github.com/spf13/cobra"
	"log"
)

func synthCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "synth",
		Short: "Synth Kubernetes manifests",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("use plumage synth <template/service/gateway/tests>")
		},
	}
	cmd.AddCommand(synthGatewayCmd(appCfg.OutputDir, appCfg.Namespace))
	cmd.AddCommand(synthTemplateCommand(appCfg.UserConfig.TemplateConfig.TemplateFile, appCfg.OutputDir, appCfg.MonitoringConfig.Collectors))
	cmd.AddCommand(synthServiceCmd(appCfg.OutputDir, appCfg.Namespace))
	alloyAddress := fmt.Sprintf("%s:%d", appCfg.MonitoringConfig.AlloyAddress, appCfg.MonitoringConfig.Collectors.OtlpGRPCPort)
	cmd.AddCommand(synthChaosCommand(appCfg.OutputDir, appCfg.Namespace, alloyAddress))
	return cmd

}

func synthProvidedCmd(cfg *config.AppConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use: "provided",
	}
	return cmd
}

func synthGatewayCmd(outputDir string, namespace string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gateway",
		Short: "Synth gateway manifests",
		Run: func(cmd *cobra.Command, args []string) {
			out := fmt.Sprintf("%s/ingress/traefik", outputDir)

			app := cdk8s.NewApp(&cdk8s.AppProps{
				Outdir:                  jsii.String(out),
				OutputFileExtension:     nil,
				RecordConstructMetadata: jsii.Bool(false),
				Resolvers:               nil,
				YamlOutputType:          cdk8s.YamlOutputType_FILE_PER_CHART,
			})

			err := kplus.SynthGateway(outputDir, namespace)
			if err != nil {
				log.Fatal(err)
			}

			//kubernetes.NewTraefikIngress(app, "traefik-chart", namespace)

			app.Synth()
		},
	}
	return cmd
}

func synthTemplateCommand(file string, outputDir string, c *config.CollectorConfig) *cobra.Command {
	var templateFile string
	cmd := &cobra.Command{
		Use:   "template",
		Short: "Synth template manifests",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println(templateFile)
			if templateFile == "" {
				templateFile = file
			}
			log.Println("synthesizing manifests from " + templateFile)
			m := c.ToStringMap()
			err := kplus.SynthTemplate(templateFile, outputDir, m)
			if err != nil {
				log.Fatal(err)

			}
		},
	}
	cmd.Flags().StringVar(&templateFile, "template-file", "", "path to template file")
	_ = cmd.MarkFlagFilename("template-file", "yaml", "yml")
	return cmd
}

func synthServiceCmd(file string, outputDir string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "service",
		Short: "Synth manifests for a service",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
				return err
			}
			m := appCfg.MonitoringConfig.Collectors.ToStringMap()

			if err := kplus.SynthService(file, outputDir, args[0], m); err != nil {
				return err
			}

			return nil
		},
	}
	return cmd
}

func synthChaosCommand(outputDir, ns, alloy string) *cobra.Command {
	var outDir string
	var fileName string
	var account string
	cmd := &cobra.Command{
		Use:   "chaos",
		Short: "Synth manifests for a chaos tests",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.ValidateRequiredFlags(); err != nil {
				return err
			}
			out := outputDir
			if outDir != "" {
				out = outDir
			}
			err := chaos.SynthTemplateFile(fileName, out, ns, alloy, account)
			if err != nil {
				log.Fatal(err)
			}
			return nil

		},
	}
	cmd.Flags().StringVar(&fileName, "file", "tests/tests.yaml", "Path to the template file")
	cmd.Flags().StringVar(&outDir, "output-dir", outDir, "Directory where the output is stored")
	cmd.Flags().StringVar(&account, "account", "", "Existing service account name")
	cmd.MarkFlagsOneRequired("file")
	return cmd
}
