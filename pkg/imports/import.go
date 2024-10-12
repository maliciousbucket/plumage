package imports

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

const (
	argoInstallUrl = "https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml"

	verticalUrl = "https://raw.githubusercontent.com/kubernetes/autoscaler/master/vertical-pod-autoscaler/deploy/vpa-crd.yaml"
)

func AddTraefikRbac(scope constructs.Construct, id string) constructs.Construct {
	traefik := cdk8s.NewInclude(scope, jsii.String(id), &cdk8s.IncludeProps{
		Url: jsii.String(""),
	})
	return traefik
}

func AddTraefikResources(scope constructs.Construct, id string) constructs.Construct {
	resources := cdk8s.NewInclude(scope, jsii.String(id), &cdk8s.IncludeProps{
		Url: jsii.String(""),
	})
	return resources
}

func InstallArgoCRDs(scope constructs.Construct, id string) constructs.Construct {
	argo := cdk8s.NewInclude(scope, jsii.String(id), &cdk8s.IncludeProps{
		Url: jsii.String(argoInstallUrl),
	})
	return argo
}

func InstallVerticalAutoscalerRbac(scope constructs.Construct, id string) constructs.Construct {
	vertical := cdk8s.NewInclude(scope, jsii.String(id), &cdk8s.IncludeProps{
		Url: jsii.String(verticalUrl),
	})
	return vertical
}
