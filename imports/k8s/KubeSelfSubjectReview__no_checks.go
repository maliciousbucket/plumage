//go:build no_runtime_type_checking

package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeSelfSubjectReview_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateKubeSelfSubjectReview_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeSelfSubjectReview_ManifestParameters(props *KubeSelfSubjectReviewProps) error {
	return nil
}

func validateKubeSelfSubjectReview_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeSelfSubjectReviewParameters(scope constructs.Construct, id *string, props *KubeSelfSubjectReviewProps) error {
	return nil
}

