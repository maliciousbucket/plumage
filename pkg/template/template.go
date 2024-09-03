package template

import (
	"context"
	"fmt"
	types2 "github.com/compose-spec/compose-go/v2/types"
	"github.com/maliciousbucket/plumage/pkg/types"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"

	"github.com/compose-spec/compose-go/v2/cli"
)

type GalahTemplate struct {
	HostAddress string             `yaml:"hostAddress"`
	Services    []*ServiceTemplate `yaml:"services"`
	Containers  []*types.ContainerSpec
}

type ServiceConfig struct {
	Name      string               `yaml:"name"`
	Namespace string               `yaml:"namespace"`
	Service   *ServiceTemplate     `yaml:"service"`
	Container *types.ContainerSpec `yaml:"container"`
}

type GalahTemplateOptions struct {
	ProjectName  string
	ProjectDir   string
	ComposeFiles []string
	TemplateFile string
	EnvFile      string
}

func LoadProject(opts *GalahTemplateOptions) (*GalahTemplate, error) {
	composeProject, err := LoadComposeFiles(opts)
	if err != nil {
		return nil, err
	}
	var containers []*types.ContainerSpec

	for _, composeSvc := range composeProject.Services {
		container, err := types.ParseServiceConfig(composeSvc)
		if err != nil {
			return nil, err
		}
		containers = append(containers, container)
	}

	template, err := loadGalahTemplate(opts)
	if err != nil {
		return nil, err
	}
	template.Containers = containers

	return template, nil
}

func loadGalahTemplate(opts *GalahTemplateOptions) (*GalahTemplate, error) {
	path := filepath.Join(opts.ProjectDir, opts.TemplateFile)
	fmt.Println(path)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var template GalahTemplate
	err = yaml.Unmarshal(content, &template)
	if err != nil {
		return nil, err
	}

	return &template, nil
}

func LoadComposeFiles(opts *GalahTemplateOptions) (*types2.Project, error) {
	if len(opts.ComposeFiles) == 0 {
		return nil, fmt.Errorf("no compose files specified")
	}

	ctx := context.Background()
	options, err := cli.NewProjectOptions(
		opts.ComposeFiles,
		cli.WithName(opts.ProjectName),
		cli.WithEnvFiles(opts.EnvFile),
		cli.WithWorkingDirectory(opts.ProjectDir),
	)
	if err != nil {
		return nil, err
	}

	project, err := options.LoadProject(ctx)
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (t *GalahTemplate) listServices() map[string]*ServiceTemplate {
	services := map[string]*ServiceTemplate{}
	for _, s := range t.Services {
		services[s.Name] = s
	}
	return services
}

func (t *GalahTemplate) listContainers() map[string]*types.ContainerSpec {
	containers := map[string]*types.ContainerSpec{}
	for _, s := range t.Containers {
		containers[s.Name] = s
	}
	return containers
}

func (t *GalahTemplate) ServiceConfigs(namesSpace string) ([]*ServiceConfig, error) {
	containers := t.listContainers()
	if len(containers) == 0 {
		return nil, fmt.Errorf("no containers found")
	}
	services := t.listServices()

	var configs []*ServiceConfig
	for _, service := range services {
		if container, ok := containers[service.Name]; ok {
			configs = append(configs, &ServiceConfig{
				Name:      service.Name,
				Namespace: namesSpace,
				Service:   service,
				Container: container,
			})
		} else {
			return nil, fmt.Errorf("no container for service found. Service: %s", service.Name)
		}
	}
	return configs, nil
}

func loadComposeFile(projectDir string, composeFilePath string) ([]*types.ContainerSpec, error) {
	return nil, nil
}

func loadServiceTemplate(projectDir string, file string) (*ServiceTemplate, error) {
	return nil, nil
}
