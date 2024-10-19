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

const (
	argoVersion             = "argo-cd-7.6.8"
	promOperatorVersion     = "14.0.0"
	kubeStateMetricsVersion = "5.25.1"
	certManagerVersion      = "v1.15.3"
	k6OperatorVersion       = "3.9.0"
	metricsVersion          = "3.12.2"
	kubePromStackVersion    = "65.2.0"
)

type AppConfig struct {
	ProjectName      string            `yaml:"project"`
	ConfigDir        string            `yaml:"getConfigDir"`
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
	OutputDir      string         `yaml:"outputDir"`
	Namespace      string         `yaml:"namespace"`
	ComposeConfig  ComposeConfig  `yaml:"compose"`
	TemplateConfig TemplateConfig `yaml:"template"`
	ChartConfig    ChartConfig    `yaml:"helmConfig"`
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
			//TODO: From Config
			WorkingDir:   "testdata",
			TemplateFile: "/chirp/template.yaml",
		},
		ChartConfig: ChartConfig{
			ArgoVersion:                   argoVersion,
			ArgoValuesFile:                "",
			KubeStateMetricsVersion:       kubeStateMetricsVersion,
			KubeStateMetricsValuesFile:    "",
			MetricsVersion:                metricsVersion,
			MetricsValuesFile:             "",
			PromOperatorVersion:           promOperatorVersion,
			PromOperatorValuesFile:        "",
			CertManagerVersion:            certManagerVersion,
			CertManagerValuesFile:         "",
			KubePrometheusStackVersion:    kubePromStackVersion,
			KubePrometheusStackValuesFile: "",
			K6Version:                     k6OperatorVersion,
			K6ValuesFile:                  "",
			Charts:                        nil,
		},
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
	WorkingDir           string   `yaml:"workingDir"`
	TemplateFile         string   `yaml:"templateFile"`
	ServiceTemplateFiles []string `json:"serviceTemplateFiles"`
}

type ComposeConfig struct {
	WorkingDir   string   `yaml:"workingDir,omitempty" json:"workingDir,omitempty"`
	ComposeFiles []string `yaml:"composeFiles,omitempty" json:"composeFiles,omitempty"`
	EnvFiles     []string `yaml:"envFiles,omitempty" json:"envFiles,omitempty"`
}

type Port struct {
	Protocol string `yaml:"protocol,omitempty" json:"protocol,omitempty"`
	Port     int    `yaml:"port,omitempty" json:"port,omitempty"`
	Name     string `yaml:"name,omitempty" json:"name,omitempty"`
	NodePort int    `yaml:"nodePort,omitempty" json:"nodePort,omitempty"`
}

type AlloyConfig struct {
}

type ChartConfig struct {
	ArgoVersion                   string             `yaml:"argoVersion,omitempty"`
	ArgoValuesFile                string             `yaml:"argoValuesFile,omitempty"`
	KubeStateMetricsVersion       string             `yaml:"stateMetricsVersion,omitempty"`
	KubeStateMetricsValuesFile    string             `yaml:"stateMetricsValuesFile,omitempty"`
	MetricsVersion                string             `yaml:"metricsVersion,omitempty"`
	MetricsValuesFile             string             `yaml:"metricsValuesFile,omitempty"`
	PromOperatorVersion           string             `yaml:"promOperatorVersion,omitempty"`
	PromOperatorValuesFile        string             `yaml:"promOperatorValuesFile,omitempty"`
	CertManagerVersion            string             `yaml:"certManagerVersion,omitempty"`
	CertManagerValuesFile         string             `yaml:"certManagerValuesFile,omitempty"`
	KubePrometheusStackVersion    string             `yaml:"kubePrometheusStackVersion,omitempty"`
	KubePrometheusStackValuesFile string             `yaml:"KubePrometheusStackValuesFile,omitempty"`
	K6Version                     string             `yaml:"k6Version,omitempty"`
	K6ValuesFile                  string             `yaml:"k6ValuesFile,omitempty"`
	Charts                        *helm.ChartsConfig `yaml:"chartsConfig,omitempty"`
}

func (c *ChartConfig) ToBaseOpts() *helm.BaseChartOpts {
	return &helm.BaseChartOpts{
		KubeStateMetrics:    c.KubeStateMetricsVersion,
		StateMetricsValues:  c.KubeStateMetricsValuesFile,
		CertManager:         c.CertManagerVersion,
		CertManagerValues:   c.CertManagerValuesFile,
		MetricsServer:       c.MetricsVersion,
		MetricsServerValues: c.MetricsValuesFile,
		PromOperatorCRDs:    c.PromOperatorVersion,
		PromOperatorValues:  c.PromOperatorValuesFile,
		ArgoCD:              c.ArgoVersion,
		ArgoValues:          c.ArgoValuesFile,
		K6Operator:          c.K6Version,
		K6Values:            c.K6ValuesFile,
	}
}

func getConfigDir(envFile string) (string, error) {
	err := godotenv.Load(envFile)
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
	dir, err := getConfigDir(envFile)
	if err != nil {
		return "", err
	}
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
