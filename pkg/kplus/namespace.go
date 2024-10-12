package kplus

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	kplus "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2"
)

func ProjectNamespace(scope constructs.Construct, ns string) kplus.Namespace {
	chart := cdk8s.NewChart(scope, jsii.String("project-namespace"), &cdk8s.ChartProps{})
	namespace := kplus.NewNamespace(chart, jsii.String(ns), &kplus.NamespaceProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Annotations:     nil,
			Finalizers:      nil,
			Labels:          nil,
			Name:            jsii.String(ns),
			Namespace:       nil,
			OwnerReferences: nil,
		},
	})
	return namespace
}

func addNamespaceDependency(namespace kplus.Namespace, chart cdk8s.Chart) {
	chart.AddDependency(namespace)
}
