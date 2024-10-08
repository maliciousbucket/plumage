package k8s


// Ingress is a collection of rules that allow inbound connections to reach the endpoints defined by a backend.
//
// An Ingress can be configured to give services externally-reachable urls, load balance traffic, terminate SSL, offer name based virtual hosting etc.
type KubeIngressProps struct {
	// Standard object's metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
	// spec is the desired state of the Ingress.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *IngressSpec `field:"optional" json:"spec" yaml:"spec"`
}

