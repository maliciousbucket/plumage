package types

import (
	"errors"
	"fmt"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2"
	compose "github.com/compose-spec/compose-go/v2/types"
	"strconv"
)

//TODO: App protocol
//TODO: Host IP

type Port struct {
	Name          *string
	ContainerPort *float64
	PublishedPort *float64
	Protocol      *Protocol
}

func (p *Port) PortProtocol() *string {
	if p.Protocol == nil {
		return nil
	}

	protocol := string(*p.Protocol)
	return &protocol
}

func (p *Port) PortPublished() *float64 {
	return p.PublishedPort
}

func (p *Port) PortContainer() *float64 {
	return p.ContainerPort
}

func (p *Port) PortName() *string {
	return p.Name
}

func ParseComposeContainerPorts(containerPorts []compose.ServicePortConfig) ([]*Port, error) {
	var ports []*Port
	var portErrors error

	for _, containerPort := range containerPorts {
		port, err := ParseComposeContainerPort(containerPort)
		if err != nil {
			portErrors = errors.Join(portErrors, err)
			continue
		}

		ports = append(ports, port)

	}
	if len(ports) == 0 {
		portErrors = errors.New("no valid ports found")
	}

	return ports, portErrors
}

func ParseComposeContainerPort(containerPort compose.ServicePortConfig) (*Port, error) {
	var port Port

	if containerPort.Name != "" {
		port.Name = &containerPort.Name
	}

	if containerPort.Target == 0 {
		return nil, fmt.Errorf("invalid target port: %d", containerPort.Target)
	}
	portNumber := float64(containerPort.Target)
	port.ContainerPort = &portNumber

	if containerPort.Published == "" {
		return nil, fmt.Errorf("published container port is required")
	}

	portNumber, err := strconv.ParseFloat(containerPort.Published, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing published port number: %s. Error: %w", containerPort.Published, err)

	}
	port.PublishedPort = &portNumber

	if containerPort.Protocol != "" {
		protocol, err := ValidProtocol(containerPort.Protocol)
		if err != nil {
			return nil, fmt.Errorf("error validating protocol: %s. Error: %w", containerPort.Protocol, err)

		}
		port.Protocol = protocol
	} else {
		protocol := ProtocolTCP
		port.Protocol = &protocol
	}
	return &port, nil
}

func (p *Port) K8sContainerPort() (*cdk8splus30.ContainerPort, error) {

	port := p.ContainerPort
	if port == nil {
		return nil, errors.New("no container port found")
	}

	hostPort := p.PublishedPort
	if hostPort == nil {
		return nil, errors.New("no host port found")
	}

	protocol := string(*p.Protocol)

	//
	//containerPort := k8s.ContainerPort{
	//	ContainerPort: port,
	//	HostPort:      hostPort,
	//	Name:          p.Name,
	//	Protocol:      &protocol,
	//}
	containerPort := cdk8splus30.ContainerPort{
		Number:   port,
		HostIp:   nil,
		HostPort: hostPort,
		Name:     p.Name,
		Protocol: cdk8splus30.Protocol(protocol),
	}

	return &containerPort, nil

}
