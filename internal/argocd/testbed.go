package argocd

import (
	"context"
	"errors"
	"fmt"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/application"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

const (
	traefikPath         = "dist/ingress/traefik"
	infraDashboardsPath = "dist/ingress/dashboards"
	appPath             = "dist"
)

func (c *Client) CreateIngressProject(ctx context.Context) error {
	var project string
	existing, _ := c.GetProject(ctx, "ingress")

	if existing == nil {
		newProj, err := c.createIngressProject(ctx)
		if err != nil {
			return err
		}
		project = newProj
	} else {
		project = "ingress"
	}

	apps, err := c.ListApplications(ctx, nil)
	if err != nil {
		return err
	}

	if apps.Items == nil || len(apps.Items) == 0 {
		for _, item := range apps.Items {
			if item.Name == "traefik" {
				log.Println("traefik app already exists")
				if item.Spec.Project != "ingress" {
					_, err = c.AddApplicationToProject(ctx, item.Name, project, true)
					if err != nil {
						return err
					}
					log.Printf("added traefik app %s to project: %s\n", item.Name, project)
					return nil
				}
				log.Printf("traefik app %s already exists in project: %s\n", item.Name, project)
				return nil
			}
		}
	}

	if err = c.createIngressApp(ctx, project); err != nil {
		return err
	}
	log.Printf("Created Ingress app in Project %s", project)
	return nil
}

func (c *Client) createIngressProject(ctx context.Context) (string, error) {
	sources := []string{argoCDRepo}
	namespaces := []string{"galah-testbed"}
	destinations := []v1alpha1.ApplicationDestination{
		{
			Server:    defaultServer,
			Namespace: "argocd",
		},
		{
			Server:    defaultServer,
			Namespace: "galah-testbed",
		},
	}
	return c.createGalahProject(ctx,
		"ingress",
		"traefik ingress for test bed",
		sources,
		namespaces,
		destinations,
	)
}

func (c *Client) CreateIngressApp(ctx context.Context) error {
	return c.createIngressApp(ctx, "ingress")
}

func (c *Client) createIngressApp(ctx context.Context, project string) error {

	proj, err := c.GetProject(ctx, project)
	if err != nil {
		return fmt.Errorf("failed to get project: %w", err)
	}

	return c.addTestBedApp(ctx, "traefik", traefikPath, proj.Name)
}

func (c *Client) CreateApplicationProject(ctx context.Context, app string) error {
	project, err := c.createApplicationProject(ctx, app)
	if err != nil {
		return err
	}
	log.Printf("Project Created Successfully: %s", project)
	return nil
}

func (c *Client) createApplicationProject(ctx context.Context, app string) (string, error) {
	sources := []string{argoCDRepo}
	namespaces := []string{
		"galah-testbed",
	}
	destinations := []v1alpha1.ApplicationDestination{
		{
			Server:    defaultServer,
			Namespace: "argocd",
		},
		{
			Server:    defaultServer,
			Namespace: "galah-testbed",
		},
	}

	return c.createGalahProject(ctx,
		app,
		"app synthesised with plumage",
		sources,
		namespaces,
		destinations)

}

func (c *Client) CreateServiceApplications(ctx context.Context, app string, services []string) error {
	if app == "" {
		return errors.New("app cannot be empty")
	}

	if services == nil || len(services) == 0 {
		return errors.New("services cannot be empty")
	}

	project, err := c.GetProject(ctx, app)
	if err != nil {
		return fmt.Errorf("project %s not found: %v", app, err)
	}

	for _, service := range services {
		path := fmt.Sprintf("%s/%s/%s/", appPath, app, service)
		appErr := c.addTestBedApp(ctx, service, path, project.Name)
		if appErr != nil {
			return fmt.Errorf("add test bed app %s failed: %v", service, appErr)
		}
		log.Printf("added test bed app %s successfully", service)

	}

	return nil
}

func (c *Client) addTestBedApp(ctx context.Context, name, path, project string) error {
	upsert := true
	validate := true

	app, err := c.applicationClient.Create(ctx, &application.ApplicationCreateRequest{
		Application: &v1alpha1.Application{
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				Namespace: "argocd",
				Finalizers: []string{
					"resources-finalizer.argocd.argoproj.io",
				},
			},
			Spec: v1alpha1.ApplicationSpec{
				Project: project,
				Source: &v1alpha1.ApplicationSource{
					RepoURL:        argoCDRepo,
					Path:           path,
					TargetRevision: "HEAD",
				},
				Destination: v1alpha1.ApplicationDestination{
					Server:    defaultServer,
					Namespace: "galah-testbed",
				},
				SyncPolicy: &v1alpha1.SyncPolicy{
					Automated: &v1alpha1.SyncPolicyAutomated{
						SelfHeal: false,
					},
				},
			},
		},
		Upsert:   &upsert,
		Validate: &validate,
	})
	if err != nil {
		return err
	}
	log.Printf("Created application %s/%s", app.GetNamespace(), app.GetName())
	return nil
}
