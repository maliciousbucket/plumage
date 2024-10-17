package traefikio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// ServersTransportTCP is the CRD implementation of a TCPServersTransport.
//
// If no tcpServersTransport is specified, a default one named default@internal will be used.
// The default@internal tcpServersTransport can be configured in the static configuration.
// More info: https://doc.traefik.io/traefik/v3.1/routing/services/#serverstransport_3
type ServersTransportTcpProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// ServersTransportTCPSpec defines the desired state of a ServersTransportTCP.
	Spec *ServersTransportTcpSpec `field:"required" json:"spec" yaml:"spec"`
}

