package plumage_template

import (
	"fmt"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2"
	"github.com/maliciousbucket/plumage/pkg/types"
)

type WebServiceProps struct {
	Name         string
	Image        string
	Ports        []*ServicePort
	Args         []string
	Commands     []string
	Volumes      []*types.ContainerVolume
	StartupProbe *types.CommandProbe
	HealthCheck  *HttpProbe
	Env          map[string]string
	Monitoring   *MonitoringConfig
	Scaling      *ScalingConfig
	Resources    *ServiceResources
}

func NewWevService(scope constructs.Construct, id string, props *WebServiceProps) constructs.Construct {
	ct := constructs.NewConstruct(scope, jsii.String(id))

	deploymentName := fmt.Sprintf("%s-deployment", props.Name)
	deployment := cdk8splus30.NewDeployment(ct, jsii.String(deploymentName), nil)
	containerProps := NewContainer(props)

	container := deployment.AddContainer(&containerProps)
	var metricsPort int
	if props.Monitoring != nil {
		metricsPort = props.Monitoring.ScrapePort
	}
	AddContainerPorts(container, props.Ports, metricsPort)
	AddContainerEnv(container, props.Env)

	servicePorts := GetServicePorts(props.Ports, metricsPort)
	deployment.ExposeViaService(&cdk8splus30.DeploymentExposeViaServiceOptions{

		Name:        jsii.String(props.Name),
		Ports:       &servicePorts,
		ServiceType: cdk8splus30.ServiceType_CLUSTER_IP,
	})
	if props.Scaling != nil {
		autoScalerId := fmt.Sprintf("%s-autoscaler", props.Name)
		autoSclaerProps := AutoScalerProps(props.Scaling, deployment)
		cdk8splus30.NewHorizontalPodAutoscaler(ct, jsii.String(autoScalerId), autoSclaerProps)

	}

	return ct

}

func NewContainer(props *WebServiceProps) cdk8splus30.ContainerProps {
	args := StringSliceToK8s(props.Args)
	commands := StringSliceToK8s(props.Commands)

	var healthProbe cdk8splus30.Probe
	if props.HealthCheck != nil {
		healthProbe = HealthCheckProbe(props.HealthCheck)
	}

	var resources *cdk8splus30.ContainerResources
	if props.Resources != nil {
		resources = AddContainerResources(props.Resources)
	}

	return cdk8splus30.ContainerProps{
		Args:            args,
		Command:         commands,
		ImagePullPolicy: cdk8splus30.ImagePullPolicy_IF_NOT_PRESENT,
		Liveness:        healthProbe,
		Name:            jsii.String(props.Name),
		Readiness:       props.StartupProbe,
		Resources:       resources,
		RestartPolicy:   cdk8splus30.ContainerRestartPolicy_ALWAYS,
		Startup:         nil,
		VolumeMounts:    nil,
		Image:           jsii.String(props.Image),
	}
}

func ServicePorts(ports []*ServicePort) []*cdk8splus30.ServicePort {

	return nil
}

func ContainerPorts(ports []*ServicePort) []cdk8splus30.ContainerPort {
	return nil
}

func HealthCheckProbe(check *HttpProbe) cdk8splus30.Probe {
	if check == nil {
		return nil
	}
	options := HealthCheckProbeOpts(check)

	return cdk8splus30.Probe_FromHttpGet(jsii.String(check.Path), &options)
}

func HealthCheckProbeOpts(check *HttpProbe) cdk8splus30.HttpGetProbeOptions {

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

	return cdk8splus30.HttpGetProbeOptions{
		FailureThreshold:    failureThreshold,
		InitialDelaySeconds: cdk8s.Duration_Seconds(initialDelaySeconds),
		PeriodSeconds:       cdk8s.Duration_Seconds(periodSeconds),
		TimeoutSeconds:      cdk8s.Duration_Seconds(timeout),
		Port:                jsii.Number(check.Port),
	}
}

func StringSliceToK8s(sl []string) *[]*string {
	var slice []*string
	for _, st := range sl {
		slice = append(slice, &st)
	}
	return &slice
}

