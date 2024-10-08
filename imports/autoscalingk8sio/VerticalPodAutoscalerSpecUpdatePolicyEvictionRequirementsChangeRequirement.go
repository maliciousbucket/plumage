package autoscalingk8sio


// EvictionChangeRequirement refers to the relationship between the new target recommendation for a Pod and its current requests, what kind of change is necessary for the Pod to be evicted.
type VerticalPodAutoscalerSpecUpdatePolicyEvictionRequirementsChangeRequirement string

const (
	// TargetHigherThanRequests.
	VerticalPodAutoscalerSpecUpdatePolicyEvictionRequirementsChangeRequirement_TARGET_HIGHER_THAN_REQUESTS VerticalPodAutoscalerSpecUpdatePolicyEvictionRequirementsChangeRequirement = "TARGET_HIGHER_THAN_REQUESTS"
	// TargetLowerThanRequests.
	VerticalPodAutoscalerSpecUpdatePolicyEvictionRequirementsChangeRequirement_TARGET_LOWER_THAN_REQUESTS VerticalPodAutoscalerSpecUpdatePolicyEvictionRequirementsChangeRequirement = "TARGET_LOWER_THAN_REQUESTS"
)

