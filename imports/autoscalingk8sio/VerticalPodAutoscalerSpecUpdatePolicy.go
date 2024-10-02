package autoscalingk8sio


// Describes the rules on how changes are applied to the pods.
//
// If not specified, all fields in the `PodUpdatePolicy` are set to their default values.
type VerticalPodAutoscalerSpecUpdatePolicy struct {
	// EvictionRequirements is a list of EvictionRequirements that need to evaluate to true in order for a Pod to be evicted.
	//
	// If more than one EvictionRequirement is specified, all of them need to be fulfilled to allow eviction.
	EvictionRequirements *[]*VerticalPodAutoscalerSpecUpdatePolicyEvictionRequirements `field:"optional" json:"evictionRequirements" yaml:"evictionRequirements"`
	// Minimal number of replicas which need to be alive for Updater to attempt pod eviction (pending other checks like PDB).
	//
	// Only positive values are allowed. Overrides global '--min-replicas' flag.
	MinReplicas *float64 `field:"optional" json:"minReplicas" yaml:"minReplicas"`
	// Controls when autoscaler applies changes to the pod resources.
	//
	// The default is 'Auto'.
	UpdateMode VerticalPodAutoscalerSpecUpdatePolicyUpdateMode `field:"optional" json:"updateMode" yaml:"updateMode"`
}

