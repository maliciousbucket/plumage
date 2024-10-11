// chaosgalah-monitoringio
package chaosgalahmonitoringio

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"chaosgalah-monitoringio.TestRunJob",
		reflect.TypeOf((*TestRunJob)(nil)).Elem(),
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
			j := jsiiProxy_TestRunJob{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"chaosgalah-monitoringio.TestRunJobProps",
		reflect.TypeOf((*TestRunJobProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"chaosgalah-monitoringio.TestRunJobSpec",
		reflect.TypeOf((*TestRunJobSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"chaosgalah-monitoringio.TestRunJobSpecEnv",
		reflect.TypeOf((*TestRunJobSpecEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"chaosgalah-monitoringio.TestRunJobSpecEnvValueFrom",
		reflect.TypeOf((*TestRunJobSpecEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"chaosgalah-monitoringio.TestRunJobSpecEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*TestRunJobSpecEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"chaosgalah-monitoringio.TestRunJobSpecEnvValueFromFieldRef",
		reflect.TypeOf((*TestRunJobSpecEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"chaosgalah-monitoringio.TestRunJobSpecEnvValueFromResourceFieldRef",
		reflect.TypeOf((*TestRunJobSpecEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"chaosgalah-monitoringio.TestRunJobSpecEnvValueFromResourceFieldRefDivisor",
		reflect.TypeOf((*TestRunJobSpecEnvValueFromResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_TestRunJobSpecEnvValueFromResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"chaosgalah-monitoringio.TestRunJobSpecEnvValueFromSecretKeyRef",
		reflect.TypeOf((*TestRunJobSpecEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"chaosgalah-monitoringio.TestRunJobSpecJobTemplate",
		reflect.TypeOf((*TestRunJobSpecJobTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"chaosgalah-monitoringio.TestRunJobSpecSchedule",
		reflect.TypeOf((*TestRunJobSpecSchedule)(nil)).Elem(),
	)
}
