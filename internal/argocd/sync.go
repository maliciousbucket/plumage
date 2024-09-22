package argocd

import (
	"context"
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

	sync, err := c.applicationClient.Sync(ctx, &application.ApplicationSyncRequest{
		Name:         &name,
		Revision:     nil,
		DryRun:       &dryRun,
		Prune:        &prune,
		AppNamespace: &ns,
		Project:      &project,
	})
	if err != nil {
		return err
	}
	fmt.Printf("Syncing %s in namespace %s with project %s\nStatus: %+v", sync.Name, sync.Namespace, project, sync.Status)
	return nil
}

func (c *Client) SyncProject(ctx context.Context, name string) error {

	project, err := c.GetProject(context.Background(), name)
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
			return err
		}
	}
	return nil
}
