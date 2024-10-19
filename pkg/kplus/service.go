package kplus

import (
	"fmt"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	kplus "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2"
	"github.com/joho/godotenv"
	"log"
	"strconv"
	"strings"
)

func NewServiceManifests(scope constructs.Construct, id string, ns string, template *ServiceTemplate, monitoring map[string]string, nodePort int) constructs.Construct {
	ct := constructs.NewConstruct(scope, jsii.String(id))

	deploymentName := "deployment"
	deployment := newServiceDeployment(ct, deploymentName, template, monitoring)

	labelSelector := kplus.LabelSelector_Of(&kplus.LabelSelectorOptions{Labels: &map[string]*string{"app": jsii.String(template.Name)}})

	deployment.Select(labelSelector)

	serviceName := fmt.Sprintf("%s", template.Name)
	service := deployment.ExposeViaService(&kplus.DeploymentExposeViaServiceOptions{
		ServiceType: kplus.ServiceType_CLUSTER_IP,
		Name:        jsii.String(serviceName),
	})
	service.Metadata().AddLabel(jsii.String("app"), jsii.String(template.Name))
	service.SelectLabel(jsii.String("app"), jsii.String(template.Name))

	accountName := fmt.Sprintf("%s-account", template.Name)
	kplus.NewServiceAccount(ct, jsii.String(accountName), &kplus.ServiceAccountProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Name: jsii.String(serviceName),
		},
	})

	var middlewareRefs []string

	if template.CircuitBreaker != nil {
		cbName := fmt.Sprintf("%s-circuitbreaker", template.Name)
		circuitBreaker := NewCircuitBreaker(ct, cbName, ns, template)
		if circuitBreaker != nil {
			middlewareRefs = append(middlewareRefs, cbName)
		}
	}

	if template.Retry != nil {
		retryName := fmt.Sprintf("%s-retry", template.Name)
		retry := NewRetry(ct, retryName, ns, template)
		if retry != nil {
			middlewareRefs = append(middlewareRefs, retryName)
		}
	}

	if template.RateLimit != nil {
		rlName := fmt.Sprintf("%s-ratelimit", template.Name)
		rl := NewRateLimit(ct, rlName, ns, template)
		if rl != nil {
			middlewareRefs = append(middlewareRefs, rlName)
		}
	}

	if template.Scaling != nil {
		NewAutoScaler(ct, deployment, template.Scaling, template.Name)
	}

	if len(template.DefaultMiddleware) > 0 {
		log.Println("Adding default middleware")
		defaults := NewDefaultMiddlewares(ct, ns, template.Name, template.DefaultMiddleware)
		if len(defaults) > 0 {
			middlewareRefs = append(middlewareRefs, defaults...)
		}
	}

	if template.DefaultAutoScaling.DefaultScaling != nil && template.Scaling == nil {
		name := fmt.Sprintf("%s-autoscaler", template.Name)
		AddDefaultScaling(ct, deployment, name, template.DefaultAutoScaling)
	}
	if template.Paths != nil && len(template.Paths) > 0 {
		routeName := fmt.Sprintf("%s-ingress-route", template.Name)
		NewIngressRoute(ct, routeName, ns, template, middlewareRefs)
	}

	return ct
}

func newServiceDeployment(scope constructs.Construct, id string, service *ServiceTemplate, m map[string]string) kplus.Deployment {

	deployment := kplus.NewDeployment(scope, jsii.String(id), &kplus.DeploymentProps{
		SecurityContext: &kplus.PodSecurityContextProps{
			EnsureNonRoot: jsii.Bool(false),
		},
	})
	deployment.PodMetadata().AddLabel(jsii.String("app"), jsii.String(service.Name))
	containerProps := newContainerProps(scope, service, m)
	container := deployment.AddContainer(containerProps)

	addEmptyDirs(scope, container, service.Name, service.EmptyDirs)
	addFileMounts(scope, container, service)
	addDirectoryMounts(scope, container, service)

	if service.Monitoring != nil {
		if service.Monitoring.ScrapePort != 0 || service.Monitoring.ScrapePath != "" {
			deployment.Metadata().AddAnnotation(jsii.String("prometheus.io/scrape"), jsii.String("true"))
		}

		if service.Monitoring.ScrapePort != 0 {
			deployment.Metadata().AddAnnotation(jsii.String("prometheus.io/port"), jsii.String(strconv.FormatInt(int64(service.Monitoring.ScrapePort), 10)))
		}
		if service.Monitoring.ScrapePath != "" {
			deployment.Metadata().AddAnnotation(jsii.String("prometheus.io/path"), jsii.String(service.Monitoring.ScrapePath))
		}
	}

	deployment.Metadata().AddLabel(jsii.String("app"), jsii.String(service.Name))

	return deployment
}

