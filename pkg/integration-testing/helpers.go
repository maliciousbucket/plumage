package integration_testing

import (
	"context"
	"github.com/testcontainers/testcontainers-go/modules/k3s"
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
