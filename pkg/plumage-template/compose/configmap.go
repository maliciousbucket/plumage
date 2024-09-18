package compose

import (
	"fmt"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	kplus "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2"
	"github.com/maliciousbucket/plumage/imports/k8s"
	plumagetemplate "github.com/maliciousbucket/plumage/pkg/plumage-template"
)

func VolumeMounts(scope constructs.Construct, svcName string, volumes []plumagetemplate.ContainerVolume) ([]*k8s.Volume, []*k8s.VolumeMount) {
	i := 0
	podVolumes := []*k8s.Volume{}
	containerVolumes := []*k8s.VolumeMount{}
	for _, volume := range volumes {

		if fileVolume, ok := volume.(*plumagetemplate.FileVolume); ok {
			name := fmt.Sprintf("%s-%s-%d", svcName, fileVolume.Name, i)
			if fileVolume.Name == "" {
				fileVolume.Name = name
			}
			configMap := newFileConfigMap(scope, fileVolume)
			podVolume, containerVolume := newVolumeMount(fileVolume, *configMap.Name())
			podVolumes = append(podVolumes, podVolume)
			containerVolumes = append(containerVolumes, containerVolume)
		}

		if dirVolume, ok := volume.(*plumagetemplate.DirVolume); ok {
			name := fmt.Sprintf("%s-%s-%d", svcName, dirVolume.Name, i)
			if dirVolume.Name == "" {
				dirVolume.Name = name
			}
			configMap := newDirConfigMap(scope, dirVolume)
			podVolume, containerVolume := newVolumeMount(dirVolume, *configMap.Name())
			podVolumes = append(podVolumes, podVolume)
			containerVolumes = append(containerVolumes, containerVolume)
		}

		if emptyVolume, ok := volume.(*plumagetemplate.EmptyVolume); ok {
			name := fmt.Sprintf("%s-%s-empty-%d", svcName, emptyVolume.Name, i)
			if emptyVolume.Name == "" {
				emptyVolume.Name = name
			}
			podVolume, containerVolume := newVolumeMount(emptyVolume, "")
			podVolumes = append(podVolumes, podVolume)
			containerVolumes = append(containerVolumes, containerVolume)
		}
		i++
	}
	return podVolumes, containerVolumes
}

func newVolumeMount(serviceVolume plumagetemplate.ContainerVolume, mapName string) (*k8s.Volume, *k8s.VolumeMount) {
	if emptyVolume, ok := serviceVolume.(*plumagetemplate.EmptyVolume); ok {
		volume := newEmptyDir(emptyVolume)
		containerVolume := &k8s.VolumeMount{
			MountPath: jsii.String(emptyVolume.TargetVolume()),
			Name:      volume.Name,
		}
		return volume, containerVolume
	}

	if fileVolume, ok := serviceVolume.(*plumagetemplate.FileVolume); ok {
		volume := &k8s.Volume{
			Name: jsii.String(fileVolume.Name),
			ConfigMap: &k8s.ConfigMapVolumeSource{
				Name: jsii.String(mapName),
			},
		}
		containerVolume := &k8s.VolumeMount{

			MountPath: jsii.String(fileVolume.TargetVolume()),
			Name:      volume.Name,
		}
		return volume, containerVolume
	}

	if dirVolume, ok := serviceVolume.(*plumagetemplate.DirVolume); ok {

		volume := &k8s.Volume{
			Name: jsii.String(dirVolume.Name),
			ConfigMap: &k8s.ConfigMapVolumeSource{
				Name: jsii.String(mapName),
			},
		}
		containerVolume := &k8s.VolumeMount{
			MountPath: jsii.String(dirVolume.TargetVolume()),
			Name:      volume.Name,
		}
		return volume, containerVolume
	}
	return nil, nil
}

func newFileConfigMap(scope constructs.Construct, volume *plumagetemplate.FileVolume) k8s.KubeConfigMap {
	data := map[string]*string{
		"-": jsii.String(volume.Source),
	}
	name := fmt.Sprintf("%s-configMap", volume.Name)

	configMap := k8s.NewKubeConfigMap(scope, jsii.String(name), &k8s.KubeConfigMapProps{
		Data:     &data,
		Metadata: &k8s.ObjectMeta{Name: jsii.String(name)},
	})
	return configMap
}

func newDirConfigMap(scope constructs.Construct, volume *plumagetemplate.DirVolume) k8s.KubeConfigMap {
	data := map[string]*string{
		"-": jsii.String(volume.Source),
	}
	name := fmt.Sprintf("%s-configMap", volume.Name)

	configMap := k8s.NewKubeConfigMap(scope, jsii.String(name), &k8s.KubeConfigMapProps{
		Data:     &data,
		Metadata: &k8s.ObjectMeta{Name: jsii.String(name)},
	})
	return configMap
}

func newEmptyDir(volume *plumagetemplate.EmptyVolume) *k8s.Volume {
	return &k8s.Volume{
		EmptyDir: &k8s.EmptyDirVolumeSource{},
		Name:     jsii.String(volume.Name),
	}
}

func ServiceFileConfigMap(scope constructs.Construct, name string, mount map[string]string) kplus.ConfigMap {
	configMap := kplus.NewConfigMap(scope, jsii.String(name), nil)
	for k, v := range mount {
		configMap.AddFile(jsii.String(k), jsii.String(v))
	}
	return configMap
}

func ServiceDirConfigMap(scope constructs.Construct, name string, mount map[string]string) kplus.ConfigMap {
	configMap := kplus.NewConfigMap(scope, jsii.String(name), nil)
	for k, _ := range mount {
		configMap.AddDirectory(jsii.String(k), nil)
	}
	return configMap
}
