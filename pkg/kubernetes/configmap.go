package kubernetes

import (
	"fmt"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2"
	"github.com/maliciousbucket/plumage/pkg/types"
)

type LocalDirectoryVolume interface {
	Name() string
	Directory() string
	Target() string
}

func LocalDirToConfigMap(scope constructs.Construct, v LocalDirectoryVolume, props *cdk8splus30.ConfigMapProps) cdk8splus30.ConfigMap {
	id := v.Name()
	target := v.Target()
	directory := v.Directory()

	configMapId := fmt.Sprintf("%s-configmap", id)

	configMap := cdk8splus30.NewConfigMap(scope, jsii.String(configMapId), props)

	//path := filepath.Join(workDir, v.Directory())
	configMap.AddDirectory(jsii.String(directory), &cdk8splus30.AddDirectoryOptions{
		KeyPrefix: jsii.String(target),
	})

	return configMap
}

type LocalFileVolume interface {
	Name() string
	File() string
	Target() string
}

func LocalFileToConfigMap(scope constructs.Construct, v LocalFileVolume, props *cdk8splus30.ConfigMapProps) cdk8splus30.ConfigMap {
	id := v.Name()
	target := v.Target()
	file := v.File()

	configMapId := fmt.Sprintf("%s-configmap", id)

	configMap := cdk8splus30.NewConfigMap(scope, jsii.String(configMapId), props)

	configMap.AddFile(jsii.String(file), jsii.String(target))

	return configMap

}

type EmptyVolume interface {
	Name() string
	Target() string
}

func EmptyDirToConfigMap(scope constructs.Construct, v EmptyVolume, props *cdk8splus30.ConfigMapProps) cdk8splus30.ConfigMap {
	id := v.Name()
	target := v.Target()
	configMapId := fmt.Sprintf("%s-configmap", id)

	configMap := cdk8splus30.NewConfigMap(scope, jsii.String(configMapId), props)
	configMap.AddDirectory(nil, &cdk8splus30.AddDirectoryOptions{
		KeyPrefix: &target,
	})

	return configMap
}

func AddVolumeConfigMaps(scope constructs.Construct, container cdk8splus30.Container, namespace string, volumes []*types.ContainerVolume) (constructs.Construct, error) {
	props := defaultConfigMapProps(namespace)
	i := 0
	for _, v := range volumes {

		volume := *v
		volumeType := volume.Type()

		switch volumeType {
		case "empty":
			emptyVolume, ok := volume.(EmptyVolume)
			if !ok {
				return nil, fmt.Errorf("expected EmptyVolume type but got %T", volume)
			}
			//name := volume.Name()
			//configMap := cdk8splus30.NewConfigMap(scope, jsii.String(emptyVolume.Name()), props)
			//data := cdk8splus30.Volume_FromEmptyDir(configMap)
			//EmptyDirToConfigMap(scope, emptyVolume, props)
			container.Mount(jsii.String(emptyVolume.Target()), cdk8splus30.Volume_FromEmptyDir(scope, jsii.String("data"), jsii.String("data"), nil), nil)

		case "file":
			fileVolume, ok := volume.(LocalFileVolume)
			if !ok {
				return nil, fmt.Errorf("expected LocalFileVolume type but got %T", volume)
			}
			configMap := LocalFileToConfigMap(scope, fileVolume, props)
			containerVolume := cdk8splus30.Volume_FromConfigMap(scope, jsii.String(fileVolume.Name()), configMap, nil)
			container.Mount(jsii.String(fileVolume.Target()), containerVolume, nil)

		case "directory":
			dirVolume, ok := volume.(LocalDirectoryVolume)
			if !ok {
				return nil, fmt.Errorf("expected VolumeFromDirectory type but got %T", volume)
			}
			configMap := LocalDirToConfigMap(scope, dirVolume, props)
			containerVolume := cdk8splus30.Volume_FromConfigMap(scope, jsii.String(dirVolume.Name()), configMap, nil)
			container.Mount(jsii.String(dirVolume.Target()), containerVolume, nil)

		default:
			return nil, fmt.Errorf("unsupported volume type: %s", volumeType)
		}
		i++
	}
	return scope, nil
}

func defaultConfigMapMetaData(namespace string) cdk8s.ApiObjectMetadata {
	return cdk8s.ApiObjectMetadata{
		Namespace: jsii.String(namespace),
	}
}

func defaultConfigMapProps(namespace string) *cdk8splus30.ConfigMapProps {
	metadata := defaultConfigMapMetaData(namespace)
	return &cdk8splus30.ConfigMapProps{
		Metadata: &metadata,
	}

}
