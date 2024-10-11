// autoscalingk8sio
package autoscalingk8sio

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"autoscalingk8sio.VerticalPodAutoscaler",
		reflect.TypeOf((*VerticalPodAutoscaler)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_VerticalPodAutoscaler{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"autoscalingk8sio.VerticalPodAutoscalerCheckpoint",
		reflect.TypeOf((*VerticalPodAutoscalerCheckpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_VerticalPodAutoscalerCheckpoint{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"autoscalingk8sio.VerticalPodAutoscalerCheckpointProps",
		reflect.TypeOf((*VerticalPodAutoscalerCheckpointProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"autoscalingk8sio.VerticalPodAutoscalerCheckpointSpec",
		reflect.TypeOf((*VerticalPodAutoscalerCheckpointSpec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"autoscalingk8sio.VerticalPodAutoscalerCheckpointV1Beta2",
		reflect.TypeOf((*VerticalPodAutoscalerCheckpointV1Beta2)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_VerticalPodAutoscalerCheckpointV1Beta2{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"autoscalingk8sio.VerticalPodAutoscalerCheckpointV1Beta2Props",
		reflect.TypeOf((*VerticalPodAutoscalerCheckpointV1Beta2Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"autoscalingk8sio.VerticalPodAutoscalerCheckpointV1Beta2Spec",
		reflect.TypeOf((*VerticalPodAutoscalerCheckpointV1Beta2Spec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"autoscalingk8sio.VerticalPodAutoscalerProps",
		reflect.TypeOf((*VerticalPodAutoscalerProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"autoscalingk8sio.VerticalPodAutoscalerSpec",
		reflect.TypeOf((*VerticalPodAutoscalerSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"autoscalingk8sio.VerticalPodAutoscalerSpecRecommenders",
		reflect.TypeOf((*VerticalPodAutoscalerSpecRecommenders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"autoscalingk8sio.VerticalPodAutoscalerSpecResourcePolicy",
		reflect.TypeOf((*VerticalPodAutoscalerSpecResourcePolicy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"autoscalingk8sio.VerticalPodAutoscalerSpecResourcePolicyContainerPolicies",
		reflect.TypeOf((*VerticalPodAutoscalerSpecResourcePolicyContainerPolicies)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"autoscalingk8sio.VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesControlledValues",
		reflect.TypeOf((*VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesControlledValues)(nil)).Elem(),
		map[string]interface{}{
			"REQUESTS_AND_LIMITS": VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesControlledValues_REQUESTS_AND_LIMITS,
			"REQUESTS_ONLY":       VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesControlledValues_REQUESTS_ONLY,
		},
	)
	_jsii_.RegisterClass(
		"autoscalingk8sio.VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMaxAllowed",
		reflect.TypeOf((*VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMaxAllowed)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMaxAllowed{}
		},
	)
	_jsii_.RegisterClass(
		"autoscalingk8sio.VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMinAllowed",
		reflect.TypeOf((*VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMinAllowed)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMinAllowed{}
		},
	)
	_jsii_.RegisterEnum(
		"autoscalingk8sio.VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMode",
		reflect.TypeOf((*VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMode)(nil)).Elem(),
		map[string]interface{}{
			"AUTO": VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMode_AUTO,
			"OFF":  VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMode_OFF,
		},
	)
	_jsii_.RegisterStruct(
		"autoscalingk8sio.VerticalPodAutoscalerSpecTargetRef",
		reflect.TypeOf((*VerticalPodAutoscalerSpecTargetRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"autoscalingk8sio.VerticalPodAutoscalerSpecUpdatePolicy",
		reflect.TypeOf((*VerticalPodAutoscalerSpecUpdatePolicy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"autoscalingk8sio.VerticalPodAutoscalerSpecUpdatePolicyEvictionRequirements",
		reflect.TypeOf((*VerticalPodAutoscalerSpecUpdatePolicyEvictionRequirements)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"autoscalingk8sio.VerticalPodAutoscalerSpecUpdatePolicyEvictionRequirementsChangeRequirement",
		reflect.TypeOf((*VerticalPodAutoscalerSpecUpdatePolicyEvictionRequirementsChangeRequirement)(nil)).Elem(),
		map[string]interface{}{
			"TARGET_HIGHER_THAN_REQUESTS": VerticalPodAutoscalerSpecUpdatePolicyEvictionRequirementsChangeRequirement_TARGET_HIGHER_THAN_REQUESTS,
			"TARGET_LOWER_THAN_REQUESTS":  VerticalPodAutoscalerSpecUpdatePolicyEvictionRequirementsChangeRequirement_TARGET_LOWER_THAN_REQUESTS,
		},
	)
	_jsii_.RegisterEnum(
		"autoscalingk8sio.VerticalPodAutoscalerSpecUpdatePolicyUpdateMode",
		reflect.TypeOf((*VerticalPodAutoscalerSpecUpdatePolicyUpdateMode)(nil)).Elem(),
		map[string]interface{}{
			"OFF":      VerticalPodAutoscalerSpecUpdatePolicyUpdateMode_OFF,
			"INITIAL":  VerticalPodAutoscalerSpecUpdatePolicyUpdateMode_INITIAL,
			"RECREATE": VerticalPodAutoscalerSpecUpdatePolicyUpdateMode_RECREATE,
			"AUTO":     VerticalPodAutoscalerSpecUpdatePolicyUpdateMode_AUTO,
		},
	)
	_jsii_.RegisterClass(
		"autoscalingk8sio.VerticalPodAutoscalerV1Beta2",
		reflect.TypeOf((*VerticalPodAutoscalerV1Beta2)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_VerticalPodAutoscalerV1Beta2{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"autoscalingk8sio.VerticalPodAutoscalerV1Beta2Props",
		reflect.TypeOf((*VerticalPodAutoscalerV1Beta2Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"autoscalingk8sio.VerticalPodAutoscalerV1Beta2Spec",
		reflect.TypeOf((*VerticalPodAutoscalerV1Beta2Spec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"autoscalingk8sio.VerticalPodAutoscalerV1Beta2SpecResourcePolicy",
		reflect.TypeOf((*VerticalPodAutoscalerV1Beta2SpecResourcePolicy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"autoscalingk8sio.VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPolicies",
		reflect.TypeOf((*VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPolicies)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"autoscalingk8sio.VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMaxAllowed",
		reflect.TypeOf((*VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMaxAllowed)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMaxAllowed{}
		},
	)
	_jsii_.RegisterClass(
		"autoscalingk8sio.VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMinAllowed",
		reflect.TypeOf((*VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMinAllowed)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMinAllowed{}
		},
	)
	_jsii_.RegisterEnum(
		"autoscalingk8sio.VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMode",
		reflect.TypeOf((*VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMode)(nil)).Elem(),
		map[string]interface{}{
			"AUTO": VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMode_AUTO,
			"OFF":  VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMode_OFF,
		},
	)
	_jsii_.RegisterStruct(
		"autoscalingk8sio.VerticalPodAutoscalerV1Beta2SpecTargetRef",
		reflect.TypeOf((*VerticalPodAutoscalerV1Beta2SpecTargetRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"autoscalingk8sio.VerticalPodAutoscalerV1Beta2SpecUpdatePolicy",
		reflect.TypeOf((*VerticalPodAutoscalerV1Beta2SpecUpdatePolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"autoscalingk8sio.VerticalPodAutoscalerV1Beta2SpecUpdatePolicyUpdateMode",
		reflect.TypeOf((*VerticalPodAutoscalerV1Beta2SpecUpdatePolicyUpdateMode)(nil)).Elem(),
		map[string]interface{}{
			"OFF":      VerticalPodAutoscalerV1Beta2SpecUpdatePolicyUpdateMode_OFF,
			"INITIAL":  VerticalPodAutoscalerV1Beta2SpecUpdatePolicyUpdateMode_INITIAL,
			"RECREATE": VerticalPodAutoscalerV1Beta2SpecUpdatePolicyUpdateMode_RECREATE,
			"AUTO":     VerticalPodAutoscalerV1Beta2SpecUpdatePolicyUpdateMode_AUTO,
		},
	)
}
