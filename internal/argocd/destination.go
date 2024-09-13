package argocd

import (
	"context"
	"fmt"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/application"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/project"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
)

func (c *Client) AddProjectDestination(ctx context.Context, projectName, server, namespace, name string) error {
	p, err := c.GetProject(ctx, projectName)
	if err != nil {
		return err
	}

	p.Spec.Destinations = []v1alpha1.ApplicationDestination{
		{
			Server:    server,
			Namespace: namespace,
			Name:      name,
		},
	}

	_, err = c.projectClient.Update(ctx, &project.ProjectUpdateRequest{Project: p})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) AddApplicationDestination(ctx context.Context, appName, server, namespace, name string) error {
	app, err := c.GetApplication(ctx, appName)
	if err != nil {
		return err
	}
	if app == nil {
		return fmt.Errorf("project %s not found", appName)
	}

	app.Spec.Destination = v1alpha1.ApplicationDestination{
		Server:    server,
		Namespace: namespace,
		Name:      name,
	}
	_, err = c.applicationClient.Update(ctx, &application.ApplicationUpdateRequest{Application: app})
	if err != nil {
		return err
	}
	return nil
}
