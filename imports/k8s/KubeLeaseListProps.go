package k8s


// LeaseList is a list of Lease objects.
type KubeLeaseListProps struct {
	// items is a list of schema objects.
	Items *[]*KubeLeaseProps `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

