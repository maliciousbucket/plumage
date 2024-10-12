package kubernetes

import (
	"fmt"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/maliciousbucket/plumage/imports/traefikio"
)

func grafanaRoute(scope constructs.Construct, ns string) traefikio.IngressRoute {
	rule := fmt.Sprintf("PathPrefix(`/%s`)", "grafana")
	return traefikio.NewIngressRoute(scope, jsii.String("grafana"), &traefikio.IngressRouteProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Name: jsii.String("grafana-route"),
		},
		Spec: &traefikio.IngressRouteSpec{
			EntryPoints: &[]*string{
				jsii.String("web"),
			},
			Routes: &[]*traefikio.IngressRouteSpecRoutes{
				{
					Kind:        traefikio.IngressRouteSpecRoutesKind_RULE,
					Match:       jsii.String(rule),
					Middlewares: &[]*traefikio.IngressRouteSpecRoutesMiddlewares{},
					Services: &[]*traefikio.IngressRouteSpecRoutesServices{
						{
							Name:      jsii.String("grafana"),
							Port:      traefikio.IngressRouteSpecRoutesServicesPort_FromNumber(jsii.Number(3000)),
							Namespace: jsii.String("galah-monitoring"),
						},
					},
				},
			},
		},
	})
}
