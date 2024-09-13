package argocd

import (
	"context"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/project"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) CreateProject(ctx context.Context, name string) (*v1alpha1.AppProject, error) {
	return c.projectClient.Create(ctx, &project.ProjectCreateRequest{
		Project: &v1alpha1.AppProject{
			ObjectMeta: v1.ObjectMeta{
				Name: name,
			},
		},
	})
}

func (c *Client) createProject(ctx context.Context, proj *v1alpha1.AppProject) error {
	_, err := c.projectClient.Create(ctx, &project.ProjectCreateRequest{
		Project: proj,
	}, nil)
	if err != nil {
		return err
	}
	return nil
}

type createProjectFunc func(project *v1alpha1.AppProject)

func projectWithNamespace(namespace string) createProjectFunc {
	return func(project *v1alpha1.AppProject) {
		project.Namespace = namespace
	}
}

func projectWithName(name string) createProjectFunc {
	return func(project *v1alpha1.AppProject) {
		project.Name = name
	}
}

func projectWithSourceNamespaces(namespaces []string) createProjectFunc {
	return func(project *v1alpha1.AppProject) {
		project.Spec.SourceNamespaces = namespaces
	}
}

func projectWithSources(sources []string) createProjectFunc {
	return func(project *v1alpha1.AppProject) {
		project.Spec.SourceRepos = sources
	}
}

func projectWithDestinations([]v1alpha1.ApplicationDestination) createProjectFunc {
	return func(project *v1alpha1.AppProject) {
		project.Spec.Destinations = make([]v1alpha1.ApplicationDestination, len(project.Spec.Destinations))
		copy(project.Spec.Destinations, project.Spec.Destinations)
	}
}

func projectWithDescription(description string) createProjectFunc {
	return func(project *v1alpha1.AppProject) {
		project.Spec.Description = description
	}
}

func (c *Client) DeleteProject(ctx context.Context, name string) error {
	_, err := c.projectClient.Delete(ctx, &project.ProjectQuery{
		Name: name,
	})
	return err
}

func (c *Client) GetProject(ctx context.Context, name string) (*v1alpha1.AppProject, error) {
	return c.projectClient.Get(ctx, &project.ProjectQuery{Name: name})
}
