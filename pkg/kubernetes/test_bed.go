package kubernetes

import (
	"fmt"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/maliciousbucket/plumage/pkg/template"
)

const (
	TraefikApiVersion     = "traefik.io/v1alpha1"
	TestbedNamespace      = "test-bed"
	GalahApiVersion       = "galah-monitoring.io/v1"
	KubernetesAppsVersion = "apps/v1"
)

func GenerateManifests(scope constructs.Construct, tmp *template.GalahTemplate, namespace string) (cdk8s.Chart, error) {
	configs, err := tmp.ServiceConfigs(namespace)
	if err != nil {
		return nil, err
	}

	chartName := fmt.Sprintf("%s-%s", namespace, "chart")
	chart := cdk8s.NewChart(scope, jsii.String(chartName), &cdk8s.ChartProps{
		DisableResourceNameHashes: nil,
		Labels:                    nil,
		Namespace:                 jsii.String(namespace),
	})

	for _, config := range configs {

		_, err := NewServiceChart(chart, *config, "yolo")
		if err != nil {
			return nil, err
		}
	}
	return chart, nil
}
