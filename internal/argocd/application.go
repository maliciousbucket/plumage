package argocd

import (
	"context"
	"fmt"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/application"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	argoApiVersion  = "argoproj.io/v1alpha1"
	applicationKind = "Application"
)

func (c *Client) GetApplication(ctx context.Context, name string) (*v1alpha1.Application, error) {
	return c.applicationClient.Get(ctx, &application.ApplicationQuery{
		Name: &name,
	})
}

func (c *Client) ListApplications(ctx context.Context, params *AppQueryParams) (*v1alpha1.ApplicationList, error) {
	var query application.ApplicationQuery

	if params != nil {
		for _, opt := range params.Options {
			opt(&query)
		}
	}

	return c.applicationClient.List(ctx, &query)
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

func (c *Client) CreateApplication(ctx context.Context) (*v1alpha1.Application, error) {
	return c.applicationClient.Create(ctx, &application.ApplicationCreateRequest{
		Application: &v1alpha1.Application{
			TypeMeta: metav1.TypeMeta{
				Kind:       "",
				APIVersion: "",
			},
			ObjectMeta: metav1.ObjectMeta{},
			Spec: v1alpha1.ApplicationSpec{
				Source: nil,
				Destination: v1alpha1.ApplicationDestination{
					Server:    "",
					Namespace: "",
					Name:      "",
				},
				Project:              "",
				SyncPolicy:           nil,
				IgnoreDifferences:    nil,
				Info:                 nil,
				RevisionHistoryLimit: nil,
				Sources:              nil,
			},
			Operation: nil,
		},
		Upsert:   nil,
		Validate: nil,
	})

}

func (c *Client) AddApplicationToProject(ctx context.Context, appName string, project string, validate bool) (*v1alpha1.ApplicationSpec, error) {
	app, err := c.GetApplication(ctx, appName)
	if err != nil {
		return nil, err
	}
	if app == nil {
		return nil, fmt.Errorf("application %s not found", appName)
	}

	spec := app.Spec
	spec.Project = project

	return c.applicationClient.UpdateSpec(ctx, &application.ApplicationUpdateSpecRequest{
		Name:         &app.Name,
		Spec:         &spec,
		Validate:     &validate,
		AppNamespace: &app.Namespace,
		Project:      &project,
	}, nil)
}

func (c *Client) UpdateApplication(ctx context.Context, appName string) (*v1alpha1.Application, error) {
	app, err := c.GetApplication(ctx, appName)
	if err != nil {
		return nil, err
	}

	return c.applicationClient.Update(ctx, &application.ApplicationUpdateRequest{
		Application: app,
	})
}
