package kplus

import (
	"fmt"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/maliciousbucket/plumage/pkg/kubernetes"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"
)

func SynthTemplate(filePath, outputDir string, monitoring map[string]string) error {
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
	if template.Name == "" {
		return fmt.Errorf("no name found in %s", filePath)
	}
	out := fmt.Sprintf("%s/%s", outputDir, template.Name)
	app := cdk8s.NewApp(&cdk8s.AppProps{
		Outdir:                  jsii.String(out),
		OutputFileExtension:     nil,
		RecordConstructMetadata: jsii.Bool(true),
		Resolvers:               nil,
		YamlOutputType:          cdk8s.YamlOutputType_FOLDER_PER_CHART_FILE_PER_RESOURCE,
	})

	namespace := "galah-testbed"
	if template.Namespace != "" {
		namespace = template.Namespace
	}

	ProjectNamespace(app, namespace)
	app.Synth()
	i := 30100
	for _, service := range template.Services {
		log.Println("Generating manifests for " + service.Name)
		serviceOut := fmt.Sprintf("%s/%s", out, service.Name)
		serviceApp := cdk8s.NewApp(&cdk8s.AppProps{
			Outdir:                  jsii.String(serviceOut),
			OutputFileExtension:     nil,
			RecordConstructMetadata: jsii.Bool(false),
			Resolvers:               nil,
			YamlOutputType:          cdk8s.YamlOutputType_FILE_PER_CHART,
		})
		name := fmt.Sprintf("%s", service.Name)
		chart := cdk8s.NewChart(serviceApp, jsii.String(name), &cdk8s.ChartProps{
			DisableResourceNameHashes: jsii.Bool(true),
			Labels:                    nil,
			Namespace:                 jsii.String(namespace),
		})

		NewServiceManifests(chart, service.Name, namespace, &service, monitoring, i)
		serviceApp.Synth()
		i++

	}

	return nil
}

func GetServices(filePath string) ([]string, string, string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, "", "", err
	}

	data, err := io.ReadAll(file)

	var template Template
	err = yaml.Unmarshal(data, &template)
	if len(template.Services) == 0 {
		return nil, "", "", fmt.Errorf("no services found in %s", filePath)
	}

	if template.Name == "" {
		return nil, "", "", fmt.Errorf("no name found in %s", filePath)
	}

	if template.Namespace == "" {
		return nil, "", "", fmt.Errorf("no namespace found in %s", filePath)
	}

	template.applyHost()

	var services []string
	for _, service := range template.Services {
		services = append(services, service.Name)
	}
	return services, template.Namespace, template.Name, nil
}

func SynthService(filePath, outputDir, service string, monitoring map[string]string) error {
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

	var serviceTemplate ServiceTemplate

	for _, svc := range template.Services {
		if svc.Name == service {
			serviceTemplate = svc
		}
	}

	if serviceTemplate.Name == "" {
		return fmt.Errorf("unable to find service: %s in template", service)
	}

	namespace := "galah-testbed"
	if template.Namespace != "" {
		namespace = template.Namespace
	}

	out := fmt.Sprintf("%s/%s", outputDir, template.Name)
	app := cdk8s.NewApp(&cdk8s.AppProps{
		Outdir:                  jsii.String(out),
		OutputFileExtension:     nil,
		RecordConstructMetadata: jsii.Bool(true),
		Resolvers:               nil,
		YamlOutputType:          cdk8s.YamlOutputType_FILE_PER_CHART,
	})

	ProjectNamespace(app, namespace)
	chart := cdk8s.NewChart(app, jsii.String(serviceTemplate.Name), &cdk8s.ChartProps{
		DisableResourceNameHashes: jsii.Bool(true),
		Labels:                    nil,
		Namespace:                 jsii.String(namespace),
	})

	NewServiceManifests(chart, serviceTemplate.Name, namespace, &serviceTemplate, monitoring, 31020)

	app.Synth()

	log.Printf("\n Manifests for %s have been created at %s/%s", service, out, service)

	return nil
}

func SynthGateway(outputDir, ns string) error {
	if _, err := os.Stat(outputDir); err != nil {
		return fmt.Errorf("output directory %s does not exist", outputDir)
	}

	out := fmt.Sprintf("%s/ingress/traefik", outputDir)

	app := cdk8s.NewApp(&cdk8s.AppProps{
		Outdir:                  jsii.String(out),
		OutputFileExtension:     nil,
		RecordConstructMetadata: jsii.Bool(false),
		Resolvers:               nil,
		YamlOutputType:          cdk8s.YamlOutputType_FILE_PER_CHART,
	})

	kubernetes.NewTraefikIngress(app, "traefik-chart", ns)
	//ProjectNamespace(app, ns)
	app.Synth()
	return nil
}

func (t *Template) applyHost() {
	if t.Host != "" {
		for _, service := range t.Services {
			service.Host = t.Host
		}
	}
}

func loadTemplate(filePath string) (*Template, error) {
	return nil, nil
}
