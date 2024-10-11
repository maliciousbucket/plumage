package traefikio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/maliciousbucket/plumage/imports/traefikio/jsii"
)

// Interval defines the frequency of the health check calls.
//
// Default: 30s.
type MiddlewareSpecErrorsServiceHealthCheckInterval interface {
	Value() interface{}
}

// The jsii proxy struct for MiddlewareSpecErrorsServiceHealthCheckInterval
type jsiiProxy_MiddlewareSpecErrorsServiceHealthCheckInterval struct {
	_ byte // padding
}

func (j *jsiiProxy_MiddlewareSpecErrorsServiceHealthCheckInterval) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func MiddlewareSpecErrorsServiceHealthCheckInterval_FromNumber(value *float64) MiddlewareSpecErrorsServiceHealthCheckInterval {
	_init_.Initialize()

	if err := validateMiddlewareSpecErrorsServiceHealthCheckInterval_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns MiddlewareSpecErrorsServiceHealthCheckInterval

	_jsii_.StaticInvoke(
		"traefikio.MiddlewareSpecErrorsServiceHealthCheckInterval",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func MiddlewareSpecErrorsServiceHealthCheckInterval_FromString(value *string) MiddlewareSpecErrorsServiceHealthCheckInterval {
	_init_.Initialize()

	if err := validateMiddlewareSpecErrorsServiceHealthCheckInterval_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns MiddlewareSpecErrorsServiceHealthCheckInterval

	_jsii_.StaticInvoke(
		"traefikio.MiddlewareSpecErrorsServiceHealthCheckInterval",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}
