package template

import (
	"github.com/maliciousbucket/plumage/pkg/resilience"
	"github.com/maliciousbucket/plumage/pkg/types"
)

type ServiceTemplate struct {
	Name       string                      `yaml:"name"`
	Paths      []*Path                     `yaml:"paths"`
	Resources  *ContainerResourcesTemplate `yaml:"resources,omitempty"`
	Resilience *resilience.ResTemplate     `yaml:"resiliencePolicy,omitempty"`
	Monitoring *MonitoringTemplate         `yaml:"monitoring,omitempty"`
	EnvFile    string                      `yaml:"envFile,omitempty"`
}

type Path struct {
	Path string `yaml:"path"`
	Port int    `yaml:"port"`
}

func (s *ServiceTemplate) WebService() (*types.WebService, error) {
	return nil, nil
}
