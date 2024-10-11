package autoscalingk8sio

// Controls how the autoscaler computes recommended resources.
//
// The resource policy may be used to set constraints on the recommendations for individual containers. If not specified, the autoscaler computes recommended resources for all containers in the pod, without additional constraints.
type VerticalPodAutoscalerV1Beta2SpecResourcePolicy struct {
	// Per-container resource policies.
	ContainerPolicies *[]*VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPolicies `field:"optional" json:"containerPolicies" yaml:"containerPolicies"`
}
