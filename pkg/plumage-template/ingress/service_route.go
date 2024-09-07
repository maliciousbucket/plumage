package ingress

import (
	"fmt"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/maliciousbucket/plumage/imports/traefikio"
)

func NewServiceIngressRoute(scope constructs.Construct, id string, props *RouteProps, middlewareRefs []*string) traefikio.IngressRoute {
	if props.Config == nil {
		return nil
	}

	metadata := ingressRouteMetaData(props.Namespace, props.Name)

	ingresRouteProps := &traefikio.IngressRouteProps{
		Metadata: metadata,
		Spec: &traefikio.IngressRouteSpec{
			Routes: &[]*traefikio.IngressRouteSpecRoutes{},
		},
	}
	var ingressRoute []*traefikio.IngressRouteSpecRoutes

	if len(props.Config.Paths) > 0 {
		for _, path := range props.Config.Paths {
			newRouteSpec := traefikio.IngressRouteSpecRoutes{}
			withMiddlewares(props.Namespace, middlewareRefs)(&newRouteSpec)
			withRule(props.Config.Host, path.Path)(&newRouteSpec)
			fmt.Println(path.Path)
			withService(props.Namespace, props.Name, *path, props.Config.EnableLoadBalancer)(&newRouteSpec)
			fmt.Printf("New Route Spec: %v", newRouteSpec)
			ingressRoute = append(ingressRoute, &newRouteSpec)
		}
	}
	ingresRouteProps.Spec.Routes = &ingressRoute

	return traefikio.NewIngressRoute(scope, jsii.String(id), ingresRouteProps)
}

type ingressRouteRoutesFunc func(r *traefikio.IngressRouteSpecRoutes)

func withMiddlewares(ns string, middlewareRefs []*string) ingressRouteRoutesFunc {
	return func(r *traefikio.IngressRouteSpecRoutes) {
		var references []*traefikio.IngressRouteSpecRoutesMiddlewares
		for _, m := range middlewareRefs {
			references = append(references, &traefikio.IngressRouteSpecRoutesMiddlewares{
				Name:      m,
				Namespace: &ns,
			})
		}
		r.Middlewares = &references
	}
}

type RouteServiceProps struct {
	Name               string
	Namespace          string
	Path               ServicePaths
	EnableLoadBalancer bool
}

func withRule(host string, prefix string) ingressRouteRoutesFunc {
	return func(r *traefikio.IngressRouteSpecRoutes) {
		fmt.Printf("\nHost: %s Path: %s\n", host, prefix)
		rule := fmt.Sprintf("PathPrefix('%s')", prefix)
		if host != "" {
			rule = fmt.Sprintf("Host('%s') && %s", host, rule)
		}

		r.Kind = RouteTypeRule
		r.Match = jsii.String(rule)
		fmt.Printf("\nRule: %s\n", rule)
	}
}

func withService(ns string, appLabel string, path ServicePaths, lb bool) ingressRouteRoutesFunc {
	return func(r *traefikio.IngressRouteSpecRoutes) {
		svc := &traefikio.IngressRouteSpecRoutesServices{
			Namespace: jsii.String(ns),
			Name:      jsii.String(appLabel),
			Kind:      traefikio.IngressRouteSpecRoutesServicesKind_SERVICE,
			Port:      traefikio.IngressRouteSpecRoutesServicesPort_FromNumber(jsii.Number(path.Port)),
			NativeLb:  jsii.Bool(lb),
		}
		if r.Services != nil {
			services := append(*r.Services, svc)
			r.Services = &services
		} else {
			r.Services = &[]*traefikio.IngressRouteSpecRoutesServices{svc}
		}

	}
}