func newContainerProps(scope constructs.Construct, service *ServiceTemplate, monitoring map[string]string) *kplus.ContainerProps {
	metricsPort := 0
	if service.Monitoring != nil {
		metricsPort = service.Monitoring.ScrapePort
	}

	ports := ContainerPorts(service, metricsPort)

	args := StringSliceToK8s(service.Args)
	commands := StringSliceToK8s(service.Commands)
	props := &kplus.ContainerProps{

		Image:           jsii.String(service.Image),
		Args:            args,
		Command:         commands,
		Ports:           &ports,
		ImagePullPolicy: kplus.ImagePullPolicy_IF_NOT_PRESENT,
		SecurityContext: &kplus.ContainerSecurityContextProps{
			EnsureNonRoot: jsii.Bool(false),
		},
	}

	if service.LivenessProbe.Probe != nil {

		props.Liveness = ToKplusProbe(service.LivenessProbe.Probe)
	}

	if service.ReadinessProbe.Probe != nil {
		props.Readiness = ToKplusProbe(service.ReadinessProbe.Probe)
	}

	if service.ReadinessProbe.Probe != nil {
		props.Readiness = ToKplusProbe(service.ReadinessProbe.Probe)
	}

	if service.WorkingDir != "" {
		props.WorkingDir = jsii.String(service.WorkingDir)
	}
	props.EnvFrom = &[]kplus.EnvFrom{}

	if service.Env != nil && len(service.Env) > 0 || service.Monitoring != nil {
		configMap := AddEnvironmentVariables(scope, service, monitoring)
		envFrom := kplus.NewEnvFrom(configMap, nil, nil)

		containerEnvFrom := append(*props.EnvFrom, envFrom)
		props.EnvFrom = &containerEnvFrom
	}

	if service.Resources != nil {
		resources := toKplusResources(service.Resources)
		props.Resources = resources
	}

	return props

}

func ContainerPorts(service *ServiceTemplate, metricsPort int) []*kplus.ContainerPort {
	ports := []*kplus.ContainerPort{}
	i := 0
	for _, port := range service.Ports {
		protocol := kplus.Protocol_TCP
		if strings.ToLower(port.Protocol) == "udp" {
			protocol = kplus.Protocol_UDP
		}

		if port.Name == "" {
			port.Name = portName(port, i)
			i++
		}

		if metricsPort != 0 && port.Port == metricsPort {
			port.Name = "http-metrics"
		}

		ports = append(ports, &kplus.ContainerPort{
			Number:   jsii.Number(port.Port),
			Name:     jsii.String(port.Name),
			Protocol: protocol,
		})
	}
	return ports
}

func containerEnv(scope constructs.Construct, name string, env map[string]string, envFile string) kplus.ConfigMap {
	configMapName := fmt.Sprintf("%s-configmap", name)
	configmap := kplus.NewConfigMap(scope, jsii.String(configMapName), &kplus.ConfigMapProps{})

	var fileEnv map[string]string
	if envFile != "" {
		fileEnv, _ = godotenv.Read(envFile)
	}

	if fileEnv != nil {
		for key, value := range fileEnv {
			configmap.AddData(jsii.String(key), jsii.String(value))
		}
	}

	for key, value := range env {
		configmap.AddData(jsii.String(key), jsii.String(value))
	}
	return configmap

}

