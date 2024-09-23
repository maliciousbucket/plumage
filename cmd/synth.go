package cmd

import (
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
		Short: "synth",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	cmd.AddCommand(synthGatewayCmd(appCfg.ConfigDir, appCfg.Namespace))
	cmd.AddCommand(synthTemplateCommand("", appCfg.OutputDir, appCfg.MonitoringConfig.Collectors))
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
		Short: "synth gateway manifests",
		Run: func(cmd *cobra.Command, args []string) {

			app := cdk8s.NewApp(&cdk8s.AppProps{
				Outdir:                  jsii.String(outputDir),
				OutputFileExtension:     nil,
				RecordConstructMetadata: jsii.Bool(true),
				Resolvers:               nil,
				YamlOutputType:          cdk8s.YamlOutputType_FOLDER_PER_CHART_FILE_PER_RESOURCE,
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
		Short: "synth template manifests",
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
