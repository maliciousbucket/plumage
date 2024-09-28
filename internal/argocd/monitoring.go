package argocd

import (
	"context"
	"fmt"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/application"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"slices"
)

const (
	galahMonitoringRepo = "https://github.com/maliciousbucket/galah-observability"
	argoCDRepo          = "https://github.com/maliciousbucket/github-api-test"
	defaultServer       = "https://kubernetes.default.svc"
	alloyPath           = "kubernetes/alloy"
	grafanaPath         = "kubernetes/grafana"
	lokiPath            = "kubernetes/loki"
	mimirPath           = "kubernetes/mimir"
	minioOperatorPath   = "kubernetes/minio-operator"
	minioTenantPath     = "kubernetes/minio-tenant"
	tempoPath           = "kubernetes/tempo/helm"
	userPath            = "kubernetes/user"

	gatewayPath = "kubernetes/gateway"
)

func (c *Client) CreateMonitoringProject(ctx context.Context) error {
	clusters, err := c.GetClusters(ctx)
	if err != nil {
		return err
	}

	if clusters == nil {
		return fmt.Errorf("no clusters found")
	}
	createCluster := true
	for _, cluster := range clusters {
		if cluster.Info.ServerVersion != "Unknown" {
			createCluster = false
		}
		fmt.Println(cluster.Name)
		fmt.Println(cluster.Server)
		fmt.Println(cluster.Info)
		fmt.Println(cluster.Info.ConnectionState.Status)
		fmt.Println(cluster.Info.ServerVersion)
	}
	if createCluster {
		cluster, clusterErr := c.CreateCluster(ctx, "galah-monitoring")
		if clusterErr != nil {
			return fmt.Errorf("failed to create galah-monitoring cluster: %v", clusterErr)
		}
		fmt.Println(cluster.Name)
	}
	var project string
	existing, _ := c.GetProject(ctx, "galah-monitoring")

	if existing != nil {
		log.Println("Galah monitoring project already exists")
		project = existing.Name
	} else {
		newProject, projErr := c.createMonitoringProject(ctx)
		if projErr != nil {
			log.Println("Create project error")
			return projErr
		}
		project = newProject
	}

	apps, err := c.ListApplications(ctx, nil)
	if err != nil {
		fmt.Println("List applications error")
		return err
	}

	appNames := []string{}
	if apps != nil {
		for _, app := range apps.Items {
			appNames = append(appNames, app.Name)
		}
	}

	createAlloy := true
	createGateway := true
	createGrafana := true
	createLoki := true
	createMimir := true
	createMinioOperator := true
	createMinioTenant := true
	createTempo := true

	if len(appNames) > 0 {
		if slices.Contains(appNames, "alloy-app") {
			createAlloy = false
		}

		if slices.Contains(appNames, "gateway-app") {
			createGateway = false
		}

		if slices.Contains(appNames, "grafana-app") {
			createGrafana = false
		}

		if slices.Contains(appNames, "loki-app") {
			createLoki = false
		}
		if slices.Contains(appNames, "mimir-app") {
			createMimir = false
		}
		if slices.Contains(appNames, "minio-operator-app") {
			createMinioOperator = false
		}
		if slices.Contains(appNames, "minio-tenant-app") {
			createMinioTenant = false
		}
		if slices.Contains(appNames, "tempo-app") {
			createTempo = false
		}

	}

	if createAlloy {
		if err = c.createAlloyApp(ctx, project); err != nil {
			fmt.Println("Create alloy app error")
			return err
		}

	}
	if createGateway {
		if err = c.createGatewayApp(ctx, project); err != nil {
			return fmt.Errorf("error creating gateway application: %w", err)
		}

	}

	if createGrafana {
		if err = c.createGrafanaApp(ctx, project); err != nil {
			return fmt.Errorf("error creating grafana application: %w", err)
		}
	}

	if createLoki {
		if err = c.createLokiApp(ctx, project); err != nil {
			return fmt.Errorf("error creating loki application: %w", err)
		}
	}

	if createMimir {
		if err = c.createMimirApp(ctx, project); err != nil {
			return fmt.Errorf("error creating mimir application: %w", err)
		}
	}

	if createMinioOperator {
		if err = c.createMinioOperatorApp(ctx, project); err != nil {
			return fmt.Errorf("error creating minio-operator application: %w", err)
		}

	}

	if createMinioTenant {
		if err = c.createMinioTenantApp(ctx, project); err != nil {
			return fmt.Errorf("error creating minio-tenant application: %w", err)
		}
	}

	if createTempo {

		if err = c.createTempoApp(ctx, project); err != nil {
			return fmt.Errorf("error creating tempo application: %w", err)
		}
	}

	return nil

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
		{
			Namespace: "galah-monitoring",
			Server:    defaultServer,
		},
		{
			Namespace: "galah-tracing",
			Server:    defaultServer,
		},
		{
			Namespace: "galah-logging",
			Server:    defaultServer,
		},
		{
			Namespace: "minio-store",
			Server:    defaultServer,
		},
		{
			Namespace: "gateway",
			Server:    defaultServer,
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

func (c *Client) CreateNetworkingProject(ctx context.Context) error {

	var project string
	existing, _ := c.GetProject(ctx, "galah-monitoring")

	if existing != nil {
		log.Println("Galah networking project already exists")
		project = existing.Name
	} else {
		newProject, err := c.createNetworkingProject(ctx)
		if err != nil {
			return err
		}
		project = newProject
	}

	apps, err := c.ListApplications(ctx, nil)
	if err != nil {
		fmt.Println("List applications error")
		return err
	}

	appNames := []string{}
	if apps != nil {
		for _, app := range apps.Items {
			appNames = append(appNames, app.Name)
		}
	}

	createGateway := true

	if slices.Contains(appNames, "gateway-app") {
		createGateway = false
	}

	if createGateway {
		if err = c.createGatewayApp(ctx, project); err != nil {
			return err
		}
	}
	log.Println("Create Networking Project")
	return nil

}

func (c *Client) createNetworkingProject(ctx context.Context) (string, error) {
	sources := []string{galahMonitoringRepo, argoCDRepo}
	namespaces := []string{
		"gateway",
		"galah-monitoring",
	}

	destinations := []v1alpha1.ApplicationDestination{
		{
			Server:    defaultServer,
			Namespace: "argocd",
		},
		{
			Namespace: "galah-monitoring",
			Server:    defaultServer,
		},
		{
			Namespace: "gateway",
			Server:    defaultServer,
		},
		{
			Namespace: "galah-testbed",
			Server:    defaultServer,
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

//func (c *Client) CreateCRDProject(ctx context.Context) error {
//	existing, _ := c.GetProject(ctx, "galah-crds")
//	if existing != nil {
//		return fmt.Errorf("galah-crd project already exists")
//	}
//
//	appList := []string{}
//	apps, err := c.ListApplications(ctx, nil)
//	if err != nil {
//		return fmt.Errorf("error listing applications: %w", err)
//	}
//
//	certManager := true
//	promCRds := true
//
//
//	project, err := c.createCRDProject(ctx)
//	if err != nil {
//		return err
//	}
//	log.Printf("Successfully created CRD Project: %s", project)
//	return nil
//}

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
		{
			Namespace: "default",
			Server:    defaultServer,
		},
		{
			Namespace: "cert-manager",
			Server:    defaultServer,
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
	proj := v1alpha1.AppProject{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		}, Spec: v1alpha1.AppProjectSpec{
			ClusterResourceWhitelist: []metav1.GroupKind{
				{
					Group: "*",
					Kind:  "*",
				},
			},
		},
	}

	projectWithName(name)(&proj)
	projectWithDescription(description)(&proj)
	projectWithSources(sources)(&proj)
	projectWithSourceNamespaces(namespaces)(&proj)
	projectWithDestinations(destinations)(&proj)
	projectWithNamespace("argocd")(&proj)
	fmt.Println(proj)

	project, err := c.createProject(ctx, &proj)
	if err != nil {
		return "", err
	}
	return project.Name, nil
}

func (c *Client) createAlloyApp(ctx context.Context, project string) error {

	return c.addMonitoringInfrastructureApp(ctx, "alloy-app", alloyPath, project)
}

func (c *Client) createDashboardApp(ctx context.Context, project string) error {
	return c.addMonitoringInfrastructureApp(ctx, "dashboard-app", userPath, project)
}

func (c *Client) createGatewayApp(ctx context.Context, project string) error {
	return c.addMonitoringInfrastructureApp(ctx, "gateway-app", gatewayPath, project)
}

func (c *Client) createGrafanaApp(ctx context.Context, project string) error {
	return c.addMonitoringInfrastructureApp(ctx, "grafana-app", grafanaPath, project)
}

func (c *Client) createLokiApp(ctx context.Context, project string) error {
	return c.addMonitoringInfrastructureApp(ctx, "loki-app", lokiPath, project)
}

func (c *Client) createMimirApp(ctx context.Context, project string) error {
	return c.addMonitoringInfrastructureApp(ctx, "mimir-app", mimirPath, project)
}

func (c *Client) createMinioOperatorApp(ctx context.Context, project string) error {
	return c.addMonitoringInfrastructureApp(ctx, "mmini-operator-app", minioOperatorPath, project)
}

func (c *Client) createMinioTenantApp(ctx context.Context, project string) error {
	return c.addMonitoringInfrastructureApp(ctx, "minio-tenant-app", minioTenantPath, project)
}

func (c *Client) createTempoApp(ctx context.Context, project string) error {
	return c.addMonitoringInfrastructureApp(ctx, "tempo-app", tempoPath, project)
}

func (c *Client) addMonitoringInfrastructureApp(ctx context.Context, name, path, project string) error {
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
					RepoURL:        galahMonitoringRepo,
					Path:           path,
					TargetRevision: "HEAD",
				},
				Destination: v1alpha1.ApplicationDestination{
					Server:    defaultServer,
					Namespace: "galah-monitoring",
				},
				SyncPolicy: &v1alpha1.SyncPolicy{
					Automated: &v1alpha1.SyncPolicyAutomated{
						SelfHeal: true,
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
	fmt.Printf("app %s/%s created\n", app.GetName(), app.GetNamespace())
	return nil
}
