package types

import compose "github.com/compose-spec/compose-go/v2/types"

//Env, volumes

type ContainerSpec struct {
	Name          string
	Image         string              `json:"image"`
	Ports         []*Port             `json:"ports"`
	Resources     *ContainerResources `json:"resources"`
	CommandProbes []*CommandProbe     `json:"command_probes"`
	HttpProbes    []*HttpProbe        `json:"http_probes"`
	Volumes       []*ContainerVolume  `json:"volumes"`
	Commands      []*string           `json:"commands"`
}

func ParseServiceConfig(config compose.ServiceConfig) (*ContainerSpec, error) {
	var ports []*Port

	for _, portSpec := range config.Ports {
		port, err := ParseComposeContainerPort(portSpec)
		if err != nil {
			return nil, err
		}
		ports = append(ports, port)
	}
	var commands []*string
	if len(config.Command) > 0 {
		for _, commandSpec := range config.Command {
			commands = append(commands, &commandSpec)
		}
	}

	var cmdProbes []*CommandProbe

	if config.HealthCheck != nil {
		probe, err := ParseHealthCheck(config.HealthCheck)
		if err != nil {
			return nil, err
		}
		cmdProbes = append(cmdProbes, probe)
	}

	container := &ContainerSpec{
		Name:          config.Name,
		Image:         config.Image,
		Ports:         ports,
		Commands:      commands,
		CommandProbes: cmdProbes,
	}
	return container, nil
}
