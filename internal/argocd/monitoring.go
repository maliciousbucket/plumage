package argocd

import (
	"context"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
)

const (
	galahMonitoringRepo = ""
	argoCDRepo          = ""
	defaultServer       = "https://kubernetes.default.svc"
	alloyPath           = "kubernetes/alloy"
	grafanaPath         = "kubernetes/grafana"
	lokiPath            = "kubernetes/loki"
	mimirPath           = "kubernetes/mimir"
	minioOperatorPath   = "kubernetes/minio-operator"
	minioTenantPath     = "kubernetes/minio-tenant"
	tempoPath           = "kubernetes/tempo"
	userPath            = "kubernetes/user"

	gatewayPath = "kubernetes/gateway"
)

func (c *Client) CreateMonitoringProject(ctx context.Context) error {
	project, err := c.createMonitoringProject(ctx)
	if err != nil {
		return err
	}

}

func (c *Client) createMonitoringProject(ctx context.Context) (string, error) {
	sources := []string{galahMonitoringRepo}
	//TODO: Change gateway to monitoring gateway
	namespaces := []string{
		"galah-monitoring",
		"galah-tracing",
		"galah-logging",
		"minio-store",
	}

	destinations := []v1alpha1.ApplicationDestination{
		{
			Server:    defaultServer,
			Namespace: "argocd",
		},
	}

	return c.createGalahProject(ctx,
		"galah-monitoring",
		"galah monitoring infrastructure",
		sources,
		namespaces,
		destinations,
	)

}

func (c *Client) createNetworkingProject(ctx context.Context) (string, error) {
	sources := []string{galahMonitoringRepo, argoCDRepo}
	namespaces := []string{
		"gateway",
		"nginx-ingress",
		"galah-monitoring",
	}

	destinations := []v1alpha1.ApplicationDestination{
		{
			Server:    defaultServer,
			Namespace: "argocd",
		},
	}

	return c.createGalahProject(ctx,
		"galah-networking",
		"galah networking infrastructure",
		sources,
		namespaces,
		destinations,
	)
}

func (c *Client) createDashboardProject(ctx context.Context) (string, error) {
	sources := []string{galahMonitoringRepo}
	namespaces := []string{"kubernetes-dashboards"}
	destinations := []v1alpha1.ApplicationDestination{
		{
			Server:    defaultServer,
			Namespace: "argocd",
		},
	}

	return c.createGalahProject(ctx,
		"galah-dashboards",
		"kubernetes dashboards",
		sources,
		namespaces,
		destinations,
	)
}

func (c *Client) createCRDProject(ctx context.Context) (string, error) {
	sources := []string{galahMonitoringRepo}
	namespaces := []string{
		"cert-manager",
		"default",
	}
	destinations := []v1alpha1.ApplicationDestination{
		{
			Server:    defaultServer,
			Namespace: "argocd",
		},
	}

	return c.createGalahProject(ctx,
		"galah-crds",
		"custom resource definitions",
		sources,
		namespaces,
		destinations,
	)
}

func (c *Client) createGalahProject(ctx context.Context, name, description string, sources, namespaces []string,
	destinations []v1alpha1.ApplicationDestination) (string, error) {
	proj := &v1alpha1.AppProject{}

	projectWithName(name)(proj)
	projectWithDescription(description)(proj)
	projectWithSources(sources)(proj)
	projectWithSourceNamespaces(namespaces)(proj)
	projectWithDestinations(destinations)(proj)
	projectWithNamespace("argocd")(proj)

	err := c.createProject(ctx, proj)
	if err != nil {
		return "", err
	}
	return proj.Name, nil
}

func (c *Client) createAlloyApp(ctx context.Context, project string) error {
	return c.addMonitoringInfrastructureApp(ctx, alloyPath, project)
}

func (c *Client) createDashboardApp(ctx context.Context, project string) error {
	return c.addMonitoringInfrastructureApp(ctx, userPath, project)
}

func (c *Client) createGatewayApp(ctx context.Context, project string) error {
	return c.addMonitoringInfrastructureApp(ctx, gatewayPath, project)
}

func (c *Client) createGrafanaApp(ctx context.Context, project string) error {
	return c.addMonitoringInfrastructureApp(ctx, grafanaPath, project)
}

func (c *Client) createLokiApp(ctx context.Context, project string) error {
	return c.addMonitoringInfrastructureApp(ctx, lokiPath, project)
}

func (c *Client) createMimirApp(ctx context.Context, project string) error {
	return c.addMonitoringInfrastructureApp(ctx, mimirPath, project)
}

func (c *Client) createMinioOperatorApp(ctx context.Context, project string) error {
	return c.addMonitoringInfrastructureApp(ctx, minioOperatorPath, project)
}

func (c *Client) createMinioTenantApp(ctx context.Context, project string) error {
	return c.addMonitoringInfrastructureApp(ctx, minioTenantPath, project)
}

func (c *Client) createTempoApp(ctx context.Context, project string) error {
	return c.addMonitoringInfrastructureApp(ctx, tempoPath, project)
}

func (c *Client) addMonitoringInfrastructureApp(ctx context.Context, path string, project string) error {

	return nil
}
