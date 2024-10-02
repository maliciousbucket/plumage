package autoscalingk8sio


// ContainerResourcePolicy controls how autoscaler computes the recommended resources for a specific container.
type VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPolicies struct {
	// Name of the container or DefaultContainerResourcePolicy, in which case the policy is used by the containers that don't have their own policy specified.
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// Specifies the maximum amount of resources that will be recommended for the container.
	//
	// The default is no maximum.
	MaxAllowed *map[string]VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMaxAllowed `field:"optional" json:"maxAllowed" yaml:"maxAllowed"`
	// Specifies the minimal amount of resources that will be recommended for the container.
	//
	// The default is no minimum.
	MinAllowed *map[string]VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMinAllowed `field:"optional" json:"minAllowed" yaml:"minAllowed"`
	// Whether autoscaler is enabled for the container.
	//
	// The default is "Auto".
	Mode VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMode `field:"optional" json:"mode" yaml:"mode"`
}

