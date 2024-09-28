package manifests

import (
	"encoding/json"
	"fmt"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	kplus "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2"
	"github.com/maliciousbucket/plumage/imports/k8s"
	"slices"
)

const (
	containerPath = "/spec/template/spec/containers/0/envFrom"
)

func AddServiceEnvironmentVariables(scope constructs.Construct, s *ServiceTemplate, target cdk8s.ApiObject, monitoring map[string]string) constructs.Construct {

	envConfig := loadMonitoringEnvWithAliases(monitoring, s.MonitoringAliases, s.MonitoringEnv)
	envValues := StringMapToK8s(envConfig)
	fmt.Printf("Env Values: %v", envValues)

	addContainerEnvironmentVariables(scope, target, s.Namespace, s.Name, envValues)
	return scope
}

func addContainerEnvironmentVariables(scope constructs.Construct, target cdk8s.ApiObject, ns string, appLabel string, data *map[string]*string) constructs.Construct {
	id := fmt.Sprintf("%s-config-env", appLabel)

	cfgMap := kplus.NewConfigMap(scope, jsii.String(id), &kplus.ConfigMapProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Name:      jsii.String(id),
			Namespace: jsii.String(ns),
		},
		Data: data,
	})

	patch := addContainerEnvPatch(cfgMap)

	target.AddJsonPatch(patch)

	return target

}

func addContainerEnvPatch(cfgMap kplus.ConfigMap) cdk8s.JsonPatch {

	envFrom := k8s.EnvFromSource{ConfigMapRef: &k8s.ConfigMapEnvSource{Name: cfgMap.Name()}}

	patch := cdk8s.JsonPatch_Add(jsii.String(containerPath), envFrom)

	return patch
}

func StructToEnvMap(obj interface{}) map[string]*string {
	var result map[string]*string
	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(obj)
	json.Unmarshal(inrec, &inInterface)

	for k, v := range inInterface {
		result[k] = v.(*string)
	}
	return result
}

func loadMonitoringEnv(values, config, env map[string]string) map[string]string {
	for k, v := range values {
		if _, ok := env[k]; ok {
			env[k] = v
		}
	}
	//If a value in the config map is a key in the values map
	//Set  the key from the config map - to the value of the value map's value
	// In the env map
	//So that some keys cna be provided for non-standard otel etc env variables
	for key, configValue := range config {
		if v, ok := values[configValue]; ok {
			env[key] = v
		}
	}
	return env
}

func loadMonitoringEnvWithAliases(values, aliases map[string]string, env []string) map[string]string {
	result := make(map[string]string)
	for k, v := range values {
		if slices.Contains(env, k) {
			result[k] = v
		}
		if configValue, ok := aliases[k]; ok {
			result[configValue] = v
		}
	}
	return result

}

func StringMapToK8s(m map[string]string) *map[string]*string {
	k8sMap := make(map[string]*string)
	for k, v := range m {
		k8sMap[k] = &v
	}
	return &k8sMap
}
