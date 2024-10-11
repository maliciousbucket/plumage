package autoscalingk8sio

// Specification of the behavior of the autoscaler.
//
// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
type VerticalPodAutoscalerV1Beta2Spec struct {
	// TargetRef points to the controller managing the set of pods for the autoscaler to control - e.g. Deployment, StatefulSet. VerticalPodAutoscaler can be targeted at controller implementing scale subresource (the pod set is retrieved from the controller's ScaleStatus) or some well known controllers (e.g. for DaemonSet the pod set is read from the controller's spec). If VerticalPodAutoscaler cannot use specified target it will report ConfigUnsupported condition. Note that VerticalPodAutoscaler does not require full implementation of scale subresource - it will not use it to modify the replica count. The only thing retrieved is a label selector matching pods grouped by the target resource.
	TargetRef *VerticalPodAutoscalerV1Beta2SpecTargetRef `field:"required" json:"targetRef" yaml:"targetRef"`
	// Controls how the autoscaler computes recommended resources.
	//
	// The resource policy may be used to set constraints on the recommendations for individual containers. If not specified, the autoscaler computes recommended resources for all containers in the pod, without additional constraints.
	ResourcePolicy *VerticalPodAutoscalerV1Beta2SpecResourcePolicy `field:"optional" json:"resourcePolicy" yaml:"resourcePolicy"`
	// Describes the rules on how changes are applied to the pods.
	//
	// If not specified, all fields in the `PodUpdatePolicy` are set to their default values.
	UpdatePolicy *VerticalPodAutoscalerV1Beta2SpecUpdatePolicy `field:"optional" json:"updatePolicy" yaml:"updatePolicy"`
}
