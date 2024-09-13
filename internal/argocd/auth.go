package argocd

import (
	"context"
	"fmt"
	account "github.com/argoproj/argo-cd/v2/pkg/apiclient/account"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/repocreds"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	"os"
)

func (c *Client) AddRepoCredentials(ctx context.Context) error {
	username := os.Getenv("GITHUB_USERNAME")
	password := os.Getenv("GITHUB_PASSWORD")
	repository := os.Getenv("ARGOCD_REPOSITORY")

	_, err := c.credsClient.CreateRepositoryCredentials(ctx, &repocreds.RepoCredsCreateRequest{
		Creds: &v1alpha1.RepoCreds{
			URL:      repository,
			Username: username,
			Password: password,
		},
		Upsert: false,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) getToken(ctx context.Context) (string, error) {
	acc, err := c.accountClient.GetAccount(ctx, &account.GetAccountRequest{
		Name: "galah",
	})
	if err != nil {
		return "", fmt.Errorf("unable to get galah account: %w", err)
	}

	token, err := c.accountClient.CreateToken(ctx, &account.CreateTokenRequest{
		Name: acc.Name,
	})
	if err != nil {
		return "", fmt.Errorf("unable to create token: %w", err)
	}
	return token.Token, nil
}
