package k8s


// PriorityLevelConfigurationList is a list of PriorityLevelConfiguration objects.
type KubePriorityLevelConfigurationListProps struct {
	// `items` is a list of request-priorities.
	Items *[]*KubePriorityLevelConfigurationProps `field:"required" json:"items" yaml:"items"`
	// `metadata` is the standard object's metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

