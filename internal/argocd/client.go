package argocd

import (
	"fmt"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/account"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/application"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/cluster"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/project"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/repocreds"
	"github.com/joho/godotenv"
	"os"
)

//https://nemre.medium.com/manage-argocd-resources-programmatically-with-golang-5fa825f1f36e

type Connection struct {
	Address string
	Token   string
}

type Client struct {
	projectClient     project.ProjectServiceClient
	clusterClient     cluster.ClusterServiceClient
	applicationClient application.ApplicationServiceClient
	accountClient     account.AccountServiceClient
	credsClient       repocreds.RepoCredsServiceClient
}

func GetConnection() (Connection, error) {
	var con Connection
	err := godotenv.Load()
	if err != nil {
		return con, err
	}

	address := os.Getenv("ARGOCD_ADDRESS")
	token := os.Getenv("ARGOCD_TOKEN")
	if address == "" {
		return con, fmt.Errorf("ARGOCD_ADDRESS environment variable not set")
	}
	if token == "" {
		return con, fmt.Errorf("ARGOCD_TOKEN environment variable not set")
	}
	con.Address = address
	con.Token = token
	return con, nil

}

func NewClient(c Connection) (*Client, error) {
	apiClient, err := apiclient.NewClient(&apiclient.ClientOptions{
		ServerAddr:           fmt.Sprintf(c.Address),
		Insecure:             true,
		AuthToken:            c.Token,
		PortForward:          true,
		PortForwardNamespace: "argocd",
	})

	if err != nil {
		return nil, err
	}

	_, projectClient, err := apiClient.NewProjectClient()
	if err != nil {
		return nil, err
	}

	_, clusterClient, err := apiClient.NewClusterClient()
	if err != nil {
		return nil, err
	}

	_, applicationClient, err := apiClient.NewApplicationClient()

	_, accountClient, err := apiClient.NewAccountClient()

	_, credsClient, err := apiClient.NewRepoCredsClient()

	return &Client{projectClient, clusterClient, applicationClient, accountClient, credsClient}, nil
}
