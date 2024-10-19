package orchestration

import (
	"context"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	"github.com/maliciousbucket/plumage/internal/argocd"
	"github.com/maliciousbucket/plumage/internal/kubeclient"
	appsV1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"time"
)

type ArgoClient interface {
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
	CreateChaosProject(ctx context.Context, ns string) error
	CreateChaosApp(ctx context.Context, ns, project, path, test string) error
}

type KubeClient interface {
	WatchDeployment(ctx context.Context, ns string, name string, meta bool) error
	WatchAppDeployment(ctx context.Context, ns string, services []string) error
	WaitPodInstanceRunning(ctx context.Context, ns string, name string) error
	WaitPodNameLabelRunning(ctx context.Context, ns string, name string) error
	CreateNamespace(ctx context.Context, ns string) (*kubeclient.NameSpaceInfo, error)
	CheckArgoExists(ctx context.Context, ns string) (*kubeclient.ServiceInfo, error)
	WaitAllArgoPods(ctx context.Context, ns string) error
	SetupArgoLb(ctx context.Context, ns, envFile string) error
	WaitServicePods(ctx context.Context, ns string, name string) error
	WaitAppPods(ctx context.Context, ns, name string, expected int, timeout time.Duration) error

	GetServiceAddress(ctx context.Context, ns string, name string) (string, error)
	CreateGalahArgoAccount(ctx context.Context, ns string) error
	GetArgoPassword(ctx context.Context, ns string) (string, error)
	CheckServiceExists(ctx context.Context, ns string, name string) (bool, error)
	ExposeService(ctx context.Context, ns string, name string, port int, nodePort int, serviceType kubeclient.ExposeServiceType) error
	ListDeployments(ctx context.Context, ns string) ([]appsV1.Deployment, error)
	ListPods(ctx context.Context, namespace string) (*v1.PodList, error)
	GetExternalAddress(ctx context.Context, ns string, name string) ([]string, error)
	GetLoadBalancersForNamespace(ctx context.Context, ns string) ([]*kubeclient.LoadBalancerInfo, error)
}

type HelmClient interface {
	InstallK6(ctx context.Context, version, valuesFile string, replace bool) error
	InstallArgoChart(ctx context.Context, version, valuesFile string) error
	InstallPromOperatorCRDs(ctx context.Context, version, valuesFile string, replace bool) error
}
