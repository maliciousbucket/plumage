package autoscalingk8sio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/maliciousbucket/plumage/imports/autoscalingk8sio/jsii"
)

type VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMaxAllowed interface {
	Value() interface{}
}

// The jsii proxy struct for VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMaxAllowed
type jsiiProxy_VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMaxAllowed struct {
	_ byte // padding
}

func (j *jsiiProxy_VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMaxAllowed) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMaxAllowed_FromNumber(value *float64) VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMaxAllowed {
	_init_.Initialize()

	if err := validateVerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMaxAllowed_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMaxAllowed

	_jsii_.StaticInvoke(
		"autoscalingk8sio.VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMaxAllowed",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMaxAllowed_FromString(value *string) VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMaxAllowed {
	_init_.Initialize()

	if err := validateVerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMaxAllowed_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMaxAllowed

	_jsii_.StaticInvoke(
		"autoscalingk8sio.VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMaxAllowed",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

