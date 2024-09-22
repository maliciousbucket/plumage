package commands

import "github.com/maliciousbucket/plumage/internal/argocd"

var (
	argoClient *argocd.Client
)

func newClient() (*argocd.Client, error) {
	if argoClient != nil {
		return argoClient, nil
	}
	conn, err := argocd.GetConnection()
	if err != nil {
		return nil, err
	}
	return argocd.NewClient(conn)
}
