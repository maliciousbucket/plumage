package orchestration

import "context"

func AddRepoCredentials(ctx context.Context, argoClient ArgoClient) error {
	return argoClient.AddRepoCredentials(ctx)
}