func StringMapToK8s(mp map[string]string) *map[string]*string {
	var k8sMap map[string]*string
	for key, value := range mp {
		k8sMap[key] = &value
	}
	return &k8sMap
}

func StringMapToEnv(mp map[string]string) *map[string]cdk8splus30.EnvValue {
	var envMap map[string]cdk8splus30.EnvValue
	for key, value := range mp {
		variable := cdk8splus30.EnvValue_FromValue(jsii.String(value))
		envMap[key] = variable
	}
	return &envMap
}

func AddContainerEnv(container cdk8splus30.Container, env map[string]string) cdk8splus30.Container {
	for key, value := range env {
		container.Env().AddVariable(jsii.String(key), cdk8splus30.EnvValue_FromValue(jsii.String(value)))
	}
	return container
}

func AddContainerPorts(container cdk8splus30.Container, ports []*ServicePort, metricsPort int) cdk8splus30.Container {
	if len(ports) == 0 {
		return container
	}
	var containerPorts []*cdk8splus30.ContainerPort
	i := 0
	for _, port := range ports {
		protocol := GetProtocol(port.Protocol)
		name := portName(port.GRPC, i)
		if metricsPort != 0 {
			if port.ContainerPort == metricsPort {
				name = "http-metrics"
			}
		}
		containerPorts = append(containerPorts, &cdk8splus30.ContainerPort{
			Number:   jsii.Number(port.ContainerPort),
			HostPort: jsii.Number(port.HostPort),
			Name:     jsii.String(name),
			Protocol: protocol,
		})
		i++
	}
	for _, containerPort := range containerPorts {
		container.AddPort(containerPort)
	}
	return container

}

func portName(rpc bool, count int) string {
	var protocol string
	switch rpc {
	case false:
		protocol = "http"
	case true:
		protocol = "grpc"
	}
	return fmt.Sprintf("%s-%d", protocol, count)
}

func AddContainerResources(resources *ServiceResources) *cdk8splus30.ContainerResources {

	if resources == nil {
		return nil
	}

	var cpu cdk8splus30.CpuResources
	if resources.CpuRequestMillis != 0 {
		cpu.Request = cdk8splus30.Cpu_Millis(jsii.Number(resources.CpuRequestMillis))
	}
	if resources.CpuLimitMillis != 0 {
		cpu.Limit = cdk8splus30.Cpu_Millis(jsii.Number(resources.CpuLimitMillis))
	}

	var memory cdk8splus30.MemoryResources
	if resources.MemoryRequestMb != 0 {
		memory.Request = cdk8s.Size_Mebibytes(jsii.Number(resources.MemoryRequestMb))
	}
	if resources.MemoryLimitMb != 0 {
		memory.Limit = cdk8s.Size_Mebibytes(jsii.Number(resources.MemoryLimitMb))
	}

	return &cdk8splus30.ContainerResources{
		Cpu:    &cpu,
		Memory: &memory,
	}
}

func GetServicePorts(ports []*ServicePort, metricsPort int) []*cdk8splus30.ServicePort {
	var servicePorts []*cdk8splus30.ServicePort
	i := 0
	for _, port := range ports {
		name := portName(port.GRPC, i)
		if metricsPort != 0 {
			if port.ContainerPort == metricsPort {
				name = "http-metrics"
			}
		}
		protocol := GetProtocol(port.Protocol)
		servicePorts = append(servicePorts, &cdk8splus30.ServicePort{
			Name:       jsii.String(name),
			Protocol:   protocol,
			TargetPort: jsii.Number(port.ContainerPort),
			Port:       jsii.Number(port.HostPort),
		})
		i++
	}
	return servicePorts
}

