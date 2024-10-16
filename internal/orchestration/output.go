package orchestration

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/go-github/v64/github"
	"github.com/maliciousbucket/plumage/pkg/config"
	"github.com/maliciousbucket/plumage/pkg/kplus"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type AppOutput struct {
	Root   string
	Charts []*ChartOutput
}

type ChartOutput struct {
	Version   string `json:"version"`
	Resources map[string]struct {
		Path string `json:"path"`
	} `json:"resources"`
}

func CommitAndPushService(ctx context.Context, cfg *config.GitHubConfig, dir string, chart string, message string) (*GitHubCommitResponse, error) {
	token := cfg.Token()
	if token == "" {
		return nil, errors.New("no auth token found")
	}
	client := github.NewClient(nil).WithAuthToken(token)
	ref, err := getRef(ctx, client, cfg)
	if err != nil {
		return nil, err
	}

	tree, err := getServiceChartTree(ctx, ref, client, cfg, dir, chart)
	if err != nil {
		return nil, err
	}

	res, err := newCommit(ctx, client, ref, tree, cfg, message)
	if err != nil {
		return nil, err
	}

	result := &GitHubCommitResponse{
		Url:     *res.URL,
		Sha:     *res.SHA,
		Message: *res.Message,
		Name:    *res.Committer.Name,
		Email:   *res.Committer.Email,
		Date:    *res.Committer.Date.GetTime(),
	}
	return result, nil

}

func getServiceChartTree(ctx context.Context, ref *github.Reference, client *github.Client, cfg *config.GitHubConfig,
	dir, chart string) (*github.Tree, error) {

	chartPath := filepath.Join(dir, chart)

	files, err := getDirectoryFiles(chartPath, nil)
	if len(files) == 0 {
		return nil, fmt.Errorf("no resources found for chart %s", chart)
	}
	if err != nil {
		return nil, err
	}
	var entries []*github.TreeEntry

	for _, path := range files {
		content, fileErr := getFileContent(path)
		if fileErr != nil {
			return nil, fileErr
		}

		entries = append(entries, &github.TreeEntry{Path: github.String(path), Type: github.String("blob"), Content: github.String(string(content)), Mode: github.String("100644")})
	}
	tree, _, err := client.Git.CreateTree(ctx, cfg.SourceOwner, cfg.SourceRepo, *ref.Object.SHA, entries)
	return tree, err
}

func CommitAndPushResource(ctx context.Context, cfg *config.GitHubConfig, dir, chart, resource string, message string) (*GitHubCommitResponse, error) {
	token := cfg.Token()
	if token == "" {
		return nil, errors.New("no auth token found")
	}
	client := github.NewClient(nil).WithAuthToken(token)
	ref, err := getRef(ctx, client, cfg)
	if err != nil {
		return nil, err
	}

	tree, err := getResourceTree(ctx, ref, client, cfg, dir, chart, resource)
	if err != nil {
		return nil, err
	}

	res, err := newCommit(ctx, client, ref, tree, cfg, message)
	if err != nil {
		return nil, err
	}

	result := &GitHubCommitResponse{
		Url:     *res.URL,
		Sha:     *res.SHA,
		Message: *res.Message,
		Name:    *res.Committer.Name,
		Email:   *res.Committer.Email,
		Date:    *res.Committer.Date.GetTime(),
	}
	return result, nil
}

func getResourceTree(ctx context.Context, ref *github.Reference, client *github.Client, cfg *config.GitHubConfig,
	dir, chart, resource string) (*github.Tree, error) {

	chartPath := filepath.Join(dir, chart)

	files, err := getDirectoryFiles(chartPath, nil)
	if len(files) == 0 {
		return nil, fmt.Errorf("no resources found for chart %s", chart)
	}
	if err != nil {
		return nil, err
	}

	var resourceFile string
	for _, path := range files {
		filename := filepath.Base(path)
		resourceName := getResourceName(filename)
		if resourceName == resource {
			resourceFile = path
		}

	}
	if resourceFile == "" {
		return nil, fmt.Errorf("no resource %s found for chart %s", resource, chart)
	}
	resourceData, err := getFileContent(resourceFile)
	if err != nil {
		return nil, err
	}

	var entries []*github.TreeEntry
	entries = append(entries, &github.TreeEntry{Path: github.String(resourceFile), Type: github.String("blob"), Content: github.String(string(resourceData)), Mode: github.String("100644")})

	tree, _, err := client.Git.CreateTree(ctx, cfg.SourceOwner, cfg.SourceRepo, *ref.Object.SHA, entries)
	return tree, err

}

func getServiceResources(service string, chart *ChartOutput) []string {
	output := []string{}
	for key, value := range chart.Resources {
		if strings.HasPrefix(key, service) {
			output = append(output, value.Path)
		}
	}
	return output
}

func removeRootPath(paths []string) []string {
	filteredPaths := []string{}
	for _, path := range paths {
		parts := strings.Split(path, "/")
		if len(parts) > 1 {
			newPath := strings.Join(parts[1:], "/")
			filteredPaths = append(filteredPaths, newPath)
		}
	}
	return filteredPaths
}

func getObjectMetaFile(dir string) (*ChartOutput, error) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil, fmt.Errorf("directory %s does not exist", dir)
	}

	objectMetaFile := filepath.Join(dir, "construct-metadata.json")
	if _, err := os.Stat(objectMetaFile); os.IsNotExist(err) {
		return nil, fmt.Errorf("metadata file %s does not exist", objectMetaFile)
	}

	file, err := os.Open(objectMetaFile)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var chartOutput ChartOutput
	if err = json.Unmarshal(data, &chartOutput); err != nil {
		return nil, fmt.Errorf("failed to unmarshal chart output: %w", err)
	}
	return &chartOutput, nil

}

func getResourceName(fileName string) string {

	parts := strings.Split(fileName, ".")

	if len(parts) > 2 {

		cleanedParts := parts[1 : len(parts)-1]

		return strings.Join(cleanedParts, ".")
	}

	return fileName
}

func CommitAndPushGateway(ctx context.Context, cfg *config.GitHubConfig, dir string) (*GitHubCommitResponse, error) {
	ingressDir := filepath.Join(dir, "ingress")
	msg := fmt.Sprintf("plumage manifests - gateway - %s", time.Now().String())

	return CommitAndPushService(ctx, cfg, ingressDir, "traefik", msg)
}

type SynthOpts struct {
	SynthTemplate bool
	SynthGateway  bool
	SynthTests    bool
	TemplateFile  string
	OutputDir     string
	Namespace     string
	Monitoring    map[string]string
}

func (s *SynthOpts) validate() error {
	var err error
	if s.Namespace == "" {
		err = errors.New("namespace is required")
	}
	if s.OutputDir == "" {
		err = errors.Join(err, errors.New("output directory is required"))
	}

	return err
}

// TODO: Move?=
func SynthDeployment(opts *SynthOpts) error {
	if !opts.SynthTemplate && !opts.SynthGateway {
		return nil
	}
	validateErr := opts.validate()
	if validateErr != nil {
		return validateErr
	}
	if opts.SynthTemplate {
		err := kplus.SynthTemplate(opts.TemplateFile, opts.OutputDir, opts.Monitoring)
		if err != nil {
			return err
		}
	}

	if opts.SynthGateway {
		err := kplus.SynthGateway(opts.OutputDir, opts.Namespace)
		if err != nil {
			return err
		}
	}

	return nil

}
