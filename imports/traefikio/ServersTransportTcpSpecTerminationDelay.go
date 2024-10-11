package traefikio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/maliciousbucket/plumage/imports/traefikio/jsii"
)

// TerminationDelay defines the delay to wait before fully terminating the connection, after one connected peer has closed its writing capability.
type ServersTransportTcpSpecTerminationDelay interface {
	Value() interface{}
}

// The jsii proxy struct for ServersTransportTcpSpecTerminationDelay
type jsiiProxy_ServersTransportTcpSpecTerminationDelay struct {
	_ byte // padding
}

func (j *jsiiProxy_ServersTransportTcpSpecTerminationDelay) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func ServersTransportTcpSpecTerminationDelay_FromNumber(value *float64) ServersTransportTcpSpecTerminationDelay {
	_init_.Initialize()

	if err := validateServersTransportTcpSpecTerminationDelay_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ServersTransportTcpSpecTerminationDelay

	_jsii_.StaticInvoke(
		"traefikio.ServersTransportTcpSpecTerminationDelay",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ServersTransportTcpSpecTerminationDelay_FromString(value *string) ServersTransportTcpSpecTerminationDelay {
	_init_.Initialize()

	if err := validateServersTransportTcpSpecTerminationDelay_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ServersTransportTcpSpecTerminationDelay

	_jsii_.StaticInvoke(
		"traefikio.ServersTransportTcpSpecTerminationDelay",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}
