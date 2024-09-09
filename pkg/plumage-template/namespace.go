package plumage_template

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	kplus "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2"
)

func NewNamespace(scope constructs.Construct, id string, name string) constructs.Construct {
	ns := kplus.NewNamespace(scope, jsii.String(id), &kplus.NamespaceProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Name: jsii.String(name),
		},
	})
	return ns
}
