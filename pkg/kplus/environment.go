package kplus

import (
	"fmt"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	kplus "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2"
)

func AddServiceEnvironmentVariables(scope constructs.Construct, s *ServiceTemplate, monitoring map[string]string) []kplus.EnvFrom {

	monitoringEnvConfig, envConfig := loadMonitoringEnvWithAliases(monitoring, s.Env, s.Env)

	mName := fmt.Sprintf("%s-monitoring-env-values", s.Name)

	configMap := kplus.NewConfigMap(scope, jsii.String(mName), nil)

	for k, v := range monitoringEnvConfig {
		configMap.AddData(jsii.String(k), jsii.String(v))
	}

	envConfigMap := kplus.NewConfigMap(scope, jsii.String(mName), nil)
	for k, v := range envConfig {
		envConfigMap.AddData(jsii.String(k), jsii.String(v))
	}

	monitoringEnvFrom := kplus.Env_FromConfigMap(configMap, nil)
	envFrom := kplus.Env_FromConfigMap(envConfigMap, nil)

	result := []kplus.EnvFrom{monitoringEnvFrom, envFrom}

	return result
}

func loadMonitoringEnvWithAliases(values, aliases map[string]string, env map[string]string) (map[string]string, map[string]string) {
	result := make(map[string]string)
	envMap := make(map[string]string)
	for k, v := range values {
		if _, ok := env[k]; ok {
			result[k] = v
		} else if configValue, found := aliases[k]; found {
			result[configValue] = v
		} else {
			envMap[k] = v
		}
	}
	return result, envMap

}
