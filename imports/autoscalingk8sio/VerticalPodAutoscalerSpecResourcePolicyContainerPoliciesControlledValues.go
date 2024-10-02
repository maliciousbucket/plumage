package autoscalingk8sio


// Specifies which resource values should be controlled.
//
// The default is "RequestsAndLimits".
type VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesControlledValues string

const (
	// RequestsAndLimits.
	VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesControlledValues_REQUESTS_AND_LIMITS VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesControlledValues = "REQUESTS_AND_LIMITS"
	// RequestsOnly.
	VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesControlledValues_REQUESTS_ONLY VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesControlledValues = "REQUESTS_ONLY"
)

