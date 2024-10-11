package integration_testing

import (
	"context"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	"github.com/maliciousbucket/plumage/internal/argocd"
	"testing"
)

type argoClient interface {
	GetClusters(ctx context.Context) ([]v1alpha1.Cluster, error)
	CreateCluster(ctx context.Context, name string) (*v1alpha1.Cluster, error)
	AddProjectDestination(ctx context.Context, projectName string, server string, namespace string, name string) error
	AddApplicationDestination(ctx context.Context, appName string, server string, namespace string, name string) error
	CreateProject(ctx context.Context, name string) (*v1alpha1.AppProject, error)
	DeleteProject(ctx context.Context, name string) error
	DeleteProjectWithApps(ctx context.Context, name string) error
	GetProject(ctx context.Context, name string) (*v1alpha1.AppProject, error)
	ListProjects(ctx context.Context) (*v1alpha1.AppProjectList, error)
	GetApplication(ctx context.Context, name string) (*v1alpha1.Application, error)
	ListApplications(ctx context.Context, params *argocd.AppQueryParams) (*v1alpha1.ApplicationList, error)
	CreateIngressApp(ctx context.Context) error
	CreateApplication(ctx context.Context) (*v1alpha1.Application, error)
	AddApplicationToProject(ctx context.Context, appName string, project string, validate bool) (*v1alpha1.ApplicationSpec, error)
	UpdateApplication(ctx context.Context, appName string) (*v1alpha1.Application, error)
	AddRepoCredentials(ctx context.Context) error
	SyncApplicationResources(ctx context.Context, name string) error
	SyncProject(ctx context.Context, name string) error

	CreateMonitoringProject(ctx context.Context) error
	CreateNetworkingProject(ctx context.Context) error
	CreateIngressProject(ctx context.Context) error
	CreateApplicationProject(ctx context.Context, app string) error
}

func TestArgoProjects(t *testing.T) {

}

func TestArgoApps(t *testing.T) {

}
