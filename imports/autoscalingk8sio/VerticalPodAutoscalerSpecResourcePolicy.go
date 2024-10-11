package autoscalingk8sio

// Controls how the autoscaler computes recommended resources.
//
// The resource policy may be used to set constraints on the recommendations for individual containers. If any individual containers need to be excluded from getting the VPA recommendations, then it must be disabled explicitly by setting mode to "Off" under containerPolicies. If not specified, the autoscaler computes recommended resources for all containers in the pod, without additional constraints.
type VerticalPodAutoscalerSpecResourcePolicy struct {
	// Per-container resource policies.
	ContainerPolicies *[]*VerticalPodAutoscalerSpecResourcePolicyContainerPolicies `field:"optional" json:"containerPolicies" yaml:"containerPolicies"`
}
