package k8s


// ResourceClaimList is a collection of claims.
type KubeResourceClaimListV1Alpha2Props struct {
	// Items is the list of resource claims.
	Items *[]*KubeResourceClaimV1Alpha2Props `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

