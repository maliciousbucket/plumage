package chaos

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type Template struct {
	Namespace      string
	ScriptDir      string `yaml:"scriptDir"`
	ServiceAccount string
	Scripts        []ScriptTemplate `yaml:"scripts"`
	LibDir         string           `yaml:"libDir"`
	LibFiles       []string         `yaml:"libFiles"`
}

func loadTemplate(configDir, file, namespace string) (*Template, error) {
	path := filepath.Join(configDir, file)
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if info.IsDir() {
		return nil, fmt.Errorf("%s is a directory", path)
	}
	template := &Template{}
	bytes, err := os.ReadFile(path)
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
	Schedule         *JobSchedule      `yaml:"schedule"`
	Labels           map[string]string `yaml:"labels"`
	Annotations      map[string]string `yaml:"annotations"`
	ExistingScript   string            `yaml:"existingScript"`
	ExistingEnv      string            `yaml:"existingEnv"`
	ExistingAccount  string            `yaml:"existingAccount"`
}

type JobSchedule struct {
	Minute     string `yaml:"minute"`
	Hour       string `yaml:"hour"`
	DayOfMonth string `yaml:"dayOfMonth"`
	Month      string `yaml:"month"`
	DayOfWeek  string `yaml:"dayOfWeek"`
}

type ScriptResources struct {
	Cpu    *ScriptResource
	Memory *ScriptResource
}

type ScriptResource struct {
	RequestMillis int
	LimitMillis   int
}
