package kubernetes

import (
	kplus "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2"
	"github.com/maliciousbucket/plumage/pkg/types"
)

func AddContainerResources(container kplus.Container, resources *types.ContainerResources) error {
	return nil
}

func AddContainerPorts(container kplus.Container, ports []*types.Port) {
	for _, port := range ports {
		protocol := port.Cdk8sProtocol()
		container.AddPort(&kplus.ContainerPort{
			Number:   port.ContainerPort,
			HostIp:   nil,
			HostPort: port.PublishedPort,
			Name:     port.Name,
			Protocol: protocol,
		})
	}
}
