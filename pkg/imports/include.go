package imports

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

func NewInclude(scope constructs.Construct, id string, path string) constructs.Construct {
	include := cdk8s.NewInclude(scope, jsii.String(id), &cdk8s.IncludeProps{Url: jsii.String(path)})
	return include
}
