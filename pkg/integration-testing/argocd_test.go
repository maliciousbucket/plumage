package integration_testing

import (
	"context"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	"github.com/joho/godotenv"
	"github.com/maliciousbucket/plumage/internal/argocd"
	"github.com/maliciousbucket/plumage/internal/orchestration"
	"github.com/stretchr/testify/assert"
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
	kubeContainer := NewKubernetesContainer()
	if err := WithHostPorts(30443)(kubeContainer); err != nil {
		t.Fatalf("error setting up test container: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()
	if err := SetupKubernetesContainer(ctx, kubeContainer); err != nil {
		t.Fatalf("error setting up clients: %v", err)
	}
	defer func() {
		if err := kubeContainer.Container.Terminate(ctx); err != nil {
			t.Fatal(err)
		}
	}()

	helm, kube, err := setupHelmKubeClients(ctx, kubeContainer)
	if err != nil {
		t.Fatalf("error setting up kubernetes and helm clients: %v", err)
	}
	t.Run("install argo chart", func(t *testing.T) {
		err = installArgo(t, ctx, helm, kube)
		if err != nil {
			t.Fatalf("error installing argo chart: %v", err)
		}
	})

	err = setArgoEnv(ctx, kubeContainer, 30443)
	assert.NoErrorf(t, err, "error setting up argo env")

	t.Run("set argo token", func(t *testing.T) {
		if err = setArgoToken(ctx, kube); err != nil {
			t.Fatalf("error setting Argo token: %v", err)
		}
	})
	var argo argoClient

	t.Run("create argo client", func(t *testing.T) {
		argoC, argoErr := newArgoClientWithEnv()
		assert.NoErrorf(t, err, "error creating Argo Client: %v", argoErr)
		if argoC == nil {
			t.Fatalf("error creating Argo Client: %v", err)
		}
		argo = argoC
	})

	t.Run("create argo project", func(t *testing.T) {
		project, err := argo.CreateProject(ctx, "testing")
		assert.NoErrorf(t, err, "error creating project: %v", err)
		if project.Name != "testing" {
			t.Errorf("error creating project: %v", project)
		}
		found, err := argo.GetProject(ctx, project.Name)
		assert.NoErrorf(t, err, "error getting project: %v", err)
		if found == nil {
			t.Fatalf("error getting project: %v", project)
		}

	})

	t.Run("create argo application", func(t *testing.T) {
		err = argo.CreateServiceApplications(ctx, "testing", "test-app", []string{"test-service"})
		assert.NoErrorf(t, err, "error creating application: %v", err)
	})

	t.Run("add application to project", func(t *testing.T) {
		_, err = argo.AddApplicationToProject(ctx, "test-app", "testing", true)
		assert.NoErrorf(t, err, "error adding application to project: %v", err)
	})

	t.Run("sync project", func(t *testing.T) {
		err = argo.SyncProject(ctx, "testing")
		assert.NoErrorf(t, err, "error syncing project: %v", err)
	})

	t.Run("delete project", func(t *testing.T) {
		err = argo.DeleteProjectWithApps(ctx, "testing")
		assert.NoErrorf(t, err, "error deleting project: %v", err)
		time.Sleep(30 * time.Second)
		foundProj, err := argo.GetProject(ctx, "testing")
		assert.Errorf(t, err, "project should be deleted: %v", err)
		if foundProj != nil {
			t.Errorf("project should be deleted: %v", foundProj)
		}

		foundApp, err := argo.GetApplication(ctx, "test-app")
		assert.Errorf(t, err, "application should be deleted: %v", err)

		if foundApp != nil {
			t.Errorf("application should be deleted: %v", foundApp)
		}
	})
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
