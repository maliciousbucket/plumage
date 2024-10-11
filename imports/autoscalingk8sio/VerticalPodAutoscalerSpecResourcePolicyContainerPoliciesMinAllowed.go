package autoscalingk8sio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/maliciousbucket/plumage/imports/autoscalingk8sio/jsii"
)

type VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMinAllowed interface {
	Value() interface{}
}

// The jsii proxy struct for VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMinAllowed
type jsiiProxy_VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMinAllowed struct {
	_ byte // padding
}

func (j *jsiiProxy_VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMinAllowed) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMinAllowed_FromNumber(value *float64) VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMinAllowed {
	_init_.Initialize()

	if err := validateVerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMinAllowed_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMinAllowed

	_jsii_.StaticInvoke(
		"autoscalingk8sio.VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMinAllowed",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMinAllowed_FromString(value *string) VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMinAllowed {
	_init_.Initialize()

	if err := validateVerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMinAllowed_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMinAllowed

	_jsii_.StaticInvoke(
		"autoscalingk8sio.VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMinAllowed",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}
