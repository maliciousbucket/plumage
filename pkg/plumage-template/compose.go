package plumage_template

import (
	"fmt"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2"
	composeTypes "github.com/compose-spec/compose-go/v2/types"
	"github.com/maliciousbucket/plumage/pkg/types"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func MatchServices(p *PlumageTemplate, compose map[string]*composeTypes.ServiceConfig) []*ServiceConfig {
	var services []*ServiceConfig
	for _, service := range p.Services {
		var composeService *composeTypes.ServiceConfig
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

func getComposeServices(p *composeTypes.Project) map[string]*composeTypes.ServiceConfig {
	services := map[string]*composeTypes.ServiceConfig{}
	for _, service := range p.Services {
		services[service.Name] = &composeTypes.ServiceConfig{}
	}
	return services
}

func GetCommandProbe(service composeTypes.ServiceConfig) (*types.CommandProbe, error) {
	if service.HealthCheck == nil {
		return nil, nil
	}
	probe, err := types.ParseHealthCheck(service.HealthCheck)
	if err != nil {
		return nil, err
	}
	return probe, nil
}

func GetVolumes(service composeTypes.ServiceConfig, projectDir string) ([]*types.ContainerVolume, error) {
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

type ContainerVolume interface {
	TargetVolume() string
}

func parseContainerVolume(projectDir string, svcName string, c composeTypes.ServiceVolumeConfig) (ContainerVolume, error) {
	if c.Target == "" {
		return nil, fmt.Errorf("invalid volume target")
	}
	var cv ContainerVolume
	if c.Source == "" {
		emptyVolume, err := parseEmptyContainerVolume(c)
		if err != nil {
			return nil, err
		}
		cv = emptyVolume
		return emptyVolume, nil
	}

	filePath := c.Source
	file, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}
	switch mode := file.Mode(); {
	case mode.IsDir():
		dirVolume, err := parseContainerVolumeFromDir(projectDir, c)
		if err != nil {
			return nil, err
		}
		cv = dirVolume
	case mode.IsRegular():
		fileVolume, err := parseContainerVolumeFromFile(projectDir, c)
		if err != nil {
			return nil, err
		}
		cv = fileVolume
	default:
		return nil, fmt.Errorf("invalid volume mode")
	}
	return cv, nil

}

func parseContainerVolumeFromDir(projectDir string, c composeTypes.ServiceVolumeConfig) (*DirVolume, error) {
	if c.Source == "" {
		return nil, fmt.Errorf("invalid volume source")
	}
	if c.Target == "" {
		return nil, fmt.Errorf("invalid volume target")
	}
	sourcePath := path.Join(projectDir, c.Source)
	dirPath, err := filepath.Abs(sourcePath)
	if err != nil {
		return nil, err
	}
	info, err := os.Stat(dirPath)
	if err != nil {
		return nil, err
	}

	volume := &DirVolume{
		Source: dirPath,
		Target: c.Target,
		Name:   info.Name(),
	}

	return volume, nil

}

func parseContainerVolumeFromFile(projectDir string, c composeTypes.ServiceVolumeConfig) (*FileVolume, error) {
	if c.Source == "" {
		return nil, fmt.Errorf("invalid volume source")
	}
	if c.Target == "" {
		return nil, fmt.Errorf("invalid volume target")
	}

	sourcePath := path.Join(projectDir, c.Source)
	filePath, err := filepath.Abs(sourcePath)
	if err != nil {
		return nil, err
	}

	info, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}

	volume := &FileVolume{
		Source: filePath,
		Target: c.Target,
		Name:   info.Name(),
	}
	return volume, nil
}

func parseEmptyContainerVolume(c composeTypes.ServiceVolumeConfig) (*EmptyVolume, error) {
	if c.Target == "" {
		return nil, fmt.Errorf("invalid volume target")
	}
	return &EmptyVolume{
		Target: c.Target,
	}, nil
}
