package types

import (
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

//
//func TestParseSampleComposeFile(t *testing.T) {
//	containerPort := float64(8080)
//	publishedPort := float64(8080)
//	protocol := ProtocolTCP
//
//	cmd1 := "CMD"
//	cmd2 := "wget"
//	cmd3 := "-O"
//	cmd4 := "/dev/null"
//	cmd5 := "-q"
//	cmd6 := "http://store-front:80/health"
//
//	timeout := "10s"
//	interval := "30s"
//	retries := float64(5)
//
//	storefront := ContainerSpec{
//		Name:  "store-front",
//		Image: "my-image",
//		Ports: []*Port{
//			&Port{
//				Name:          nil,
//				ContainerPort: &containerPort,
//				PublishedPort: &publishedPort,
//				Protocol:      &protocol,
//			},
//		},
//		Resources: nil,
//		HealthCheck: &CommandProbe{
//
//			Commands: []*string{
//				&cmd1,
//				&cmd2,
//				&cmd3,
//				&cmd4,
//				&cmd5,
//				&cmd6,
//			},
//			Delay:    nil,
//			Timeout:  &timeout,
//			Interval: &interval,
//			Retries:  &retries,
//		},
//
//		ReadinessProbe: nil,
//		Volumes:        nil,
//		Commands:       nil,
//	}
//
//	t.Run("aks microservice - store front", func(t *testing.T) {
//		composeFilePath := "../../testdata/compose/aks/docker-compose.yml"
//		projectName := "testproject"
//		ctx := context.Background()
//
//		options, err := cli.NewProjectOptions(
//			[]string{composeFilePath},
//			cli.WithName(projectName),
//		)
//		if err != nil {
//			t.Fatal(err)
//		}
//
//		project, err := options.LoadProject(ctx)
//		if err != nil {
//			t.Fatal(err)
//		}
//
//		svc, err := project.GetService("store-front")
//		if err != nil {
//			t.Fatal(err)
//		}
//
//		container, err := ParseServiceConfig(svc)
//		require.NoError(t, err)
//
//		require.Equal(t, container.Name, "store-front")
//		assertContainerSpecEqual(t, container, &storefront)
//	})
//}

func assertContainerSpecEqual(t *testing.T, got *ContainerSpec, want *ContainerSpec) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
