package autoscalingk8sio


// ContainerResourcePolicy controls how autoscaler computes the recommended resources for a specific container.
type VerticalPodAutoscalerSpecResourcePolicyContainerPolicies struct {
	// Name of the container or DefaultContainerResourcePolicy, in which case the policy is used by the containers that don't have their own policy specified.
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// Specifies the type of recommendations that will be computed (and possibly applied) by VPA.
	//
	// If not specified, the default of [ResourceCPU, ResourceMemory] will be used.
	ControlledResources *[]*string `field:"optional" json:"controlledResources" yaml:"controlledResources"`
	// Specifies which resource values should be controlled.
	//
	// The default is "RequestsAndLimits".
	ControlledValues VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesControlledValues `field:"optional" json:"controlledValues" yaml:"controlledValues"`
	// Specifies the maximum amount of resources that will be recommended for the container.
	//
	// The default is no maximum.
	MaxAllowed *map[string]VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMaxAllowed `field:"optional" json:"maxAllowed" yaml:"maxAllowed"`
	// Specifies the minimal amount of resources that will be recommended for the container.
	//
	// The default is no minimum.
	MinAllowed *map[string]VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMinAllowed `field:"optional" json:"minAllowed" yaml:"minAllowed"`
	// Whether autoscaler is enabled for the container.
	//
	// The default is "Auto".
	Mode VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMode `field:"optional" json:"mode" yaml:"mode"`
}

