package config

type AppConfig struct {
	ConfigDir     string
	ProjectDir    string
	ComposeConfig *ComposeConfig
	OutputDir     string
}

type ComposeConfig struct {
	WorkingDir   string   `yaml:"-,omitempty" json:"-,omitempty"`
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

type Port struct {
	Protocol string `yaml:"protocol,omitempty" json:"protocol,omitempty"`
	Port     int    `yaml:"port,omitempty" json:"port,omitempty"`
	Name     string `yaml:"name,omitempty" json:"name,omitempty"`
	NodePort int    `yaml:"nodePort,omitempty" json:"nodePort,omitempty"`
}

type AlloyConfig struct {
}
