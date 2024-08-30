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

	container := &ContainerSpec{
		Name:     config.Name,
		Image:    config.Image,
		Ports:    ports,
		Commands: commands,
	}
	return container, nil
}
