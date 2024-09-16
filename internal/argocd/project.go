package argocd

import (
	"context"
	"fmt"
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

func (c *Client) createProject(ctx context.Context, proj *v1alpha1.AppProject) (*v1alpha1.AppProject, error) {
	if proj == nil {
		return nil, fmt.Errorf("project is nil")
	}

	if c.projectClient == nil {
		fmt.Println("Project client is nil?")
	}

	output, err := c.projectClient.Create(ctx, &project.ProjectCreateRequest{
		Project: proj,
		Upsert:  true,
	})

	if err != nil {
		return nil, err
	}
	fmt.Println(output.Name)
	return output, nil
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

func projectWithDestinations(destinations []v1alpha1.ApplicationDestination) createProjectFunc {
	return func(project *v1alpha1.AppProject) {
		project.Spec.Destinations = make([]v1alpha1.ApplicationDestination, len(destinations))
		copy(project.Spec.Destinations, destinations)
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
