package autoscalingk8sio

// Specification of the checkpoint.
//
// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
type VerticalPodAutoscalerCheckpointSpec struct {
	// Name of the checkpointed container.
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// Name of the VPA object that stored VerticalPodAutoscalerCheckpoint object.
	VpaObjectName *string `field:"optional" json:"vpaObjectName" yaml:"vpaObjectName"`
}
