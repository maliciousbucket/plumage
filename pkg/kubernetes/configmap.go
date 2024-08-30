package kubernetes

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2"
	"path/filepath"
)

type LocalDirectoryVolume interface {
	Name() string
	Directory() string
	Target() string
}

func LocalDirToConfigMap(scope constructs.Construct, v LocalDirectoryVolume, props cdk8splus30.ConfigMapProps, workDir string) cdk8splus30.ConfigMap {
	id := v.Name()

	configMap := cdk8splus30.NewConfigMap(scope, &id, &props)

	path := filepath.Join(workDir, v.Directory())
	configMap.AddDirectory(&path, nil)

	return configMap
}

type LocalFileVolume interface {
	Name() string
	File() string
	Target() string
}

func LocalFileToConfigMap(scope constructs.Construct, v LocalFileVolume, props cdk8splus30.ConfigMapProps, workDir string) cdk8splus30.ConfigMap {
	id := v.Name()

	configMap := cdk8splus30.NewConfigMap(scope, &id, &props)

	path := filepath.Join(workDir, v.File())
	configMap.AddFile(&path, nil)

	return configMap
}
