package k8s


// IngressClassList is a collection of IngressClasses.
type KubeIngressClassListProps struct {
	// items is the list of IngressClasses.
	Items *[]*KubeIngressClassProps `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

