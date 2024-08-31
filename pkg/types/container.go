package types

import (
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2"
	compose "github.com/compose-spec/compose-go/v2/types"
)

const (
	DefaultImagePullPolicy = "IfNotPresent"
)

//Env, volumes
//TODO: Command probe to live check / startup
//TODO: Host IP - Deployment, Service, Container, Ingress Route

type ContainerSpec struct {
	Name           string
	Image          string              `json:"image"`
	Ports          []*Port             `json:"ports"`
	Resources      *ContainerResources `json:"resources"`
	HealthCheck    *CommandProbe       `json:"commandProbes"`
	ReadinessProbe *HttpProbe          `json:"http_probes"`
	Volumes        []*ContainerVolume  `json:"volumes"`
	Commands       []*string           `json:"commands"`
	Args           []*string           `json:"args"`
	WorkingDir     string              `json:"workingDir"`
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
	var args []*string

	if len(config.Command) > 0 {
		for _, commandSpec := range config.Command {
			args = append(args, &commandSpec)
		}
	}

	var commands []*string
	if len(config.Entrypoint) > 0 {
		for _, entrypointSpec := range config.Entrypoint {
			commands = append(commands, &entrypointSpec)
		}
	}

	var cmdProbe *CommandProbe

	if config.HealthCheck != nil {
		probe, err := ParseHealthCheck(config.HealthCheck)
		if err != nil {
			return nil, err
		}
		cmdProbe = probe
	}
	var workingDir string
	if config.WorkingDir != "" {
		workingDir = config.WorkingDir
	}

	container := &ContainerSpec{
		Name:        config.Name,
		Image:       config.Image,
		Ports:       ports,
		Commands:    commands,
		Args:        args,
		HealthCheck: cmdProbe,
		WorkingDir:  workingDir,
	}
	return container, nil
}

func (c *ContainerSpec) ToContainerProps() (*cdk8splus30.ContainerProps, error) {
	var ports []*cdk8splus30.ContainerPort
	for _, portSpec := range c.Ports {
		port, err := portSpec.K8sContainerPort()
		if err != nil {
			return nil, err
		}
		ports = append(ports, port)
	}

	var startup cdk8splus30.Probe
	if c.HealthCheck != nil {
		probe, err := c.HealthCheck.Probe()
		if err != nil {
			return nil, err
		}
		startup = probe
	}

	return &cdk8splus30.ContainerProps{
		Args:            &c.Args,
		Command:         &c.Commands,
		EnvFrom:         nil,
		EnvVariables:    nil,
		ImagePullPolicy: DefaultImagePullPolicy,
		Lifecycle:       nil,
		Liveness:        nil,
		Name:            &c.Name,
		Port:            nil,
		PortNumber:      nil,
		Ports:           &ports,
		Readiness:       c.ReadinessProbe,
		Resources:       nil,
		RestartPolicy:   "",
		SecurityContext: nil,
		Startup:         &startup,
		VolumeMounts:    nil,
		WorkingDir:      nil,
		Image:           &c.Image,
	}, nil
}
