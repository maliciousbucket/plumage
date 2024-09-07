package kubernetes

import (
	"fmt"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	kplus "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2"
	"github.com/maliciousbucket/plumage/pkg/template"
)

func defaultServiceAnnotations() *map[string]*string {
	return nil
}

func defaultServiceLabels(svcName string) *map[string]*string {
	labels := map[string]*string{
		"app": jsii.String(svcName),
	}
	return &labels
}

func DefaultServiceMetadata(nameSpace, svcName string) cdk8s.ApiObjectMetadata {

	return cdk8s.ApiObjectMetadata{
		Name:      jsii.String(svcName),
		Namespace: jsii.String(nameSpace),
		Labels:    defaultMiddlewareLabels(svcName),
	}
}

func ServicePorts(config *template.ServiceConfig) []*kplus.ServicePort {
	if config.Service != nil {
		if config.Service.Monitoring != nil {
			if config.Service.Monitoring.ScrapeConfig != nil {
				if config.Service.Monitoring.ScrapeConfig.MetricsPort != 0 {
					metricsPort := jsii.Number(config.Service.Monitoring.ScrapeConfig.MetricsPort)
					for _, port := range config.Container.Ports {
						if port.ContainerPort == metricsPort {
							port.Name = jsii.String("http-metrics")
						}
					}
				}
			}
		}

	}

	var ports []*kplus.ServicePort
	i := 0
	for _, port := range config.Container.Ports {
		name := port.Name
		if name == nil || *name == "" {
			nameFmt := fmt.Sprintf("http-%d", i)
			name = jsii.String(nameFmt)
		}
		ports = append(ports, &kplus.ServicePort{
			Name:       name,
			Port:       port.PublishedPort,
			TargetPort: port.ContainerPort,
			Protocol:   port.Cdk8sProtocol(),
		})

	}

	return ports
}
