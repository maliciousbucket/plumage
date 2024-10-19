package kplus

import (
	"fmt"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	kplus "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2"
	"slices"
)

func AddEnvironmentVariables(scope constructs.Construct, s *ServiceTemplate, monitoring map[string]string) kplus.ConfigMap {
	var monitoringEnvConfig = make(map[string]string)
	if s.Monitoring != nil {
		monitoringEnvConfig = loadMonitoringEnvWithAliases(monitoring, s.Monitoring.Aliases, s.Monitoring.MonitoringEnv)
	}

	env := loadEnv(monitoringEnvConfig, s.Env)
	configMap := kplus.NewConfigMap(scope, jsii.String("env"), nil)

	for k, v := range env {
		value := fmt.Sprintf("%s", v)
		configMap.AddData(jsii.String(k), jsii.String(value))
	}
	return configMap
}

func loadEnv(monitoring, env map[string]string) map[string]string {
	if env == nil {
		env = make(map[string]string)
	}
	for key, value := range monitoring {
		if _, ok := env[key]; ok {
			env[key] = monitoring[key]
		} else {
			env[key] = value
		}
	}
	return env
}

func loadMonitoringEnvWithAliases(values, aliases map[string]string, monitoringEnv []string) map[string]string {
	result := make(map[string]string)
	for k, v := range values {
		//If the monitoring slice from the template has a standard key,
		//the key is set to the cfg value in the result
		if slices.Contains(monitoringEnv, k) {
			result[k] = v
			//if there is an alias, the alias key's value is set
			//to the matching value from the cfg
		} else if configValue, ok := aliases[k]; ok {
			result[configValue] = v
		}

	}
	return result
}
