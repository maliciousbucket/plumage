package argocd

import (
	"context"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/application"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
)

func (c *Client) GetApplication(name string) (*v1alpha1.Application, error) {
	return c.applicationClient.Get(context.Background(), &application.ApplicationQuery{
		Name: &name,
	})
}

func (c *Client) ListApplications(params *AppQueryParams) (*v1alpha1.ApplicationList, error) {
	var query application.ApplicationQuery

	for _, opt := range params.Options {
		opt(&query)
	}
	return c.applicationClient.List(context.Background(), &query)
}

type AppQueryParams struct {
	Options []AppQueryFunc
}

type AppQueryFunc func(query *application.ApplicationQuery)

func WithName(name string) AppQueryFunc {
	return func(query *application.ApplicationQuery) {
		query.Name = &name
	}
}

func WithNamespace(namespace string) AppQueryFunc {
	return func(query *application.ApplicationQuery) {
		query.AppNamespace = &namespace
	}
}

func WithRepository(repository string) AppQueryFunc {
	return func(query *application.ApplicationQuery) {
		query.Repo = &repository
	}
}

func WithProject(project string) AppQueryFunc {
	return func(query *application.ApplicationQuery) {
		query.Project = []string{project}
	}
}
