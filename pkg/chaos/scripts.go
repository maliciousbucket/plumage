package chaos

import (
	"fmt"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	kplus "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2"
)

var (
	testAnnotation    = "galah-monitoring.io/test"
	testEnvAnnotation = "galah-monitoring.io/test-env"
)

func newTestFileConfigMap(scope constructs.Construct, ns, name, scriptDir, libDir, scriptName string, libFiles []string) kplus.ConfigMap {
	configMap := kplus.NewConfigMap(scope, jsii.String(name), &kplus.ConfigMapProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Name:      jsii.String(name),
			Namespace: jsii.String(ns),
		},
	})
	addScript(configMap, scriptDir, scriptName)
	addLibFiles(configMap, libDir, scriptDir, libFiles)
	configMap.Metadata().AddAnnotation(jsii.String(testAnnotation), jsii.String(name))
	return configMap
}

func addScript(configMap kplus.ConfigMap, scriptDir, script string) kplus.ConfigMap {
	file := fmt.Sprintf("%s/%s", scriptDir, script)
	configMap.AddFile(jsii.String(file), jsii.String(script))
	return configMap
}

func addLibFiles(configMap kplus.ConfigMap, libDir, scriptDir string, libFiles []string) kplus.ConfigMap {
	if libDir == "" && libFiles == nil {
		return configMap
	}
	if libDir == "" && len(libFiles) == 0 {
		return configMap
	}

	if libDir != "" && len(libFiles) > 0 {
		for _, libFile := range libFiles {
			file := fmt.Sprintf("%s/%s", libDir, libFile)
			configMap.AddFile(jsii.String(file), jsii.String(libFile))
		}
		return configMap
	}

	if libDir != "" && len(libFiles) == 0 {
		configMap.AddDirectory(jsii.String(libDir), nil)
		return configMap
	}
	if libDir == "" && len(libFiles) > 0 {
		for _, libFile := range libFiles {
			file := fmt.Sprintf("%s/%s", scriptDir, libFile)
			configMap.AddFile(jsii.String(file), jsii.String(libFile))
		}
		return configMap
	}
	return configMap
}

func newEnvConfigMap(scope constructs.Construct, name, ns string, env map[string]string) kplus.ConfigMap {
	mapName := fmt.Sprintf("%s-env", name)
	configMap := kplus.NewConfigMap(scope, jsii.String(name), &kplus.ConfigMapProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Name:      jsii.String(mapName),
			Namespace: jsii.String(ns),
		},
	})

	for k, v := range env {
		configMap.AddData(jsii.String(k), jsii.String(v))
	}
	configMap.Metadata().AddAnnotation(jsii.String(testEnvAnnotation), jsii.String(name))
	return configMap
}
