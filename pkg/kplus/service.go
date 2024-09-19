package kplus

import (
	"fmt"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	kplus "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2"
	"github.com/joho/godotenv"
	"strconv"
	"strings"
)

func NewServiceManifests(scope constructs.Construct, id string, template *ServiceTemplate) constructs.Construct {
	ct := constructs.NewConstruct(scope, jsii.String(id))

	deploymentName := fmt.Sprintf("%s-deployment", template.Name)
	deployment := newServiceDeployment(ct, deploymentName, template)

	labelSelector := kplus.LabelSelector_Of(&kplus.LabelSelectorOptions{Labels: &map[string]*string{"app": jsii.String(template.Name)}})

	deployment.Select(labelSelector)

	serviceName := fmt.Sprintf("%s-service", template.Name)
	service := deployment.ExposeViaService(&kplus.DeploymentExposeViaServiceOptions{
		ServiceType: kplus.ServiceType_CLUSTER_IP,
		Name:        jsii.String(serviceName),
	})
	service.SelectLabel(jsii.String("app"), jsii.String(template.Name))
	accountName := fmt.Sprintf("%s-account", template.Name)
	kplus.NewServiceAccount(ct, jsii.String(accountName), &kplus.ServiceAccountProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Name: jsii.String(serviceName),
		},
	})

	var mioddlewareRefs []string

	if template.CircuitBreaker != nil {
		cbName := fmt.Sprintf("%s-circuitbreaker", template.Name)
		circuitBreaker := NewCircuitBreaker(ct, cbName, template)
		if circuitBreaker != nil {
			mioddlewareRefs = append(mioddlewareRefs, *circuitBreaker.Name())
		}
	}

	if template.Retry != nil {
		retryName := fmt.Sprintf("%s-retry", template.Name)
		retry := NewRetry(ct, retryName, template)
		if retry != nil {
			mioddlewareRefs = append(mioddlewareRefs, *retry.Name())
		}
	}

	if template.RateLimit != nil {
		rlName := fmt.Sprintf("%s-ratelimit", template.Name)
		rl := NewRateLimit(ct, rlName, template)
		if rl != nil {
			mioddlewareRefs = append(mioddlewareRefs, *rl.Name())
		}
	}

	if template.Scaling != nil {
		NewAutoScaler(ct, deployment, template.Scaling, template.Name)
	}
	routeName := fmt.Sprintf("%s-ingress-route", template.Name)
	NewIngressRoute(ct, routeName, template, mioddlewareRefs)

	return ct
}

func newServiceDeployment(scope constructs.Construct, id string, service *ServiceTemplate) kplus.Deployment {
	deployment := kplus.NewDeployment(scope, jsii.String(id), nil)

	containerProps := newContainerProps(scope, service)
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

func newContainerProps(scope constructs.Construct, service *ServiceTemplate) *kplus.ContainerProps {
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
	}

	if service.LivenessProbe != nil {

		props.Liveness = ToKplusProbe(service.LivenessProbe)
	}

	if service.ReadinessProbe != nil {
		props.Readiness = ToKplusProbe(service.ReadinessProbe)
	}

	if service.ReadinessProbe != nil {
		props.Readiness = ToKplusProbe(service.ReadinessProbe)
	}

	if len(service.Env) != 0 || service.EnvFile != "" {
		env := containerEnv(scope, service.Name, service.Env, service.EnvFile)
		source := kplus.Env_FromConfigMap(env, nil)
		props.EnvFrom = &[]kplus.EnvFrom{
			source,
		}
	}

	if service.WorkingDir != "" {
		props.WorkingDir = jsii.String(service.WorkingDir)
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
	if len(emptyDirs) > 0 {
		i := 0
		for _, emptyDir := range emptyDirs {
			name := fmt.Sprintf("%s-empty-dir-%d", name, i)
			mount := kplus.Volume_FromEmptyDir(scope, jsii.String(name), jsii.String(name), nil)
			container.Mount(jsii.String(emptyDir), mount, nil)
		}
	}
	return container
}

func portName(port Port, num int) string {
	protocol := "http"
	if port.GRPC {
		protocol = "grpc"
	}
	return fmt.Sprintf("%s-%d", protocol, num)
}

func StringMapToK8s(m map[string]string) *map[string]*string {
	var k8sMap map[string]*string
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
