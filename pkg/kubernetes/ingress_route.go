package kubernetes

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/maliciousbucket/plumage/imports/traefikio"
)

const (
	defaultRouteType = "Rule"
)

func newIngressRoute(scope constructs.Construct, id string) traefikio.IngressRoute {
	return nil
}

func newProps(scope constructs.Construct, id string) traefikio.IngressRouteProps {
	return nil
}

func newSpec(scope constructs.Construct, id string) traefikio.IngressRouteSpec {
	return traefikio.IngressRouteSpec{
		Routes: traefikio.IngressRouteSpecRoutes{
			Kind:        "",
			Match:       nil,
			Middlewares: nil,
			Priority:    nil,
			Services:    nil,
			Syntax:      nil,
		},
		EntryPoints: nil,
		Tls:         nil,
	}
}
