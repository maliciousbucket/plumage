package orchestration

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/ProtonMail/go-crypto/openpgp"
	"github.com/google/go-github/v64/github"
	"github.com/maliciousbucket/plumage/pkg/config"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

func CommitAndPush(ctx context.Context, cfg *config.GitHubConfig, dir string, message string) (*GitHubCommitResponse, error) {
	token := cfg.Token()
	if token == "" {
		return nil, errors.New("no auth token found")
	}
	client := github.NewClient(nil).WithAuthToken(token)
	ref, err := getRef(ctx, client, cfg)
	if err != nil {
		return nil, err
	}
	denyList := []string{}
	tree, err := getTree(ctx, ref, client, cfg, dir, denyList)
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

type GitHubCommitResponse struct {
	Url     string    `json:"url"`
	Sha     string    `json:"sha"`
	Message string    `json:"message"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Date    time.Time `json:"date"`
}

func CommitAndPushTestBed(ctx context.Context, cfg *config.AppConfig, app string) (*GitHubCommitResponse, *GitHubCommitResponse, error) {
	ghCfg, err := config.NewGithubConfig(cfg.ConfigDir, "github.yaml")
	if err != nil {
		log.Fatal(err)
	}
	path := fmt.Sprintf("%s/%s", cfg.OutputDir, app)
	msg := fmt.Sprintf("Plumage Manifests - %s - %s", app, time.Now().String())
	templateCommit, err := CommitAndPush(ctx, ghCfg, path, msg)
	if err != nil {
		return nil, nil, err
	}
	log.Printf("\nCommits for App %s successful\n", app)
	gatewayCommit, err := CommitAndPushGateway(ctx, ghCfg, cfg.OutputDir)
	if err != nil {
		return templateCommit, nil, err
	}
	return templateCommit, gatewayCommit, nil
}

func newCommit(ctx context.Context, client *github.Client, ref *github.Reference, tree *github.Tree,
	cfg *config.GitHubConfig, message string) (*github.Commit, error) {

	commit, err := createCommit(ctx, client, ref, tree, cfg, message)
	if err != nil {
		return nil, err
	}
	opts := github.CreateCommitOptions{}

	if cfg.PrivateKey() != "" {
		armoredBlock, e := os.ReadFile(cfg.PrivateKey())
		if e != nil {
			return nil, e
		}
		keyring, e := openpgp.ReadArmoredKeyRing(bytes.NewReader(armoredBlock))
		if e != nil {
			return nil, e
		}
		if len(keyring) != 1 {
			return nil, errors.New("expected exactly one key in the keyring")
		}
		key := keyring[0]
		opts.Signer = github.MessageSignerFunc(func(w io.Writer, r io.Reader) error {
			return openpgp.ArmoredDetachSign(w, key, r, nil)
		})
	}

	newGitCommit, _, err := client.Git.CreateCommit(ctx, cfg.SourceOwner, cfg.SourceRepo, commit, &opts)
	if err != nil {
		return nil, err
	}

	ref.Object.SHA = newGitCommit.SHA
	_, _, err = client.Git.UpdateRef(ctx, cfg.SourceOwner, cfg.SourceRepo, ref, false)
	return newGitCommit, err

}

func createCommit(ctx context.Context, client *github.Client, ref *github.Reference, tree *github.Tree,
	cfg *config.GitHubConfig, message string) (*github.Commit, error) {

	parent, _, err := client.Repositories.GetCommit(ctx, cfg.SourceOwner, cfg.SourceRepo, *ref.Object.SHA, nil)
	if err != nil {
		return nil, err
	}
	parent.Commit.SHA = parent.SHA

	date := time.Now()
	author := &github.CommitAuthor{Date: &github.Timestamp{Time: date}, Name: &cfg.AuthorName, Email: &cfg.AuthorEmail}
	commit := &github.Commit{Author: author, Message: &message, Tree: tree, Parents: []*github.Commit{parent.Commit}}
	return commit, nil
}

func getTree(ctx context.Context, ref *github.Reference, client *github.Client, cfg *config.GitHubConfig, dir string, denyList []string) (*github.Tree, error) {
	files, err := getDirectoryFiles(dir, denyList)
	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, fmt.Errorf("no files found in %s", dir)
	}

	var entries []*github.TreeEntry

	for _, file := range files {
		content, err := getFileContent(file)
		if err != nil {
			return nil, err
		}
		entries = append(entries, &github.TreeEntry{Path: github.String(file), Type: github.String("blob"), Content: github.String(string(content)), Mode: github.String("100644")})

	}

	tree, _, err := client.Git.CreateTree(ctx, cfg.SourceOwner, cfg.SourceRepo, *ref.Object.SHA, entries)

	return tree, err
}

func getDirectoryFiles(dir string, denyList []string) ([]string, error) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return []string{}, fmt.Errorf("directory %s does not exist", dir)
	}
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		fmt.Println(dir)
		fmt.Println(path)
		if err != nil {
			return err
		}
		for _, deny := range denyList {
			if deny, _ := regexp.MatchString(deny, info.Name()); !deny {

			}
		}
		if info.Mode().IsRegular() {
			files = append(files, path)
		}

		return nil

	})
	return files, err

}

func getFileContent(file string) ([]byte, error) {

	b, err := os.ReadFile(file)
	return b, err
}

func getRef(ctx context.Context, client *github.Client, cfg *config.GitHubConfig) (*github.Reference, error) {
	cmtBranch := fmt.Sprintf("refs/heads/%s", cfg.CommitBranch)
	if ref, _, err := client.Git.GetRef(ctx, cfg.SourceOwner, cfg.SourceRepo, cmtBranch); err == nil {
		return ref, nil
	}

	//if cfg.CommitBranch == cfg.BaseBranch {
	//	return nil, errors.New("no commit branch found")
	//}

	if cfg.BaseBranch == "" {
		return nil, errors.New("the base branch should not be empty when the commit branch does not exist")
	}
	//null pointer dereference xdd
	var baseRef *github.Reference
	var err error
	baseBranch := fmt.Sprintf("refs/heads/%s", cfg.CommitBranch)
	if baseRef, _, err = client.Git.GetRef(ctx, cfg.SourceOwner, cfg.SourceRepo, baseBranch); err != nil {
		return nil, err
	}
	newRef := &github.Reference{Ref: github.String(cmtBranch), Object: &github.GitObject{SHA: baseRef.Object.SHA}}
	ref, _, err := client.Git.CreateRef(ctx, cfg.SourceOwner, cfg.SourceRepo, newRef)
	return ref, err

}