func AutoScalerProps(config *ScalingConfig, dep cdk8splus30.Deployment) *cdk8splus30.HorizontalPodAutoscalerProps {
	if config == nil {
		return nil
	}
	var maxReplicas *float64
	if config.MaxReplicas != 0 {
		maxReplicas = jsii.Number(config.MaxReplicas)
	}
	var minReplicas *float64
	if config.MinReplicas != 0 {
		minReplicas = jsii.Number(config.MinReplicas)
	}

	metrics := AutoScalerMetrics(config)

	return &cdk8splus30.HorizontalPodAutoscalerProps{
		Metadata:    nil,
		MaxReplicas: maxReplicas,
		Target:      dep,
		Metrics:     &metrics,
		MinReplicas: minReplicas,
		ScaleDown:   nil,
		ScaleUp:     nil,
	}
}

func AutoScalerRules(config *ScalingConfig) *cdk8splus30.ScalingRules {
	if config == nil {
		return nil
	}

	return &cdk8splus30.ScalingRules{
		Policies: &[]*cdk8splus30.ScalingPolicy{
			&cdk8splus30.ScalingPolicy{
				Replicas: nil,
				Duration: nil,
			},
		},
		StabilizationWindow: nil,
		Strategy:            cdk8splus30.ScalingStrategy_MIN_CHANGE,
	}
}

func AutoScalerMetrics(config *ScalingConfig) []cdk8splus30.Metric {
	if config == nil {
		return nil
	}
	var metrics []cdk8splus30.Metric
	cpuUtil, cpuAmount := AddCpuMetrics(config)
	if cpuUtil != nil {

		cpuUtilMetric := cdk8splus30.Metric_ContainerCpu(cpuUtil)
		metrics = append(metrics, cpuUtilMetric)
	}
	if cpuAmount != nil {
		cpuAmountMetric := cdk8splus30.Metric_ContainerCpu(cpuAmount)
		metrics = append(metrics, cpuAmountMetric)
	}
	memoryUtil, memoryAmount := AddMemoryMetrics(config)
	if memoryUtil != nil {
		memoryUtilMetric := cdk8splus30.Metric_ContainerMemory(memoryUtil)
		metrics = append(metrics, memoryUtilMetric)
	}
	if memoryAmount != nil {
		memoryAmountMetric := cdk8splus30.Metric_ContainerMemory(memoryAmount)
		metrics = append(metrics, memoryAmountMetric)
	}
	return metrics
}

func AddCpuMetrics(config *ScalingConfig) (*cdk8splus30.MetricContainerResourceOptions, *cdk8splus30.MetricContainerResourceOptions) {
	if config == nil {
		return nil, nil
	}
	var utilOptions cdk8splus30.MetricContainerResourceOptions
	if config.TargetCpuPercent != 0 {
		target := cdk8splus30.MetricTarget_AverageUtilization(jsii.Number(config.TargetCpuPercent))
		cpuTarget := cdk8splus30.Metric_ResourceCpu(target)
		utilOptions.Target = cpuTarget
	}
	var amountOptions cdk8splus30.MetricContainerResourceOptions
	if config.TargetCpuAmount != 0 {
		target := cdk8splus30.MetricTarget_AverageValue(jsii.Number(config.TargetCpuAmount))
		cpuTarget := cdk8splus30.Metric_ResourceCpu(target)
		amountOptions.Target = cpuTarget
	}
	return &utilOptions, &amountOptions
}

func AddMemoryMetrics(config *ScalingConfig) (*cdk8splus30.MetricContainerResourceOptions, *cdk8splus30.MetricContainerResourceOptions) {
	if config == nil {
		return nil, nil
	}
	var utilOptions cdk8splus30.MetricContainerResourceOptions
	if config.TargetMemoryPercent != 0 {
		target := cdk8splus30.MetricTarget_AverageUtilization(jsii.Number(config.TargetMemoryPercent))
		memoryTarget := cdk8splus30.Metric_ResourceMemory(target)
		utilOptions.Target = memoryTarget
	}
	var amountOptions cdk8splus30.MetricContainerResourceOptions
	if config.TargetMemoryAmount != 0 {
		target := cdk8splus30.MetricTarget_AverageValue(jsii.Number(config.TargetMemoryAmount))
		memoryTarget := cdk8splus30.Metric_ResourceMemory(target)
		amountOptions.Target = memoryTarget
	}

	return &utilOptions, &amountOptions
}
