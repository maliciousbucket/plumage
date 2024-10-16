package integration_testing

import (
	"context"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	"github.com/joho/godotenv"
	"github.com/maliciousbucket/plumage/internal/argocd"
	"github.com/maliciousbucket/plumage/internal/orchestration"
	"log"
	"testing"
	"time"
)

type argoClient interface {
	AddProjectDestination(ctx context.Context, projectName string, server string, namespace string, name string) error
	AddApplicationDestination(ctx context.Context, appName string, server string, namespace string, name string) error
	CreateProject(ctx context.Context, name string) (*v1alpha1.AppProject, error)
	DeleteProject(ctx context.Context, name string) error
	DeleteProjectWithApps(ctx context.Context, name string) error
	GetProject(ctx context.Context, name string) (*v1alpha1.AppProject, error)
	ListProjects(ctx context.Context) (*v1alpha1.AppProjectList, error)
	GetApplication(ctx context.Context, name string) (*v1alpha1.Application, error)
	ListApplications(ctx context.Context, params *argocd.AppQueryParams) (*v1alpha1.ApplicationList, error)
	CreateIngressApp(ctx context.Context, ns string) error
	CreateApplication(ctx context.Context) (*v1alpha1.Application, error)
	AddApplicationToProject(ctx context.Context, appName string, project string, validate bool) (*v1alpha1.ApplicationSpec, error)
	UpdateApplication(ctx context.Context, appName string) (*v1alpha1.Application, error)
	AddRepoCredentials(ctx context.Context, envFile string) error
	SyncApplicationResources(ctx context.Context, name string) error
	SyncProject(ctx context.Context, name string) error

	CreateMonitoringProject(ctx context.Context) error
	CreateIngressProject(ctx context.Context, ns string) error
	CreateApplicationProject(ctx context.Context, app string) error
	CreateServiceApplications(ctx context.Context, ns, app string, services []string) error
}

func newArgoClient() (argoClient, error) {
	conn, err := argocd.GetConnection()
	if err != nil {
		return nil, err
	}
	client, err := argocd.NewClient(conn)
	if err != nil {
		return nil, err
	}
	return client, nil

}

func newArgoClientWithEnv() (*argocd.Client, error) {
	conn, err := argocd.GetConnectionFromEnv(".env.test")
	if err != nil {
		return nil, err
	}
	client, err := argocd.NewClient(conn)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func TestArgoProjects(t *testing.T) {

}

func TestArgoApps(t *testing.T) {

}

func installArgo(t *testing.T, ctx context.Context, helm helmClient, kube kubeClient) error {
	err := helm.InstallArgoChart(ctx, argoVersion, "")
	if err != nil {
		return err
	}
	log.Println("Waiting for argocd deployments....")
	time.Sleep(30 * time.Second)

	if err = orchestration.WatchArgoDeployment(ctx, kube); err != nil {
		return err
	}

	return nil
}

func setArgoToken(ctx context.Context, kube kubeClient) error {

	if err := orchestration.SetArgoToken(ctx, kube); err != nil {
		return err
	}

	return nil
}

func setArgoCreds(ctx context.Context, argo argoClient) error {
	if err := orchestration.AddRepoCredentials(ctx, argo, ".env.test"); err != nil {
		return err
	}
	return nil
}

func setArgoEnv(ctx context.Context, kubeContainer *KubernetesContainer, port int) error {
	ip, err := kubeContainer.Container.Host(ctx)
	if err != nil {
		return err
	}

	ctrPort, err := kubeContainer.Container.MappedPort(ctx, "30443/tcp")
	if err != nil {
		return err
	}
	err = godotenv.Load(".env.test")
	if err != nil {
		return err
	}
	log.Println(ctrPort)

	env, err := godotenv.Read(".env.test")
	if err != nil {
		return err
	}
	log.Println(ip)
	env["ARGO_ADDRESS"] = "127.0.0.1:30443"
	err = godotenv.Write(env, ".env.test")
	if err != nil {
		return err
	}
	return nil

}
