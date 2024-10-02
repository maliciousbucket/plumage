package argocd

import (
	"context"
	"log"
	"strings"

	"fmt"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/application"
)

func (c *Client) SyncApplicationResources(ctx context.Context, name string) error {
	app, err := c.GetApplication(ctx, name)
	if err != nil {
		return err
	}
	if app == nil {
		return fmt.Errorf("application %s not found", name)
	}

	return c.syncApplication(ctx, app.Name, app.Namespace, app.Spec.Project)
}

func (c *Client) syncApplication(ctx context.Context, name, ns, project string) error {
	prune := true
	dryRun := false

	_, err := c.applicationClient.Sync(ctx, &application.ApplicationSyncRequest{
		Name:         &name,
		Revision:     nil,
		DryRun:       &dryRun,
		Prune:        &prune,
		AppNamespace: &ns,
		Project:      &project,
	})
	if err != nil {
		if strings.Contains(err.Error(), "another operation is already in progress") {
			log.Println("Operation already in progress.")
			return nil
		}
		return err
	}
	return nil
}

func (c *Client) SyncProject(ctx context.Context, name string) error {

	project, err := c.GetProject(ctx, name)
	if err != nil {
		return err
	}
	if project == nil {
		return fmt.Errorf("project %s not found", name)
	}

	params := &AppQueryParams{}
	params.Options = append(params.Options, WithProject(project.Name))

	apps, err := c.ListApplications(ctx, params)
	if err != nil {
		return err
	}
	if len(apps.Items) == 0 {
		return fmt.Errorf("no applications found for project %s", name)
	}

	for _, app := range apps.Items {
		err = c.syncApplication(ctx, app.Name, app.Namespace, project.Name)
		if err != nil {
			return fmt.Errorf("failed to sync application %s: %v", app.Name, err)
		}
	}
	return nil
}

func (c *Client) SyncAllProjects(ctx context.Context) error {
	projects, err := c.ListProjects(ctx)
	if err != nil {
		return err
	}

	if len(projects.Items) == 0 {
		return fmt.Errorf("no projects found in cluster")
	}

	for _, project := range projects.Items {
		if project.Name == "default" {
			//Default ArgoCD project isn't used
			continue
		}
		err = c.SyncProject(ctx, project.Name)
		if err != nil {
			return err
		}
	}
	return nil
}
