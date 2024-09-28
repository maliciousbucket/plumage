package cmd

import (
	"fmt"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/maliciousbucket/plumage/pkg/config"
	"github.com/maliciousbucket/plumage/pkg/kplus"
	"github.com/maliciousbucket/plumage/pkg/kubernetes"
	"github.com/spf13/cobra"
	"log"
)

const (
	//tmplFile = "testdata/kplus/template.yaml"
	tmpFile = "testdata/manifests/gcs.yaml"
)

func synthCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "synth",
		Short: "Synth Kubernetes manifests",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("use plumage synth <template/service/gateway>")
		},
	}
	cmd.AddCommand(synthGatewayCmd(appCfg.OutputDir, appCfg.Namespace))
	cmd.AddCommand(synthTemplateCommand("testdata/chirp/template.yaml", appCfg.OutputDir, appCfg.MonitoringConfig.Collectors))
	cmd.AddCommand(synthServiceCmd(appCfg.OutputDir, appCfg.Namespace))
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

			kubernetes.NewTraefikIngress(app, "traefik-chart", namespace)

			app.Synth()
		},
	}
	return cmd
}

func synthTemplateCommand(file string, outputDir string, c *config.CollectorConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "template",
		Short: "Synth template manifests",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("synthesizing manifests from " + file)
			m := c.ToStringMap()
			err := kplus.SynthTemplate(file, outputDir, m)
			if err != nil {
				log.Fatal(err)

			}
		},
	}
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
