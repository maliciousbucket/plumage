package integration_testing

import (
	"context"
	"errors"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/k3s"
	"github.com/testcontainers/testcontainers-go/wait"
	"time"
)

var (
	k3sImage = "docker.io/rancher/k3s:v1.27.1-k3s1"
)

type KubernetesContainer struct {
	Container *k3s.K3sContainer
	Options   []testcontainers.ContainerCustomizer
}

func NewKubernetesContainer() *KubernetesContainer {
	return &KubernetesContainer{Options: []testcontainers.ContainerCustomizer{}}
}

func SetupKubernetesContainer(ctx context.Context, c *KubernetesContainer) error {

	k3sContainer, err := k3s.Run(ctx, "docker.io/rancher/k3s:v1.27.1-k3s1", c.Options...)
	if err != nil {
		return fmt.Errorf("error creating k3s container: %w", err)
	}

	if k3sContainer != nil {
		c.Container = k3sContainer
	} else {
		return fmt.Errorf("k3s container not found")
	}

	return nil

}

func (c *KubernetesContainer) KubeConfig(ctx context.Context) ([]byte, error) {
	if c.Container == nil {
		return nil, fmt.Errorf("k3s container not found")
	}
	config, err := c.Container.GetKubeConfig(ctx)
	if err != nil {
		return nil, err
	}
	return config, nil
}

type KubeContainerOption func(c *KubernetesContainer) error

func WithManifests(manifests []string) KubeContainerOption {
	return func(c *KubernetesContainer) error {
		if manifests == nil {
			return fmt.Errorf("manifests is nil")
		}
		for _, manifest := range manifests {
			c.Options = append(c.Options, k3s.WithManifest(manifest))
		}
		return nil
	}
}

type WaitCondition struct {
	Resource  string
	Name      string
	Namespace string
	Condition string
	Timeout   time.Duration
}

func (w *WaitCondition) Validate() error {
	var err error
	if w.Resource == "" {
		err = errors.Join(err, errors.New("resource is empty"))
	}
	if w.Name == "" {
		err = errors.Join(err, errors.New("name is empty"))
	}
	if w.Condition == "" {
		err = errors.Join(err, errors.New("condition is empty"))
	}
	if w.Timeout == 0 {
		err = errors.Join(err, errors.New("timeout is empty"))
	}
	return err
}

func WithWaitStrategy(condition *WaitCondition) KubeContainerOption {
	return func(c *KubernetesContainer) error {
		if err := condition.Validate(); err != nil {
			return err
		}
		strategy := []string{"kubectl", "wait"}
		strategy = append(strategy, condition.Resource, condition.Name)
		if condition.Namespace != "" {
			namespace := fmt.Sprintf("--namespace=%s", condition.Namespace)
			strategy = append(strategy, namespace)
		}
		timeout := fmt.Sprintf("--timeout=%s", condition.Timeout)
		strategy = append(strategy, timeout)
		c.Options = append(c.Options, testcontainers.WithWaitStrategy(wait.ForExec(strategy)))
		return nil
	}
}
