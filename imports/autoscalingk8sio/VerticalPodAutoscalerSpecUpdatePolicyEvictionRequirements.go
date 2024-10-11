package autoscalingk8sio

// EvictionRequirement defines a single condition which needs to be true in order to evict a Pod.
type VerticalPodAutoscalerSpecUpdatePolicyEvictionRequirements struct {
	// EvictionChangeRequirement refers to the relationship between the new target recommendation for a Pod and its current requests, what kind of change is necessary for the Pod to be evicted.
	ChangeRequirement VerticalPodAutoscalerSpecUpdatePolicyEvictionRequirementsChangeRequirement `field:"required" json:"changeRequirement" yaml:"changeRequirement"`
	// Resources is a list of one or more resources that the condition applies to.
	//
	// If more than one resource is given, the EvictionRequirement is fulfilled if at least one resource meets `changeRequirement`.
	Resources *[]*string `field:"required" json:"resources" yaml:"resources"`
}
