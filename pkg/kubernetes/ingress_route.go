package kubernetes

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/maliciousbucket/plumage/imports/traefikio"
)

const (
	defaultRouteType = "Rule"
)

func newIngressRoute(scope constructs.Construct, id string) traefikio.IngressRoute {
	return nil
}

//func newProps(scope constructs.Construct, id string) traefikio.IngressRouteProps {
//	return nil
//}

//func newSpec(scope constructs.Construct, id string) traefikio.IngressRouteSpec {
//	return traefikio.IngressRouteSpec{
//		Routes: traefikio.IngressRouteSpecRoutes{
//			Kind:        "",
//			Match:       nil,
//			Middlewares: nil,
//			Priority:    nil,
//			Services:    nil,
//			Syntax:      nil,
//		},
//		EntryPoints: nil,
//		Tls:         nil,
//	}
//}

func defaultIngressRouteAnnotations(name string) *map[string]*string {
	return nil
}

func defaultIngressRouteLabels(name string) *map[string]string {
	return nil
}

func defaultIngressRouteMetadata(svcName string) *cdk8s.ApiObjectMetadata {
	namespace := TestbedNamespace

	return &cdk8s.ApiObjectMetadata{
		Namespace:   jsii.String(namespace),
		Annotations: defaultIngressRouteAnnotations(svcName),
		Labels:      defaultMiddlewareLabels(svcName),
	}
}

func ingressRouteSpec() *traefikio.IngressRouteSpec {
	return &traefikio.IngressRouteSpec{
		Routes:      nil,
		EntryPoints: nil,
		Tls:         nil,
	}
}

func newIngressRouteProps(svcName string) *traefikio.IngressRouteProps {
	metadata := defaultIngressRouteMetadata(svcName)
	return &traefikio.IngressRouteProps{
		Metadata: metadata,
		Spec:     nil,
	}
}
