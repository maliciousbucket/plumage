package orchestration

import (
	"context"
	"log"
)

func Setup(ctx context.Context, helmClient HelmClient, kubeClient KubeClient, ns, argoVersion, valuesFile, envFIle string) error {
	_, err := kubeClient.CreateNamespace(ctx, ns)
	if err != nil {
		return err
	}
	_, err = kubeClient.CreateNamespace(ctx, "argocd")
	if err != nil {
		return err
	}

	err = helmClient.InstallArgoChart(ctx, argoVersion, valuesFile)
	if err != nil {
		return err
	}

	err = kubeClient.SetupArgoLb(ctx, ns, envFIle)
	if err != nil {
		return err
	}
	_, err = kubeClient.CheckArgoExists(ctx, "argocd")
	if err != nil {
		return err
	}

	err = SetArgoToken(ctx, kubeClient)
	if err != nil {
		return err
	}
	log.Println("ArgoCD has been set up")

	return nil
}
