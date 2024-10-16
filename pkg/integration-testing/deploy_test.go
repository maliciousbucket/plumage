package integration_testing

import (
	"context"
	"github.com/maliciousbucket/plumage/internal/kubeclient"
	"github.com/maliciousbucket/plumage/internal/orchestration"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

var (
	testBedNamespace = "galah-testbed"
)

// TODO
func TestDeployTemplate(t *testing.T) {

}

func TestDeployGateway(t *testing.T) {
	kubeContainer := NewKubernetesContainer()
	err := WithHostPorts(30443)(kubeContainer)
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()
	if err = SetupKubernetesContainer(ctx, kubeContainer); err != nil {
		t.Fatalf("error setting up clients: %v", err)
	}
	defer func() {
		if err = kubeContainer.Container.Terminate(ctx); err != nil {
			t.Fatal(err)
		}
	}()

	helm, kube, err := setupHelmKubeClients(ctx, kubeContainer)
	if err != nil {
		t.Fatal(err)
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
			t.Fatal("error creating Argo Client")
		}
		argo = argoC
	})

	t.Run("create test bed namespace", func(t *testing.T) {
		err = orchestration.CreateNamespace(ctx, kube, testBedNamespace)
		assert.NoErrorf(t, err, "error creating namespace: %v", err)
	})

	t.Run("set argo git credentials", func(t *testing.T) {
		err = setArgoCreds(ctx, argo)
		assert.NoErrorf(t, err, "error setting git credentials: %v", err)
	})

	t.Run("deploy gateway", func(t *testing.T) {
		if err = orchestration.DeployGateway(ctx, argo, kube, testBedNamespace); err != nil {
			t.Fatalf("error deploying gateway: %v", err)
		}
	})

	log.Println("yay")

}

func TestDeployMonitoring(t *testing.T) {
	kubeContainer := NewKubernetesContainer()
	err := WithHostPorts(30443)(kubeContainer)
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()
	if err = SetupKubernetesContainer(ctx, kubeContainer); err != nil {
		t.Fatalf("error setting up clients: %v", err)
	}
	defer func() {
		if err = kubeContainer.Container.Terminate(ctx); err != nil {
			t.Fatal(err)
		}
	}()

	helm, kube, err := setupHelmKubeClients(ctx, kubeContainer)
	if err != nil {
		t.Fatal(err)
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
			t.Fatal("error creating Argo Client")
		}
		argo = argoC
	})

	t.Run("set argo git credentials", func(t *testing.T) {
		err = setArgoCreds(ctx, argo)
		assert.NoErrorf(t, err, "error setting git credentials: %v", err)
	})

	t.Run("deploy monitoring", func(t *testing.T) {
		if err = orchestration.DeployMonitoring(ctx, argo, kube); err != nil {
			t.Fatalf("error deploying gateway: %v", err)
		}
	})
	time.Sleep(1 * time.Minute)
	DescribeResource(t, ctx, kubeContainer, "pods", "galah-monitoring")
	DescribeResource(t, ctx, kubeContainer, "pods", "galah-logging")
	DescribeResource(t, ctx, kubeContainer, "deployment", "galah-monitoring")
	DescribeResource(t, ctx, kubeContainer, "deployment", "galah-logging")

	t.Run("wait for monitoring deployment", func(t *testing.T) {
		err = orchestration.WaitForMonitoringDeployment(ctx, kube)
		assert.NoErrorf(t, err, "error waiting for monitoring deployment: %v", err)
	})

	log.Println("yay")
}

func setupHelmKubeClients(ctx context.Context, kubeContainer *KubernetesContainer) (helmClient, kubeClient, error) {
	config, err := kubeContainer.KubeConfig(ctx)
	if err != nil {
		return nil, nil, err
	}

	helm, err := newHelmClient(config)
	if err != nil {
		return nil, nil, err
	}

	kube, err := kubeclient.NewClientFromConfig(config)
	if err != nil {
		return nil, nil, err
	}

	return helm, kube, nil
}
