package disruptor

import (
	"encoding/json"
	"fmt"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	kplus "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2"
	"github.com/maliciousbucket/plumage/imports/k8s"
	"log"
)

const (
	serviceTypePath = "/spec/type"
	nodePortPath    = "/spec/nodePort"
)

type Template struct {
	DeploymentName string
	Port           int
}

func AddDisruptorContainer(scope constructs.Construct, target cdk8s.ApiObject, params *Template) constructs.Construct {
	name := fmt.Sprintf("%s-%s", *target.Name(), "disruptor")
	container := k8s.Container{
		Name:            jsii.String(name),
		Image:           jsii.String("ghcr.io/grafana/xk6-disruptor-agent:latest"),
		ImagePullPolicy: jsii.String("IfNotPresent"),
		Ports: &[]*k8s.ContainerPort{
			{
				ContainerPort: jsii.Number(params.Port),
				HostPort:      jsii.Number(params.Port),
			},
		},
	}
	jsonData, err := json.Marshal(container)
	if err != nil {
		log.Fatalf("Error converting container to JSON: %v", err)
	}
	path := "/spec/template/spec/containers/1/"
	patch := cdk8s.JsonPatch_Add(jsii.String(path), jsonData)
	target.AddJsonPatch(patch)
	return target
}

func AddDisruptor(object cdk8s.ApiObject) constructs.Construct {
	patch := cdk8s.JsonPatch_Add(jsii.String(serviceTypePath), kplus.ServiceType_LOAD_BALANCER)

	object.AddJsonPatch(patch)
	return object
}

func AddDisruptorWithNodePort(object cdk8s.ApiObject, port int) constructs.Construct {
	typePatch := cdk8s.JsonPatch_Replace(jsii.String(serviceTypePath), kplus.ServiceType_NODE_PORT)
	portPatch := cdk8s.JsonPatch_Add(jsii.String(nodePortPath), jsii.Number(port))
	object.AddJsonPatch(typePatch, portPatch)
	return object
}
