package autoscalingk8sio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/maliciousbucket/plumage/imports/autoscalingk8sio/jsii"
)

type VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMaxAllowed interface {
	Value() interface{}
}

// The jsii proxy struct for VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMaxAllowed
type jsiiProxy_VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMaxAllowed struct {
	_ byte // padding
}

func (j *jsiiProxy_VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMaxAllowed) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMaxAllowed_FromNumber(value *float64) VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMaxAllowed {
	_init_.Initialize()

	if err := validateVerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMaxAllowed_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMaxAllowed

	_jsii_.StaticInvoke(
		"autoscalingk8sio.VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMaxAllowed",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMaxAllowed_FromString(value *string) VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMaxAllowed {
	_init_.Initialize()

	if err := validateVerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMaxAllowed_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMaxAllowed

	_jsii_.StaticInvoke(
		"autoscalingk8sio.VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMaxAllowed",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

