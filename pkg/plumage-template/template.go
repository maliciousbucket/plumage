package plumage_template

import types2 "github.com/compose-spec/compose-go/v2/types"

type PlumageTemplate struct {
	Services            []*Service
	InitContainers      []*InitContainer
	UnmatchedContainers []string
	UnmatchedServices   []string
	Compose             []types2.ServiceConfig
}

func (p *PlumageTemplate) GetService(serviceName string) *Service {
	for _, service := range p.Services {
		if service.Name == serviceName {
			return service
		}
	}
	return nil
}

func (p *PlumageTemplate) GetInitContainer(name string) *InitContainer {
	for _, initContainer := range p.InitContainers {
		if initContainer.Name == name {
			return initContainer
		}
	}
	return nil
}

func (p *PlumageTemplate) ListServiceNames() []string {
	var serviceNames []string
	for _, service := range p.Services {
		serviceNames = append(serviceNames, service.Name)
	}
	return serviceNames
}

func (p *PlumageTemplate) findUnmatched(project types2.Project) {
	containers := project.ServiceNames()
	services := p.ListServiceNames()
	p.UnmatchedContainers = append(p.UnmatchedContainers, containers...)
	p.UnmatchedServices = append(p.UnmatchedServices, services...)

}

func (p *PlumageTemplate) ListUnmatched(project types2.Project) map[string][]string {
	p.findUnmatched(project)
	unmatched := make(map[string][]string)
	unmatched["Services"] = p.UnmatchedServices
	unmatched["Containers"] = p.UnmatchedContainers
	return unmatched
}

func (p *PlumageTemplate) AddComposeServices(project types2.Project) {
	services := project.AllServices()
	for _, service := range services {
		p.Compose = append(p.Compose, service)
	}
}

type ServiceConfig struct {
	Service *Service
	Compose *types2.ServiceConfig
}

type Service struct {
	Name            string            `yaml:"name"`
	Image           string            `yaml:"image"`
	Host            string            `yaml:"host"`
	Ports           []*ServicePort    `yaml:"ports"`
	Commands        []string          `yaml:"commands"`
	Args            []string          `yaml:"args"`
	HttpHealthCheck string            `yaml:"httpHealthCheck"`
	Env             map[string]string `yaml:"env"`

	Resources          *ServiceResources `yaml:"resources"`
	Scaling            *ScalingConfig    `yaml:"scaling"`
	Monitoring         *MonitoringConfig `yaml:"monitoring"`
	InitContainerNames []string          `yaml:"initContainers"`
	Middleware         []string          `yaml:"middleware"`
	Paths              []*ServicePaths   `yaml:"paths"`
	LoadBalancer       bool              `yaml:"loadBalancer"`
	SynthOptions       *SynthOptions     `yaml:"synthOptions"`
}

type SynthOptions struct {
	Deployment     bool `yaml:"deployment,omitempty"`
	IngressRoute   bool `yaml:"ingressRoute,omitempty"`
	AutoScaling    bool `yaml:"autoScaling,omitempty"`
	Service        bool `yaml:"service,omitempty"`
	ComposeImage   bool `yaml:"useComposeImage,omitempty"`
	ComposeVolumes bool `yaml:"composeVolumes"`
	InitContainers bool `yaml:"initContainers,omitempty"`
	ComposePorts   bool `yaml:"composePorts,omitempty"`
}

type IngressOptions struct {
	Paths              []*ServicePaths `yaml:"paths"`
	EnableLoadBalancer bool            `yaml:"enableLoadBalancer,omitempty"`
}

type InitContainer struct {
	Name      string            `yaml:"name"`
	Image     string            `yaml:"image"`
	Commands  []string          `yaml:"commands"`
	Resources *ServiceResources `yaml:"resources"`
}

type ServicePort struct {
	ContainerPort int    `yaml:"containerPort"`
	HostPort      int    `yaml:"hostPort,omitempty"`
	Protocol      string `yaml:"protocol,omitempty"`
	Name          string `yaml:"name,omitempty"`
	Grpc          bool   `yaml:"grpc,omitempty"`
}

type ServicePaths struct {
	Path string
	Port int
}

type MonitoringConfig struct {
	MetricsPath   string `yaml:"metricsPath"`
	ScrapePort    int    `yaml:"scrapePort"`
	MonitoringEnv map[string]string
}

type ScalingConfig struct {
	TargetCpuAmount     int              `yaml:"targetCpuAmount"`
	TargetCpuPercent    int              `yaml:"targetCpuPercent"`
	TargetMemoryAmount  int              `yaml:"targetMemoryAmount"`
	TargetMemoryPercent int              `yaml:"targetMemoryPercent"`
	MinReplicas         int              `yaml:"minReplicas"`
	MaxReplicas         int              `yaml:"maxReplicas"`
	TargetReplicas      int              `yaml:"targetReplicas"`
	Resources           *ResourceScaling `yaml:"resources"`
}

type ResourceScaling struct {
	TargetMemoryUsagePercent float64 `yaml:"targetMemoryUsagePercent"`
	TargetMemoryUsageAmount  float64 `yaml:"targetMemoryUsageAmount"`
	TargetCpuUsagePercent    float64 `yaml:"targetCpuUsagePercent"`
	TargetCpuUsageAmount     float64 `yaml:"targetCpuUsageAmount"`
}

type ServiceResources struct {
	CpuRequestMillis int `yaml:"cpuRequestMillis"`
	CpuLimitMillis   int `yaml:"cpuLimitMillis"`
	MemoryRequestMb  int `yaml:"memoryRequestMb"`
	MemoryLimitMb    int `yaml:"memoryLimitMb"`
}

type HttpProbe struct {
	Path                string `yaml:"path"`
	Port                int    `yaml:"port"`
	Threshold           int    `yaml:"threshold"`
	InitialDelaySeconds int    `yaml:"initialDelaySeconds"`
	TimeoutSeconds      int    `yaml:"timeoutSeconds"`
	PeriodSeconds       int    `yaml:"periodSeconds"`
}

type EmptyVolume struct {
	Target string `yaml:"target"`
}

type FileVolume struct {
}
