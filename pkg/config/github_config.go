package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"
)

type GitHubConfig struct {
	SourceOwner  string `yaml:"sourceOwner"`
	SourceRepo   string `yaml:"sourceRepo"`
	CommitBranch string `yaml:"commitBranch"`
	BaseBranch   string `yaml:"baseBranch"`
	AuthorName   string `yaml:"authorName"`
	AuthorEmail  string `yaml:"authorEmail"`
	EnvFile      string `yaml:"envFile"`
	TargetDir    string `yaml:"targetDir"`
	privateKey   string
	githubToken  string
}

func (g *GitHubConfig) Token() string {
	return g.githubToken
}

func (g *GitHubConfig) PrivateKey() string {
	return g.privateKey
}

func NewGithubConfig(configDir string, file string) (*GitHubConfig, error) {
	filename := "github.yaml"

	if file != "" {
		filename = file
	}
	filePath := filepath.Join(configDir, filename)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("GitHub config file %s does not exist", filePath)
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	config := &GitHubConfig{}
	if err := yaml.Unmarshal(content, config); err != nil {
		return nil, err
	}
	err = loadGitHubEnv(configDir, config.EnvFile, config)
	if err != nil {
		return nil, err
	}
	return config, nil

}

func (g *GitHubConfig) LoadExtraEnv(files string) error {
	envFiles := strings.Split(files, ",")
	err := g.loadExtraEnv(envFiles)
	if err != nil {
		return err
	}
	return nil
}

func (g *GitHubConfig) loadExtraEnv(files []string) error {
	for _, file := range files {
		err := godotenv.Load(file)
		if err != nil {
			return fmt.Errorf("error loafing Env file. File: %s", file)
		}
	}
	sourceOwner := os.Getenv("SOURCE_OWNER")
	sourceRepo := os.Getenv("SOURCE_REPO")
	commitBranch := os.Getenv("COMMIT_BRANCH")
	baseBranch := os.Getenv("BASE_BRANCH")
	authorName := os.Getenv("AUTHOR_NAME")
	authorEmail := os.Getenv("AUTHOR_EMAIL")
	targetDir := os.Getenv("TARGET_DIR")

	if sourceOwner != "" {
		g.SourceOwner = sourceOwner
	}
	if sourceRepo != "" {
		g.SourceRepo = sourceRepo
	}
	if commitBranch != "" {
		g.CommitBranch = commitBranch
	}
	if baseBranch != "" {
		g.BaseBranch = baseBranch
	}
	if authorName != "" {
		g.AuthorName = authorName
	}
	if authorEmail != "" {
		g.AuthorEmail = authorEmail
	}
	if targetDir != "" {
		g.TargetDir = targetDir
	}
	return nil
}

func loadGitHubEnv(configDir string, envFile string, g *GitHubConfig) error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	if envFile != "" {
		path := filepath.Join(configDir, envFile)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			return fmt.Errorf("GitHub Env file %s does not exist", path)
		}
		err = godotenv.Load(envFile)
	}

	pvtKey := os.Getenv("GITHUB_PRIVATE_KEY")
	g.privateKey = pvtKey
	gitHubToken := os.Getenv("GITHUB_TOKEN")
	if gitHubToken == "" {
		return fmt.Errorf("GITHUB_TOKEN environment variable not set")
	}

	g.githubToken = gitHubToken
	return nil
}
