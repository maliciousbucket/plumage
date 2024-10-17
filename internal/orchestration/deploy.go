package orchestration

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/maliciousbucket/plumage/internal/argocd"
	"github.com/maliciousbucket/plumage/internal/kubeclient"
	"github.com/maliciousbucket/plumage/pkg/kplus"
	"log"
	"slices"
	"sync"
	"time"
)

var (
	chaosPath = "/tests"
)

func DeployTemplate() {}

func deployTemplate() {

}

func DeployApp(ctx context.Context, argoClient ArgoClient, ns, path string) (string, error) {
	return deployApp(ctx, argoClient, ns, path)
}

func deployApp(ctx context.Context, argoClient ArgoClient, ns, path string) (string, error) {
	if err := AddRepoCredentials(ctx, argoClient, ""); err != nil {
		return "", err
	}

	services, _, appName, err := kplus.GetServices(path)
	if err != nil {
		return "", err
	}
	if len(services) == 0 {
		return "", fmt.Errorf("no services found in %s", path)
	}

	if err = argoClient.CreateServiceApplications(ctx, ns, appName, services); err != nil {
		return "", err
	}
	return appName, nil

}

func DeployGateway(ctx context.Context, argoClient ArgoClient, kubeClient KubeClient, ns string) error {
	return deployGateway(ctx, argoClient, kubeClient, ns)
}

func deployGateway(ctx context.Context, argoClient ArgoClient, kubeClient KubeClient, ns string) error {
	if gatewayProj, _ := argoClient.GetProject(ctx, "ingress"); gatewayProj == nil {
		if err := argoClient.CreateIngressProject(ctx, ns); err != nil {
			return err
		}
	} else {
		params := &argocd.AppQueryParams{Options: []argocd.AppQueryFunc{
			argocd.WithProject("ingress"),
		}}
		apps, _ := argoClient.ListApplications(ctx, params)
		if apps == nil || len(apps.Items) == 0 {
			if err := argoClient.CreateIngressApp(ctx, ns); err != nil {
				return err
			}
		}

	}
	if err := argoClient.SyncProject(ctx, "ingress"); err != nil {
		return err
	}
	projects, err := argoClient.ListProjects(ctx)
	if err != nil {
		return err
	}
	if err = prettyPrint(projects); err != nil {
		return err
	}

	apps, err := argoClient.ListApplications(ctx, &argocd.AppQueryParams{Options: []argocd.AppQueryFunc{
		argocd.WithProject("ingress"),
	}})
	if err != nil {
		return err
	}
	fmt.Println("Apps?")
	if err = prettyPrint(apps); err != nil {
		return err
	}
	log.Println("Pods???")
	pods, err := kubeClient.ListPods(ctx, ns)
	if err != nil {
		return err
	}
	log.Println(len(pods.Items))
	if err = prettyPrint(pods); err != nil {
		return err
	}
	log.Println("Deployments?")
	deployments, err := kubeClient.ListDeployments(ctx, ns)
	if err != nil {
		return err
	}
	if err = prettyPrint(deployments); err != nil {
		return err
	}

	return nil
}

func WaitForGatewayDeployment(ctx context.Context, kubeClient KubeClient, ns string) error {
	return waitForGatewayDeployment(ctx, kubeClient, ns)
}

func waitForGatewayDeployment(ctx context.Context, kubeClient KubeClient, ns string) error {
	err := kubeClient.WatchDeployment(ctx, ns, "traefik", true)
	if err != nil {
		return err
	}
	errChan := make(chan error)
	go func() {

		errChan <- watchInfrastructure(ctx, kubeClient, ns, "traefik")
		close(errChan)
	}()
	select {
	case watchErr := <-errChan:
		return watchErr
	case <-ctx.Done():
		return ctx.Err()
	}
}

func DeployMonitoring(ctx context.Context, argoClient ArgoClient, kubeClient KubeClient) error {
	return deployMonitoring(ctx, argoClient, kubeClient)
}

func deployMonitoring(ctx context.Context, argoClient ArgoClient, kubeClient KubeClient) error {
	if monitoringProj, _ := argoClient.GetProject(ctx, "galah-monitoring"); monitoringProj == nil {
		if err := argoClient.CreateMonitoringProject(ctx); err != nil {
			return err
		}
	} else {
		if err := argoClient.SyncProject(ctx, "galah-monitoring"); err != nil {
			return err
		}
	}

	return nil

}

func WaitForMonitoringDeployment(ctx context.Context, kubeClient KubeClient) error {
	return waitForMonitoringDeployment(ctx, kubeClient)
}

