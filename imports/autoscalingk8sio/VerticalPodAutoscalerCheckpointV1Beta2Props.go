package autoscalingk8sio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// VerticalPodAutoscalerCheckpoint is the checkpoint of the internal state of VPA that is used for recovery after recommender's restart.
type VerticalPodAutoscalerCheckpointV1Beta2Props struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// Specification of the checkpoint.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
	Spec *VerticalPodAutoscalerCheckpointV1Beta2Spec `field:"optional" json:"spec" yaml:"spec"`
}
