package autoscalingk8sio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// VerticalPodAutoscaler is the configuration for a vertical pod autoscaler, which automatically manages pod resources based on historical and real time resource utilization.
type VerticalPodAutoscalerProps struct {
	// Specification of the behavior of the autoscaler.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
	Spec     *VerticalPodAutoscalerSpec `field:"required" json:"spec" yaml:"spec"`
	Metadata *cdk8s.ApiObjectMetadata   `field:"optional" json:"metadata" yaml:"metadata"`
}
