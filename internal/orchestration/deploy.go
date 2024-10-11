package orchestration

import (
	"context"
	"errors"
	"github.com/maliciousbucket/plumage/internal/argocd"
	"github.com/maliciousbucket/plumage/internal/kubeclient"
	"sync"
	"time"
)

func DeployTemplate() {}

func DeployGateway(ctx context.Context, argoClient ArgoClient, kubeClient KubeClient, ns string) error {
	return deployGateway(ctx, argoClient, kubeClient, ns)
}

func deployGateway(ctx context.Context, argoClient ArgoClient, kubeClient KubeClient, ns string) error {
	if gatewayProj, _ := argoClient.GetProject(ctx, "ingress"); gatewayProj == nil {
		if err := argoClient.CreateIngressProject(ctx); err != nil {
			return err
		}
	} else {
		params := &argocd.AppQueryParams{Options: []argocd.AppQueryFunc{
			argocd.WithProject("ingress"),
		}}
		apps, _ := argoClient.ListApplications(ctx, params)
		if apps == nil || len(apps.Items) == 0 {
			if err := argoClient.CreateIngressApp(ctx); err != nil {
				return err
			}
		}

		if err := argoClient.SyncProject(ctx, "ingress"); err != nil {
			return err
		}
	}
	time.Sleep(10 * time.Second)
	errChan := make(chan error)
	go func() {

		errChan <- watchInfrastructure(ctx, kubeClient, ns, "traefik")
		close(errChan)
	}()
	select {
	case err := <-errChan:
		return err
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

func deployAndWaitForApp(ctx context.Context, argoClient ArgoClient, kubeClient KubeClient, ns, app string, services []string) error {
	if err := argoClient.CreateServiceApplications(ctx, app, services); err != nil {
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
