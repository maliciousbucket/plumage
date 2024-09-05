package k8s

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/maliciousbucket/plumage/imports/k8s"
	"github.com/maliciousbucket/plumage/pkg/types"
	"time"

	plumagetemplate "github.com/maliciousbucket/plumage/pkg/plumage-template"
)

type ContainerProps struct {
	Name         string
	Image        string
	Commands     []string
	Args         []string
	StartupProbe *types.CommandProbe
	HealthCheck  *plumagetemplate.HttpProbe
	Ports        []*plumagetemplate.ServicePort
	Resources    *plumagetemplate.ServiceResources
	Monitoring   *plumagetemplate.MonitoringConfig
	Env          map[string]string
	WorkingDir   string
}

func NewContainer(props *ContainerProps) k8s.Container {
	args := StringSliceToK8s(props.Args)
	commands := StringSliceToK8s(props.Commands)
	env := StringMapToEnv(props.Env)

	metricsPort := 0
	if props.Monitoring != nil {
		metricsPort = props.Monitoring.ScrapePort
	}

	ports := ContainerPorts(props.Ports, metricsPort)

	if props.HealthCheck != nil {
	}

	container := k8s.Container{
		Name:            jsii.String(props.Name),
		Args:            args,
		Command:         commands,
		Env:             env,
		EnvFrom:         nil,
		Image:           jsii.String(props.Image),
		ImagePullPolicy: jsii.String("IfNotPresent"),
		Lifecycle:       nil,
		LivenessProbe:   nil,
		Ports:           &ports,
		ReadinessProbe:  nil,
		Resources:       nil,
		SecurityContext: nil,
		StartupProbe:    nil,
		VolumeDevices:   nil,
		VolumeMounts:    nil,
		WorkingDir:      nil,
	}
	if props.Resources != nil {
		resources := ContainerResources(props.Resources)
		container.Resources = resources
	}

	if props.HealthCheck != nil {
		healthCheck := HttpProbeToLiveCheck(props.HealthCheck)
		container.LivenessProbe = healthCheck
	}

	if props.StartupProbe != nil {
		probe := StartupProbeToK8s(props.StartupProbe)
		container.StartupProbe = probe
	}

	if props.WorkingDir != "" {
		container.WorkingDir = jsii.String(props.WorkingDir)
	}
	return container
}

func ContainerResources(resources *plumagetemplate.ServiceResources) *k8s.ResourceRequirements {
	if resources == nil {
		return nil
	}

	var cpuLimit k8s.Quantity
	var memoryLimit k8s.Quantity
	var cpuRequest k8s.Quantity
	var memoryRequest k8s.Quantity
	if resources.CpuLimitMillis != 0 {
		cpuLimit = k8s.Quantity_FromNumber(jsii.Number(resources.CpuLimitMillis))
	}

	if resources.MemoryRequestMb != 0 {
		memoryRequest = k8s.Quantity_FromNumber(jsii.Number(resources.MemoryRequestMb))
	}

	if resources.CpuRequestMillis != 0 {
		cpuRequest = k8s.Quantity_FromNumber(jsii.Number(resources.CpuRequestMillis))
	}

	if resources.MemoryRequestMb != 0 {
		memoryRequest = k8s.Quantity_FromNumber(jsii.Number(resources.MemoryRequestMb))
	}

	limits := map[string]k8s.Quantity{
		"memory": memoryLimit,
		"cpu":    cpuLimit,
	}

	requests := map[string]k8s.Quantity{
		"cpu":    cpuRequest,
		"memory": memoryRequest,
	}

	return &k8s.ResourceRequirements{
		Limits:   &limits,
		Requests: &requests,
	}
}

func ContainerPorts(ports []*plumagetemplate.ServicePort, metricsPort int) []*k8s.ContainerPort {
	var containerPorts []*k8s.ContainerPort
	i := 0
	for _, port := range ports {
		name := PortName(port.GRPC, i)
		if metricsPort != 0 {
			if port.ContainerPort == metricsPort {
				name = "http-metrics"
			}
		}
		protocol := GetProtocol(port.Protocol)
		containerPorts = append(containerPorts, &k8s.ContainerPort{
			ContainerPort: jsii.Number(port.ContainerPort),
			HostIp:        nil,
			HostPort:      jsii.Number(port.HostPort),
			Name:          jsii.String(name),
			Protocol:      jsii.String(protocol),
		})
		i++
	}
	return containerPorts
}

func HttpProbeToLiveCheck(check *plumagetemplate.HttpProbe) *k8s.Probe {
	if check == nil {
		return nil
	}
	var failureThreshold *float64
	var initialDelaySeconds *float64
	var periodSeconds *float64
	var timeout *float64

	if check.Threshold != 0 {
		failureThreshold = jsii.Number(check.Threshold)
	}

	if check.InitialDelaySeconds != 0 {
		initialDelaySeconds = jsii.Number(check.InitialDelaySeconds)
	}

	if check.TimeoutSeconds != 0 {
		timeout = jsii.Number(check.TimeoutSeconds)
	}

	if check.PeriodSeconds != 0 {
		periodSeconds = jsii.Number(check.PeriodSeconds)
	}

	return &k8s.Probe{
		FailureThreshold: failureThreshold,

		HttpGet: &k8s.HttpGetAction{
			Port:        k8s.IntOrString_FromNumber(jsii.Number(check.Port)),
			Host:        nil,
			HttpHeaders: nil,
			Path:        jsii.String(check.Path),
			Scheme:      nil,
		},
		InitialDelaySeconds:           initialDelaySeconds,
		PeriodSeconds:                 periodSeconds,
		SuccessThreshold:              nil,
		TerminationGracePeriodSeconds: nil,
		TimeoutSeconds:                timeout,
	}
}

func StartupProbeToK8s(commandProbe *types.CommandProbe) *k8s.Probe {
	if commandProbe == nil {
		return nil
	}

	exec := commandProbe.Commands
	delay, _ := time.ParseDuration(*commandProbe.Delay)
	delaySeconds := delay.Seconds()
	timeout, _ := time.ParseDuration(*commandProbe.Timeout)
	timeoutSeconds := timeout.Seconds()
	period, _ := time.ParseDuration(*commandProbe.Interval)
	periodSeconds := period.Seconds()

	probe := k8s.Probe{
		Exec:                &k8s.ExecAction{Command: &exec},
		FailureThreshold:    commandProbe.Retries,
		InitialDelaySeconds: &delaySeconds,
		PeriodSeconds:       &periodSeconds,
		TimeoutSeconds:      &timeoutSeconds,
	}
	return &probe
}
