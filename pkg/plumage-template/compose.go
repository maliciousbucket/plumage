package plumage_template

import (
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2"
	types2 "github.com/compose-spec/compose-go/v2/types"
	"github.com/maliciousbucket/plumage/pkg/types"
	"strings"
)

func MatchServices(p *PlumageTemplate, compose map[string]*types2.ServiceConfig) []*ServiceConfig {
	var services []*ServiceConfig
	for _, service := range p.Services {
		var composeService *types2.ServiceConfig
		if comp, ok := compose[service.Name]; ok {
			composeService = comp
		}
		services = append(services, &ServiceConfig{
			service,
			composeService,
		})
	}
	return services
}

func getComposeServices(p *types2.Project) map[string]*types2.ServiceConfig {
	services := map[string]*types2.ServiceConfig{}
	for _, service := range p.Services {
		services[service.Name] = &types2.ServiceConfig{}
	}
	return services
}

func GetCommandProbe(service types2.ServiceConfig) (*types.CommandProbe, error) {
	if service.HealthCheck == nil {
		return nil, nil
	}
	probe, err := types.ParseHealthCheck(service.HealthCheck)
	if err != nil {
		return nil, err
	}
	return probe, nil
}

func GetVolumes(service types2.ServiceConfig, projectDir string) ([]*types.ContainerVolume, error) {
	if len(service.Volumes) == 0 {
		return nil, nil
	}
	var volumes []*types.ContainerVolume
	for _, volume := range service.Volumes {
		containerVolume, err := types.ParseContainerVolume(projectDir, service.Name, volume)
		if err != nil {
			return nil, err
		}
		volumes = append(volumes, containerVolume)
	}
	return volumes, nil
}

func GetProtocol(prt string) cdk8splus30.Protocol {
	protocol := strings.ToLower(prt)
	switch protocol {
	case "tcp":
		return cdk8splus30.Protocol_TCP
	case "udp":
		return cdk8splus30.Protocol_UDP
	default:
		return cdk8splus30.Protocol_TCP
	}
}
