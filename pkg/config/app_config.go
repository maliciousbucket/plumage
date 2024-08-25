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
