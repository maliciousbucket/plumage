package types

import (
	"context"
	"github.com/compose-spec/compose-go/v2/cli"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestParseServiceConfig(t *testing.T) {

	t.Run("valid config", func(t *testing.T) {

	})

	t.Run("invalid config", func(t *testing.T) {

	})

	t.Run("test sample config", func(t *testing.T) {

	})
}

func TestParseSampleComposeFile(t *testing.T) {
	containerPort := float64(8080)
	publishedPort := float64(8080)
	protocol := ProtocolTCP

	storefront := ContainerSpec{
		Name:  "store-front",
		Image: "my-image",
		Ports: []*Port{
			&Port{
				Name:          nil,
				ContainerPort: &containerPort,
				PublishedPort: &publishedPort,
				Protocol:      &protocol,
			},
		},
		Resources:     nil,
		CommandProbes: nil,
		HttpProbes:    nil,
		Volumes:       nil,
		Commands:      nil,
	}

	t.Run("booking microservice - store front", func(t *testing.T) {
		composeFilePath := "../../testdata/compose/aks/docker-compose.yml"
		projectName := "testproject"
		ctx := context.Background()

		options, err := cli.NewProjectOptions(
			[]string{composeFilePath},
			cli.WithName(projectName),
		)
		if err != nil {
			t.Fatal(err)
		}

		project, err := options.LoadProject(ctx)
		if err != nil {
			t.Fatal(err)
		}

		svc, err := project.GetService("store-front")
		if err != nil {
			t.Fatal(err)
		}

		container, err := ParseServiceConfig(svc)
		require.NoError(t, err)

		require.Equal(t, container.Name, "store-front")
		assertContainerSpecEqual(t, container, &storefront)
	})
}

func assertContainerSpecEqual(t *testing.T, got *ContainerSpec, want *ContainerSpec) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
