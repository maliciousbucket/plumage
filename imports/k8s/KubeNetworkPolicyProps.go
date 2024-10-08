package k8s


// NetworkPolicy describes what network traffic is allowed for a set of Pods.
type KubeNetworkPolicyProps struct {
	// Standard object's metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
	// spec represents the specification of the desired behavior for this NetworkPolicy.
	Spec *NetworkPolicySpec `field:"optional" json:"spec" yaml:"spec"`
}

