package integration_testing

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go/modules/k3s"
	"io"
	"strings"
	"testing"
	"time"
)

type CheckResourceOpts struct {
	Type      string
	Name      string
	Namespace string
}

func AssertResourceReady(t *testing.T, ctx context.Context, client *kubeClient, opts *CheckResourceOpts) error {
	t.Helper()

	return nil
}

func AssertResourceExists(t *testing.T, ctx context.Context, opts *CheckResourceOpts, ctr *k3s.K3sContainer) error {
	t.Helper()

	return nil
}

type WaitResourceOpts struct {
	Type      string
	Name      string
	Namespace string
	Condition string
	Timeout   time.Duration
}

func WaitForResource(t *testing.T) {
	t.Helper()
}

func WaitForDeployment(t *testing.T, ctx context.Context, client kubeClient, kubeContainer *KubernetesContainer, name, ns string, meta bool) {
	t.Helper()

	err := client.WatchDeployment(ctx, ns, name, meta)
	assert.NoError(t, err)
}

func DescribeResource(t *testing.T, ctx context.Context, kubeContainer *KubernetesContainer, resource, ns string) {
	t.Helper()

	res, reader, err := kubeContainer.Container.Exec(ctx, []string{"kubectl", "describe", resource, "-n", ns})

	assert.NoError(t, err)

	if res != 0 {
		t.Errorf("expecting 0 but got %d", res)
	}

	buf := new(strings.Builder)
	_, err = io.Copy(buf, reader)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(buf.String())

}
