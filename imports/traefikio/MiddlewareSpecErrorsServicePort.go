package traefikio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/maliciousbucket/plumage/imports/traefikio/jsii"
)

// Port defines the port of a Kubernetes Service.
//
// This can be a reference to a named port.
type MiddlewareSpecErrorsServicePort interface {
	Value() interface{}
}

// The jsii proxy struct for MiddlewareSpecErrorsServicePort
type jsiiProxy_MiddlewareSpecErrorsServicePort struct {
	_ byte // padding
}

func (j *jsiiProxy_MiddlewareSpecErrorsServicePort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func MiddlewareSpecErrorsServicePort_FromNumber(value *float64) MiddlewareSpecErrorsServicePort {
	_init_.Initialize()

	if err := validateMiddlewareSpecErrorsServicePort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns MiddlewareSpecErrorsServicePort

	_jsii_.StaticInvoke(
		"traefikio.MiddlewareSpecErrorsServicePort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func MiddlewareSpecErrorsServicePort_FromString(value *string) MiddlewareSpecErrorsServicePort {
	_init_.Initialize()

	if err := validateMiddlewareSpecErrorsServicePort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns MiddlewareSpecErrorsServicePort

	_jsii_.StaticInvoke(
		"traefikio.MiddlewareSpecErrorsServicePort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}
