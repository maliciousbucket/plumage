package kubernetes

import (
	"fmt"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/maliciousbucket/plumage/imports/traefikio"
)

func NewInfraDashboardRoutes(scope constructs.Construct, id, ns string) constructs.Construct {
	chart := cdk8s.NewChart(scope, jsii.String(id), &cdk8s.ChartProps{
		DisableResourceNameHashes: jsii.Bool(true),
		Namespace:                 jsii.String(ns),
	})

	stripInfraRoutesMiddleware(chart)
	alloyRoute(chart)
	grafanaRoute(chart)
	argoRoute(chart)
	return chart
}

func alloyRoute(scope constructs.Construct) traefikio.IngressRoute {
	id := fmt.Sprintf("alloy-route")
	return newInfraRoute(scope, id, "alloy", "alloy", "galah-monitoring", 12345)
}

func grafanaRoute(scope constructs.Construct) traefikio.IngressRoute {
	id := fmt.Sprintf("grafana-route")
	return newInfraRoute(scope, id, "grafana", "grafana", "galah-monitoring", 3000)
}

func argoRoute(scope constructs.Construct) traefikio.IngressRoute {
	id := fmt.Sprintf("argo-route")
	return newInfraRoute(scope, id, "argo", "argocd-helm-server", "argocd", 443)
}

func newInfraRoute(scope constructs.Construct, id string, name, service, ns string, port int) traefikio.IngressRoute {
	rule := fmt.Sprintf("PathPrefix(`/testbed/%s`)", name)
	return traefikio.NewIngressRoute(scope, jsii.String(id), &traefikio.IngressRouteProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Name: jsii.String(id),
			//Namespace: jsii.String(ns),
		},
		Spec: &traefikio.IngressRouteSpec{
			EntryPoints: &[]*string{
				jsii.String("web"),
			},
			Routes: &[]*traefikio.IngressRouteSpecRoutes{
				{
					Kind:  traefikio.IngressRouteSpecRoutesKind_RULE,
					Match: jsii.String(rule),
					Middlewares: &[]*traefikio.IngressRouteSpecRoutesMiddlewares{
						{
							Name: jsii.String("strip-testbed"),
						},
						{
							Name: jsii.String("strip-infra-prefix"),
						},
					},
					Services: &[]*traefikio.IngressRouteSpecRoutesServices{
						{
							Name:      jsii.String(service),
							Port:      traefikio.IngressRouteSpecRoutesServicesPort_FromNumber(jsii.Number(port)),
							Namespace: jsii.String(ns),
						},
					},
				},
			},
		},
	})
}

func stripInfraRoutesMiddleware(scope constructs.Construct) traefikio.Middleware {
	middleware := traefikio.NewMiddleware(scope, jsii.String("strip-infra-prefix"), &traefikio.MiddlewareProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Name: jsii.String("strip-infra-prefix"),
		},
		Spec: &traefikio.MiddlewareSpec{
			StripPrefix: &traefikio.MiddlewareSpecStripPrefix{
				Prefixes: &[]*string{
					jsii.String("/argo"),
					jsii.String("/alloy"),
				},
			},
		},
	})
	return middleware
}
