package k8s


// FlowSchema defines the schema of a group of flows.
//
// Note that a flow is made up of a set of inbound API requests with similar attributes and is identified by a pair of strings: the name of the FlowSchema and a "flow distinguisher".
type KubeFlowSchemaProps struct {
	// `metadata` is the standard object's metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
	// `spec` is the specification of the desired behavior of a FlowSchema.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *FlowSchemaSpec `field:"optional" json:"spec" yaml:"spec"`
}

