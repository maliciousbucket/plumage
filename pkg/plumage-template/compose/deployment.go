package compose

import (
	"fmt"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/maliciousbucket/plumage/imports/k8s"
	plumagetemplate "github.com/maliciousbucket/plumage/pkg/plumage-template"
	"github.com/maliciousbucket/plumage/pkg/types"
)

type DeploymentProps struct {
	Name           string
	Image          string
	Commands       []string
	Args           []string
	StartupProbe   *types.CommandProbe
	HealthCheck    *plumagetemplate.HttpProbe
	Ports          []plumagetemplate.ServicePort
	Resources      *plumagetemplate.ServiceResources
	Monitoring     *plumagetemplate.MonitoringConfig
	InitContainers []*plumagetemplate.InitContainer
	MinReplicas    int
	Env            map[string]string
}

func NewDeployment(scope constructs.Construct, id string, ns string, appLabel string, props *DeploymentProps) k8s.KubeDeployment {
	meta := deploymentMeta(ns, props.Name)
	labels := deploymentLabels(props.Name)

	var replicas float64
	if props.MinReplicas != 0 {
		replicas = float64(props.MinReplicas)
	}

	container := NewContainer(&ContainerProps{
		Name:         props.Name,
		Image:        props.Image,
		Commands:     props.Commands,
		Args:         props.Args,
		StartupProbe: props.StartupProbe,
		HealthCheck:  props.HealthCheck,
		Ports:        props.Ports,
		Resources:    props.Resources,
		Monitoring:   props.Monitoring,
		Env:          props.Env,
		WorkingDir:   "",
	})

	initContainers := InitContainersToK8s(props.InitContainers)

	deployment := k8s.NewKubeDeployment(scope, jsii.String(id), &k8s.KubeDeploymentProps{
		Metadata: meta,
		Spec: &k8s.DeploymentSpec{
			Selector: &k8s.LabelSelector{
				MatchLabels: labels,
			},
			Template: &k8s.PodTemplateSpec{
				Metadata: &k8s.ObjectMeta{
					Annotations: nil,
					Labels:      labels,
					Name:        deploymentName(appLabel),
					Namespace:   jsii.String(ns),
				},
				Spec: &k8s.PodSpec{
					Containers: &[]*k8s.Container{
						&container,
					},
					ActiveDeadlineSeconds:        nil,
					AutomountServiceAccountToken: nil,
					HostAliases:                  nil,
					InitContainers:               initContainers,
					PriorityClassName:            nil,
					ReadinessGates:               nil,
					RestartPolicy:                jsii.String("Always"),
					Volumes: &[]*k8s.Volume{
						{
							Name:                 nil,
							AwsElasticBlockStore: nil,
							AzureDisk:            nil,
							AzureFile:            nil,
							Cephfs:               nil,
							Cinder:               nil,
							ConfigMap: &k8s.ConfigMapVolumeSource{
								DefaultMode: nil,
								Items:       nil,
								Name:        nil,
								Optional:    nil,
							},
							Csi:                   nil,
							DownwardApi:           nil,
							EmptyDir:              nil,
							Ephemeral:             nil,
							Fc:                    nil,
							FlexVolume:            nil,
							Flocker:               nil,
							GcePersistentDisk:     nil,
							GitRepo:               nil,
							Glusterfs:             nil,
							HostPath:              nil,
							Iscsi:                 nil,
							Nfs:                   nil,
							PersistentVolumeClaim: nil,
							PhotonPersistentDisk:  nil,
							PortworxVolume:        nil,
							Projected:             nil,
							Quobyte:               nil,
							Rbd:                   nil,
							ScaleIo:               nil,
							Secret:                nil,
							Storageos:             nil,
							VsphereVolume:         nil,
						},
					},
				},
			},
			MinReadySeconds:         nil,
			Paused:                  nil,
			ProgressDeadlineSeconds: nil,
			Replicas:                &replicas,
			RevisionHistoryLimit:    nil,
			Strategy:                nil,
		},
	})

	return deployment
}

func deploymentMeta(ns string, appLabel string) *k8s.ObjectMeta {
	//name := fmt.Sprintf("deployment-%s", appLabel)
	labels := deploymentLabels(appLabel)

	return &k8s.ObjectMeta{
		Annotations: nil,
		Labels:      labels,
		Name:        jsii.String(appLabel),
		Namespace:   jsii.String(ns),
	}
}

func deploymentLabels(appLabel string) *map[string]*string {
	return &map[string]*string{
		AppLabel: jsii.String(appLabel),
	}
}

func deploymentName(appLabel string) *string {
	name := fmt.Sprintf("deployment-%s", appLabel)
	return &name
}