func waitForMonitoringDeployment(ctx context.Context, kubeClient KubeClient) error {
	resources := map[string]string{
		"alloy":    "galah-monitoring",
		"tempo":    "galah-tracing",
		"loki":     "galah-logging",
		"mimir":    "galah-monitoring",
		"grafana":  "galah-monitoring",
		"operator": "minio-store",
	}
	time.Sleep(10 * time.Second)
	var watchErr error
	errChan := make(chan error, len(resources)+1)
	var wg sync.WaitGroup
	time.Sleep(1 * time.Minute)
	for res, namespace := range resources {
		wg.Add(1)
		go func(res, namespace string) {
			defer wg.Done()
			err := watchInfrastructure(ctx, kubeClient, namespace, res)
			errChan <- err
		}(res, namespace)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for {
		select {
		case err, ok := <-errChan:
			if ok {
				if err != nil {
					watchErr = errors.Join(watchErr, err)
				}
			} else {
				return watchErr
			}
		case <-time.After(time.Second * 30):
			ctx.Done()
			wg.Done()
			return watchErr
		case <-ctx.Done():
			watchErr = errors.Join(watchErr, ctx.Err())
			return watchErr
		}
	}
}

// DeployChaos TODO: Wait for k6
func DeployChaos(ctx context.Context, argoClient ArgoClient, kubeClient KubeClient, helmClient HelmClient, ns, version, outputDir string) error {
	if err := helmClient.InstallK6(ctx, version, "", false); err != nil {
		return fmt.Errorf("install k6: %w", err)
	}

	chaosProj, _ := argoClient.GetProject(ctx, "chaos")
	if chaosProj == nil {
		if err := argoClient.CreateChaosProject(ctx, ns); err != nil {
			return err
		}
	} else {
		if err := argoClient.SyncProject(ctx, chaosProj.Name); err != nil {
			return err
		}
	}

	tests, err := getChaosTests(outputDir, chaosPath)
	if err != nil {
		return err
	}
	if len(tests) == 0 {
		return errors.New("no tests found")
	}
	apps, err := argoClient.ListApplications(ctx, &argocd.AppQueryParams{Options: []argocd.AppQueryFunc{
		argocd.WithProject("chaos"),
	}})
	if err != nil {
		return err
	}
	if apps != nil && len(apps.Items) > 0 {
		for _, app := range apps.Items {
			if slices.Contains(tests, app.Name) {
				for i, test := range tests {
					if test == app.Name {
						tests = append(tests[:i], tests[i+1:]...)
						break
					}
				}
			}
		}
	}
	if len(tests) == 0 {
		log.Println("No new tests found")
		return nil
	}
	path := fmt.Sprintf("%s/%s", outputDir, "tests")
	if err = argoClient.CreateChaosApp(ctx, ns, "chaos", path); err != nil {
		return err
	}
	//TODO: wait for deployment
	return nil

}

func CreateNamespace(ctx context.Context, kubeClient KubeClient, ns string) error {
	_, err := kubeClient.CreateNamespace(ctx, ns)
	if err != nil {
		return err
	}
	return nil
}

func DeployAndWaitForApp(ctx context.Context, argoClient ArgoClient, kubeClient KubeClient, ns, app string, services []string) error {
	if err := argoClient.CreateServiceApplications(ctx, ns, app, services); err != nil {
		return err
	}

	if err := argoClient.SyncProject(ctx, app); err != nil {
		return err
	}
	if err := kubeClient.WatchAppDeployment(ctx, ns, services); err != nil {
		return err
	}
	return nil
}

func watchInfrastructure(ctx context.Context, client kubeclient.Client, ns, name string) error {
	if err := client.WaitAppPods(ctx, ns, name, 1, 2*time.Minute); err != nil {
		return err
	}
	return nil
}

var (
	argoServer        = "argo-helm-argocd-server"
	argoAppController = "argo-helm-argocd-applicationset-controller"
	argoDexServer     = "argo-helm-argocd-dex-server"
	argoRepoServer    = "argo-helm-argocd-repo-server"
)

func WatchArgoDeployment(ctx context.Context, kubeClient KubeClient) error {
	var wg sync.WaitGroup
	errChan := make(chan error, 4)

	waitArgoResource := func(resource string) {
		err := kubeClient.WatchDeployment(ctx, "argocd", resource, false)
		errChan <- err
		wg.Done()
	}
	wg.Add(4)
	go waitArgoResource(argoServer)
	go waitArgoResource(argoAppController)
	go waitArgoResource(argoDexServer)
	go waitArgoResource(argoRepoServer)

	go func() {
		wg.Wait()
		close(errChan)
	}()
	var watchErr error

	for err := range errChan {
		if err != nil {
			watchErr = errors.Join(watchErr, err)
		}
	}
	return watchErr

}

func prettyPrint(v any) error {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(jsonData))
	return nil
}
