package chaos

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Template struct {
	Namespace string
	ScriptDir string           `yaml:"scriptDir"`
	Scripts   []ScriptTemplate `yaml:"scripts"`
	LibDir    string           `yaml:"libDir"`
	LibFiles  []string         `yaml:"libFiles"`
}

func loadTemplate(file, namespace string) (*Template, error) {
	info, err := os.Stat(file)
	if err != nil {
		return nil, err
	}
	if info.IsDir() {
		return nil, fmt.Errorf("%s is a directory", file)
	}
	template := &Template{}
	bytes, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(bytes, template); err != nil {
		return nil, err
	}
	template.Namespace = namespace
	return template, nil
}

type ScriptTemplate struct {
	Name             string            `yaml:"name"`
	ScriptName       string            `yaml:"script"`
	LibFiles         []string          `yaml:"libFiles"`
	ResourceRequests ScriptResources   `yaml:"resourceRequests"`
	ResourceLimits   ScriptResources   `yaml:"resourceLimits"`
	Args             []string          `yaml:"args"`
	Env              map[string]string `yaml:"env"`
	RunOnce          bool              `yaml:"runOnce"`
	Schedule         string            `yaml:"schedule"`
	Labels           map[string]string `yaml:"labels"`
	Annotations      map[string]string `yaml:"annotations"`
}

type ScriptResources struct {
	Cpu    *ScriptResource
	Memory *ScriptResource
}

type ScriptResource struct {
	RequestMillis int
	LimitMillis   int
}
