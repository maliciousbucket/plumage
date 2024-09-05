package types

import (
	"errors"
	"strings"
)

const (
	ProtocolTCP = Protocol("TCP")
	ProtocolUDP = Protocol("UDP")

	ServiceTypeLoadBalancer = ServiceType("LoadBalancer")
	ServiceTypeClusterIP    = ServiceType("ClusterIP")
)

type ServiceType string

type Protocol string

func ValidProtocol(input string) (*Protocol, error) {
	protocol := strings.ToUpper(input)
	switch Protocol(protocol) {
	case ProtocolTCP:
		p := ProtocolTCP
		return &p, nil
	case ProtocolUDP:
		p := ProtocolUDP
		return &p, nil
	default:
		return nil, errors.New("invalid Protocol")
	}

}

type WebService struct {
	Name      string
	Container *ContainerSpec
	//ResiliencePolicy *resilience.Policy
	//CircuitBreakerPolicy *template.CircuitBreakerConfig
	//RetryPolicy          *template.RetryConfig
	//RateLimitPolicy      *template.RateLimitConfig
	Labels map[string]*string
	Env    map[string]string
}
