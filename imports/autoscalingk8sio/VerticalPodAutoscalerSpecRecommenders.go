package autoscalingk8sio


// VerticalPodAutoscalerRecommenderSelector points to a specific Vertical Pod Autoscaler recommender.
//
// In the future it might pass parameters to the recommender.
type VerticalPodAutoscalerSpecRecommenders struct {
	// Name of the recommender responsible for generating recommendation for this object.
	Name *string `field:"required" json:"name" yaml:"name"`
}

