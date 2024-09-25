package orchestration

import (
	"context"
	"errors"
	"github.com/maliciousbucket/plumage/internal/kubeclient"
)

func CheckMonitoringInfraExists(ctx context.Context, c kubeclient.Client) error {
	argo, err := c.CheckServiceExists(ctx, "argocd", "argocd-helm-server")
	if err != nil {
		return err
	}
	if !argo {
		return errors.New("argocd is not installed")
	}
	alloy, err := c.CheckServiceExists(ctx, "galah-monitoring", "alloy")
	if err != nil {
		return err
	}
	if !alloy {
		return errors.New("alloy is not installed")
	}

	grafana, err := c.CheckServiceExists(ctx, "galah-monitoring", "grafana")
	if err != nil {
		return err
	}
	if !grafana {
		return errors.New("galah-monitoring is not installed")
	}
	return nil
}
