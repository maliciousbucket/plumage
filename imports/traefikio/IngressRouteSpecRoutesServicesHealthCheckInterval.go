package traefikio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/maliciousbucket/plumage/imports/traefikio/jsii"
)

// Interval defines the frequency of the health check calls.
//
// Default: 30s.
type IngressRouteSpecRoutesServicesHealthCheckInterval interface {
	Value() interface{}
}

// The jsii proxy struct for IngressRouteSpecRoutesServicesHealthCheckInterval
type jsiiProxy_IngressRouteSpecRoutesServicesHealthCheckInterval struct {
	_ byte // padding
}

func (j *jsiiProxy_IngressRouteSpecRoutesServicesHealthCheckInterval) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func IngressRouteSpecRoutesServicesHealthCheckInterval_FromNumber(value *float64) IngressRouteSpecRoutesServicesHealthCheckInterval {
	_init_.Initialize()

	if err := validateIngressRouteSpecRoutesServicesHealthCheckInterval_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns IngressRouteSpecRoutesServicesHealthCheckInterval

	_jsii_.StaticInvoke(
		"traefikio.IngressRouteSpecRoutesServicesHealthCheckInterval",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func IngressRouteSpecRoutesServicesHealthCheckInterval_FromString(value *string) IngressRouteSpecRoutesServicesHealthCheckInterval {
	_init_.Initialize()

	if err := validateIngressRouteSpecRoutesServicesHealthCheckInterval_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns IngressRouteSpecRoutesServicesHealthCheckInterval

	_jsii_.StaticInvoke(
		"traefikio.IngressRouteSpecRoutesServicesHealthCheckInterval",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}
