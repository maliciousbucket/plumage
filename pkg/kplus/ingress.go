package kplus

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/maliciousbucket/plumage/imports/traefikio"
	"github.com/maliciousbucket/plumage/pkg/plumage-template/ingress"
)

func NewIngressRoute(scope constructs.Construct, id string, ns string, service *ServiceTemplate, middleware []string) traefikio.IngressRoute {
	var paths []*ingress.ServicePaths
	for _, path := range service.Paths {
		paths = append(paths, &ingress.ServicePaths{
			Path: path.Path,
			Port: path.Port,
		})
	}

	props := &ingress.RouteProps{
		Name:      service.Name,
		Namespace: ns,
		Config: &ingress.RouteConfig{
			Host:               "",
			Paths:              paths,
			EnableLoadBalancer: false,
		},
		Middlewares: middleware,
		HealthCheck: nil,
	}

	ingressRoute := ingress.NewServiceIngressRoute(scope, id, props, middleware)
	return ingressRoute
}
