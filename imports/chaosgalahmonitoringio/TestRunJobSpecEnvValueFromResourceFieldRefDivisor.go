package chaosgalahmonitoringio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/maliciousbucket/plumage/imports/chaosgalahmonitoringio/jsii"
)

// Specifies the output format of the exposed resources, defaults to "1".
type TestRunJobSpecEnvValueFromResourceFieldRefDivisor interface {
	Value() interface{}
}

// The jsii proxy struct for TestRunJobSpecEnvValueFromResourceFieldRefDivisor
type jsiiProxy_TestRunJobSpecEnvValueFromResourceFieldRefDivisor struct {
	_ byte // padding
}

func (j *jsiiProxy_TestRunJobSpecEnvValueFromResourceFieldRefDivisor) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func TestRunJobSpecEnvValueFromResourceFieldRefDivisor_FromNumber(value *float64) TestRunJobSpecEnvValueFromResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validateTestRunJobSpecEnvValueFromResourceFieldRefDivisor_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns TestRunJobSpecEnvValueFromResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"chaosgalah-monitoringio.TestRunJobSpecEnvValueFromResourceFieldRefDivisor",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func TestRunJobSpecEnvValueFromResourceFieldRefDivisor_FromString(value *string) TestRunJobSpecEnvValueFromResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validateTestRunJobSpecEnvValueFromResourceFieldRefDivisor_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns TestRunJobSpecEnvValueFromResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"chaosgalah-monitoringio.TestRunJobSpecEnvValueFromResourceFieldRefDivisor",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}
