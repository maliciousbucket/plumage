package kplus

import (
	"fmt"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

func SynthTemplate(filePath, outputDir string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	data, err := io.ReadAll(file)

	var template Template
	err = yaml.Unmarshal(data, &template)

	if len(template.Services) == 0 {
		return fmt.Errorf("no services found in %s", filePath)
	}

	app := cdk8s.NewApp(&cdk8s.AppProps{
		Outdir:                  jsii.String(outputDir),
		OutputFileExtension:     nil,
		RecordConstructMetadata: jsii.Bool(true),
		Resolvers:               nil,
		YamlOutputType:          cdk8s.YamlOutputType_FOLDER_PER_CHART_FILE_PER_RESOURCE,
	})

	for _, service := range template.Services {
		fmt.Println(service.Name)
		fmt.Println(service)
		name := fmt.Sprintf("%s-chart", service.Name)
		chart := cdk8s.NewChart(app, jsii.String(name), &cdk8s.ChartProps{
			DisableResourceNameHashes: jsii.Bool(true),
			Labels:                    nil,
			Namespace:                 jsii.String(template.Namespace),
		})

		NewServiceManifests(chart, service.Name, &service)
	}

	app.Synth()

	return nil
}

func loadTemplate(filePath string) (*Template, error) {
	return nil, nil
}
