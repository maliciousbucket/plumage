package traefikio

// ServersTransportTCPSpec defines the desired state of a ServersTransportTCP.
type ServersTransportTcpSpec struct {
	// DialKeepAlive is the interval between keep-alive probes for an active network connection.
	//
	// If zero, keep-alive probes are sent with a default value (currently 15 seconds), if supported by the protocol and operating system. Network protocols or operating systems that do not support keep-alives ignore this field. If negative, keep-alive probes are disabled.
	DialKeepAlive ServersTransportTcpSpecDialKeepAlive `field:"optional" json:"dialKeepAlive" yaml:"dialKeepAlive"`
	// DialTimeout is the amount of time to wait until a connection to a backend server can be established.
	DialTimeout ServersTransportTcpSpecDialTimeout `field:"optional" json:"dialTimeout" yaml:"dialTimeout"`
	// TerminationDelay defines the delay to wait before fully terminating the connection, after one connected peer has closed its writing capability.
	TerminationDelay ServersTransportTcpSpecTerminationDelay `field:"optional" json:"terminationDelay" yaml:"terminationDelay"`
	// TLS defines the TLS configuration.
	Tls *ServersTransportTcpSpecTls `field:"optional" json:"tls" yaml:"tls"`
}