func addFileMounts(scope constructs.Construct, container kplus.Container, service *ServiceTemplate) kplus.Container {

	if len(service.FileMounts) > 0 {

		for i := 0; i < len(service.FileMounts); i++ {
			if len(service.FileMounts[i]) > 0 {
				name := fmt.Sprintf("%s-file-configmap-%d", service.Name, i)
				configMap := kplus.NewConfigMap(scope, jsii.String(name), nil)
				for k, v := range service.FileMounts[i] {
					configMap.AddFile(jsii.String(k), jsii.String(v))
				}
				mountId := fmt.Sprintf("%s-volumeMount", *configMap.Name())
				mount := kplus.Volume_FromConfigMap(scope, jsii.String(mountId), configMap, &kplus.ConfigMapVolumeOptions{})
				container.Mount(jsii.String(""), mount, nil)

			}
		}
	}

	return container
}

func addDirectoryMounts(scope constructs.Construct, container kplus.Container, service *ServiceTemplate) kplus.Container {
	if len(service.VolumeMounts) > 0 {
		i := 0
		for local, target := range service.VolumeMounts {
			name := fmt.Sprintf("%s-directory-configmap-%d", service.Name, i)
			configMap := kplus.NewConfigMap(scope, jsii.String(name), nil)
			configMap.AddDirectory(jsii.String(local), nil)
			mountId := fmt.Sprintf("%s-volumeMount", *configMap.Name())
			mount := kplus.Volume_FromConfigMap(scope, jsii.String(mountId), configMap, &kplus.ConfigMapVolumeOptions{})
			container.Mount(jsii.String(""), mount, &kplus.MountOptions{
				SubPath: jsii.String(target),
			})

		}
	}

	return container
}

func addEmptyDirs(scope constructs.Construct, container kplus.Container, name string, emptyDirs []string) kplus.Container {
	fmt.Println("Adding dirs")
	var emptyVolumes []*kplus.Volume
	if len(emptyDirs) > 0 {
		i := 0
		for _, emptyDir := range emptyDirs {
			dirName := fmt.Sprintf("%s-empty-dir-%d", name, i)
			mount := kplus.Volume_FromEmptyDir(scope, jsii.String(dirName), jsii.String(dirName), nil)
			emptyVolumes = append(emptyVolumes, &mount)

			container.Mount(jsii.String(emptyDir), mount, nil)
		}
	}
	return container
}

func toKplusResources(resources *Resources) *kplus.ContainerResources {
	if resources == nil {
		return nil
	}
	containerResources := &kplus.ContainerResources{
		Cpu:              nil,
		EphemeralStorage: nil,
		Memory:           nil,
	}
	if resources.Memory != nil {
		containerResources.Memory = &kplus.MemoryResources{}
		if resources.Memory.Request != 0 {
			containerResources.Memory.Request = cdk8s.Size_Mebibytes(jsii.Number(resources.Memory.Request))
		}
		if resources.Memory.Limit != 0 {
			containerResources.Memory.Limit = cdk8s.Size_Mebibytes(jsii.Number(resources.Memory.Limit))
		}
	}

	if resources.CPU != nil {
		containerResources.Cpu = &kplus.CpuResources{}
		if resources.CPU.Request != 0 {
			containerResources.Cpu.Request = kplus.Cpu_Millis(jsii.Number(resources.CPU.Request))
		}
		if resources.CPU.Limit != 0 {
			containerResources.Cpu.Limit = kplus.Cpu_Millis(jsii.Number(resources.CPU.Limit))
		}
	}
	return containerResources
}

func portName(port Port, num int) string {
	protocol := "http"
	if port.GRPC {
		protocol = "grpc"
	}
	return fmt.Sprintf("%s-%d", protocol, num)
}

func StringMapToK8s(m map[string]string) *map[string]*string {
	k8sMap := make(map[string]*string)
	for k, v := range m {
		k8sMap[k] = &v
	}
	return &k8sMap
}

func StringSliceToK8s(sl []string) *[]*string {
	var k8sSlice []*string
	for _, v := range sl {
		k8sSlice = append(k8sSlice, &v)
	}
	return &k8sSlice
}
