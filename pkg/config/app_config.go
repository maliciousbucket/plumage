package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/maliciousbucket/plumage/internal/helm"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
)

type AppConfig struct {
	ProjectName      string            `yaml:"project"`
	ConfigDir        string            `yaml:"configDir"`
	ProjectDir       string            `yaml:"projectDir"`
	OutputDir        string            `yaml:"outputDir"`
	Namespace        string            `yaml:"namespace"`
	UserConfig       *UserConfig       `yaml:"userConfig"`
	MonitoringConfig *MonitoringConfig `yaml:"monitoringConfig"`
}

func NewAppConfig(projectDir string, namespace string, outputDir string, monitoringEnv string) (*AppConfig, error) {
	configDir, err := findOrCreateConfigDir("", ".env")
	if err != nil {
		return nil, err
	}
	userConfig, err := loadUserConfigWithDefaults(configDir)
	if err != nil {
		return nil, err
	}
	fmt.Println(userConfig)
	monitoringConfig, err := NewMonitoringConfig(configDir, monitoringEnv)
	if err != nil {
		return nil, err
	}

	outDir := "./dist"
	if outputDir != "" {
		outDir = outputDir
	}

	appConfig := &AppConfig{
		ConfigDir:        configDir,
		ProjectDir:       projectDir,
		Namespace:        namespace,
		UserConfig:       userConfig,
		MonitoringConfig: monitoringConfig,
		OutputDir:        outDir,
	}
	appConfig.OutputDir = userConfig.OutputDir
	appConfig.Namespace = userConfig.Namespace
	return appConfig, nil
}

type UserConfig struct {
	OutputDir      string             `yaml:"outputDir"`
	Namespace      string             `yaml:"namespace"`
	ComposeConfig  ComposeConfig      `yaml:"compose"`
	TemplateConfig TemplateConfig     `yaml:"template"`
	TraefikConfig  TraefikConfig      `yaml:"traefik"`
	ChartConfig    *helm.ChartsConfig `yaml:"helm"`
}

func getDefaultUserConfig() UserConfig {
	composeDir := os.Getenv("PROJECT_DIR")
	if composeDir == "" {
		composeDir = "./config/compose"
	}
	templateDir := os.Getenv("TEMPLATE_DIR")
	if templateDir == "" {
		templateDir = "/templates"
	}
	return UserConfig{
		ComposeConfig: ComposeConfig{
			WorkingDir:   composeDir,
			ComposeFiles: []string{},
			EnvFiles:     []string{},
		},
		TemplateConfig: TemplateConfig{
			WorkingDir:              templateDir,
			TemplateFile:            "galah-template.yaml",
			ServiceTemplateFiles:    []string{},
			MonitoringTemplateFiles: []string{},
			ResilienceTemplateFiles: []string{},
		},
		TraefikConfig: getDefaultTraefikConfig(),
	}
}

func loadUserConfigWithDefaults(configDir string) (*UserConfig, error) {
	config := getDefaultUserConfig()
	return loadUserConfig(configDir, &config)
}

func loadUserConfig(configDir string, base *UserConfig) (*UserConfig, error) {
	fileName := filepath.Join(configDir, "config.yaml")

	if _, err := os.Stat(fileName); err != nil {
		if os.IsNotExist(err) {
			file, err := os.Create(fileName)
			if err != nil {
				return nil, err
			}
			file.Close()
		} else {
			return nil, err
		}
	}

	content, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(content, base); err != nil {
		return nil, err
	}
	return base, nil
}

type TemplateConfig struct {
	WorkingDir              string   `yaml:"workingDir"`
	TemplateFile            string   `yaml:"templateFile"`
	ServiceTemplateFiles    []string `json:"serviceTemplateFiles"`
	MonitoringTemplateFiles []string `json:"monitoringTemplateFiles"`
	ResilienceTemplateFiles []string `json:"resilienceTemplateFiles"`
}

type ComposeConfig struct {
	WorkingDir   string   `yaml:"workingDir,omitempty" json:"workingDir,omitempty"`
	ComposeFiles []string `yaml:"composeFiles,omitempty" json:"composeFiles,omitempty"`
	EnvFiles     []string `yaml:"envFiles,omitempty" json:"envFiles,omitempty"`
}

type TraefikConfig struct {
	Image        string  `yaml:"image,omitempty" json:"image,omitempty"`
	Name         string  `yaml:"name,omitempty" json:"name,omitempty"`
	Namespace    string  `yaml:"namespace,omitempty" json:"namespace,omitempty"`
	AdminEnabled bool    `yaml:"adminEnabled,omitempty" json:"adminEnabled,omitempty"`
	AdminPort    int     `yaml:"adminPort,omitempty" json:"adminPort,omitempty"`
	Ports        []*Port `yaml:"ports,omitempty" json:"ports,omitempty"`
}

func getDefaultTraefikConfig() TraefikConfig {
	image := os.Getenv("TRAEFIK_IMAGE")
	if image == "" {
		image = "traefik:v3.1"
	}

	return TraefikConfig{
		Image:        image,
		Name:         "traefik",
		Namespace:    "test-bed",
		AdminEnabled: true,
		AdminPort:    8080,
		Ports: []*Port{
			&Port{
				Protocol: "TCP",
				Port:     8000,
				Name:     "http",
				NodePort: 9000,
			},
			&Port{
				Protocol: "TCP",
				Port:     8080,
				Name:     "admin",
			},
		},
	}
}

type Port struct {
	Protocol string `yaml:"protocol,omitempty" json:"protocol,omitempty"`
	Port     int    `yaml:"port,omitempty" json:"port,omitempty"`
	Name     string `yaml:"name,omitempty" json:"name,omitempty"`
	NodePort int    `yaml:"nodePort,omitempty" json:"nodePort,omitempty"`
}

type AlloyConfig struct {
}

func configDir(envfile string) (string, error) {
	err := godotenv.Load(envfile)
	if err != nil {
		return "", err
	}

	envConfigDir := os.Getenv("CONFIG_DIR")
	if envConfigDir != "" {
		return envConfigDir, nil
	}
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	dir := fmt.Sprintf("%s%c%s", userConfigDir, os.PathSeparator, "/plumage")
	return dir, nil
}

func findOrCreateConfigDir(projectName string, envFile string) (string, error) {
	dir, err := configDir(envFile)
	if err != nil {
		return "", err
	}
	log.Println(dir)
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		log.Printf("\nCreating config directory: %s\n", dir)
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return "", err
		}
		return dir, nil
	}
	if dir != "" {
		return dir, nil
	}

	return dir, nil
}
