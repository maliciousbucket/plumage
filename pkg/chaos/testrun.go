package chaos

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	chaos "github.com/maliciousbucket/plumage/imports/chaosgalahmonitoringio"
)

var (
	defaultImage = "docker.io/maliciousbucket/k6:latest"
)

type TestRunProps struct {
	Namespace        string
	ScriptDir        string
	LibDir           string
	Name             string
	ScriptName       string
	LibFiles         []string
	ResourceRequests ScriptResources
	ResourceLimits   ScriptResources
	Args             []string
	Env              map[string]string
	RunOnce          bool
	Schedule         string
	Labels           map[string]string
	Annotations      map[string]string
	ExistingEnv      *ExistingResource
	ExistingScript   *ExistingResource
	ExistingAccount  *ExistingResource
}

type ExistingResource struct {
	Name string
}

func NewTestRun(scope constructs.Construct, id string, template *Template, script *ScriptTemplate) constructs.Construct {
	chaos.NewTestRunJob(scope, jsii.String(id), &chaos.TestRunJobProps{
		Metadata: nil,
		Spec: &chaos.TestRunJobSpec{
			Args:                       nil,
			JobTemplate:                nil,
			Schedule:                   nil,
			ServiceAccount:             nil,
			TestName:                   nil,
			Env:                        nil,
			EnvConfigMap:               nil,
			FailedJobsHistoryLimit:     nil,
			RunOnce:                    nil,
			ScriptConfigMap:            nil,
			StartingDeadlineSeconds:    nil,
			SuccessfulJobsHistoryLimit: nil,
			Suspend:                    nil,
			TestRunJobHistoryLimit:     nil,
			Image:                      jsii.String(defaultImage),
		},
	})
	return nil
}
