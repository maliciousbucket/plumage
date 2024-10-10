package chaos

import (
	"errors"
	"fmt"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	kplus "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2"
	chaos "github.com/maliciousbucket/plumage/imports/chaosgalahmonitoringio"
	"strings"
)

var (
	defaultImage        = "docker.io/maliciousbucket/k6:latest"
	otelOutput          = "experimental-opentelemetry"
	otelGrpcEndpoint    = "K6_OTEL_GRPC_EXPORTER_ENDPOINT"
	defaultMetricPrefix = "k6_"
	k6OtelInsecure      = "K6_OTEL_GRPC_EXPORTER_INSECURE=true"
	k6MetricPrefixArg   = "K6_OTEL_METRIC_PREFIX"
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
	Args             string
	Env              map[string]string
	RunOnce          bool
	Schedule         *JobSchedule
	Labels           map[string]string
	Annotations      map[string]string
	ExistingEnv      *ExistingResource
	envMap           string
	ExistingScript   *ExistingResource
	scriptMap        string
	ExistingAccount  *ExistingResource
	account          string
}

type ExistingResource struct {
	Name string
}

func NewTestRunFromTemplate(scope constructs.Construct, id string, ns, alloyAddr string, template *Template, script *ScriptTemplate) (constructs.Construct, error) {
	chartProps := &cdk8s.ChartProps{
		DisableResourceNameHashes: jsii.Bool(true),
	}
	if ns != "" {
		chartProps.Namespace = jsii.String(ns)
	}

	ct := cdk8s.NewChart(scope, jsii.String(id), chartProps)

	props := &TestRunProps{
		Namespace:        template.Namespace,
		ScriptDir:        template.ScriptDir,
		LibDir:           template.LibDir,
		Name:             script.Name,
		ScriptName:       script.ScriptName,
		LibFiles:         script.LibFiles,
		Env:              script.Env,
		ResourceRequests: script.ResourceRequests,
		ResourceLimits:   script.ResourceLimits,
		RunOnce:          script.RunOnce,
		Labels:           script.Labels,
		Annotations:      script.Annotations,
	}
	opts := []TestRunOpt{
		WithOtelOutput(alloyAddr, defaultMetricPrefix),
	}

	if template.ServiceAccount != "" {
		account := &ExistingResource{
			Name: template.ServiceAccount,
		}
		script.ExistingAccount = account.Name
	}

	if script.ExistingAccount != "" {
		opts = append(opts, WithExistingAccount(script.ExistingAccount))
	} else {
		opts = append(opts, WithNewAccount(ct))
	}

	if script.ExistingScript != "" {
		opts = append(opts, WithExistingScriptMap(script.ExistingScript))
	} else {
		opts = append(opts, WithScript(ct, template.ScriptDir, template.LibDir, script.LibFiles))
	}

	if script.ExistingEnv != "" {
		opts = append(opts, WithExistingEnvMap(script.ExistingEnv))
	}

	if script.Env != nil {
		opts = append(opts, WithNewEnv(ct))
	}

	if script.Schedule != nil {
		opts = append(opts, WithSchedule(script.Schedule))
	} else {
		opts = append(opts, WithDefaultSchedule())
	}

	if script.Args != nil {
		opts = append(opts, WithArgs(script.Args))
	}

	testId := fmt.Sprintf("%s-%s", script.Name, "testrun")

	return newTestRun(ct, testId, props, opts...)

}

func newTestRun(scope constructs.Construct, id string, props *TestRunProps, opts ...TestRunOpt) (constructs.Construct, error) {
	if len(opts) == 0 {
		return nil, fmt.Errorf("no test run options provided")
	}
	var optsErr error
	for _, opt := range opts {
		err := opt(props)
		if err != nil {
			optsErr = errors.Join(optsErr, err)
		}
	}
	if optsErr != nil {
		return nil, optsErr
	}
	var env []*chaos.TestRunJobSpecEnv
	if props.Env != nil {
		env = envToK8s(props.Env)
	}

	if props.ExistingAccount != nil {
		account := props.ExistingAccount.Name
		props.account = account
	}

	if props.ExistingScript != nil {
		props.scriptMap = props.ExistingScript.Name
	}

	if props.ExistingEnv != nil {
		props.envMap = props.ExistingEnv.Name
	}

	jobTemplate := newJobTemplate(props.Labels, props.Annotations)

	job := chaos.NewTestRunJob(scope, jsii.String(id), &chaos.TestRunJobProps{
		Metadata: &cdk8s.ApiObjectMetadata{},
		Spec: &chaos.TestRunJobSpec{
			Args:        jsii.String(props.Args),
			JobTemplate: jobTemplate,
			Schedule: &chaos.TestRunJobSpecSchedule{
				DayOfMonth: jsii.String(props.Schedule.DayOfMonth),
				DayOfWeek:  jsii.String(props.Schedule.DayOfWeek),
				Hour:       jsii.String(props.Schedule.Hour),
				Minute:     jsii.String(props.Schedule.Minute),
				Month:      jsii.String(props.Schedule.Month),
			},
			ServiceAccount:             jsii.String(props.account),
			TestName:                   jsii.String(props.Name),
			Env:                        &env,
			EnvConfigMap:               jsii.String(props.envMap),
			FailedJobsHistoryLimit:     jsii.Number(3),
			Image:                      jsii.String(defaultImage),
			RunOnce:                    jsii.Bool(props.RunOnce),
			ScriptConfigMap:            jsii.String(props.scriptMap),
			StartingDeadlineSeconds:    jsii.Number(30),
			SuccessfulJobsHistoryLimit: jsii.Number(3),
			Suspend:                    jsii.Bool(false),
			TestRunJobHistoryLimit:     jsii.Number(5),
		},
	})
	return job, nil
}

func (t *TestRunProps) setNewJobServiceAccount(account kplus.ServiceAccount) error {

	if account == nil {
		return fmt.Errorf("account is nil")
	}

	t.account = *account.Metadata().Name()
	return nil

}

type TestRunOpt func(t *TestRunProps) error

func WithExistingScriptMap(configMap string) TestRunOpt {
	return func(t *TestRunProps) error {
		if configMap == "" {
			return fmt.Errorf("empty config map name")
		}
		resource := &ExistingResource{Name: configMap}
		t.ExistingScript = &ExistingResource{Name: resource.Name}
		return nil
	}
}

func WithExistingEnvMap(envConfigMap string) TestRunOpt {
	return func(t *TestRunProps) error {
		if envConfigMap == "" {
			return fmt.Errorf("envConfigMap is empty")
		}
		resource := &ExistingResource{Name: envConfigMap}
		t.ExistingScript = resource
		return nil
	}
}

func WithNewEnv(scope constructs.Construct) TestRunOpt {
	return func(t *TestRunProps) error {
		if t.Env == nil || len(t.Env) == 0 {
			return nil
		}
		if t.ExistingEnv == nil {
			envMap := newEnvConfigMap(scope, t.Name, t.Namespace, t.Env)
			resource := &ExistingResource{Name: *envMap.Metadata().Name()}
			t.envMap = resource.Name
			t.Env = nil
			return nil
		}
		return nil
	}
}

func WithExistingAccount(account string) TestRunOpt {
	return func(t *TestRunProps) error {
		if account == "" {
			return fmt.Errorf("account name is empty")
		}
		resource := &ExistingResource{Name: account}
		t.ExistingAccount = &ExistingResource{Name: resource.Name}
		return nil
	}
}

func WithScript(scope constructs.Construct, scriptDir, libDir string, libFiles []string) TestRunOpt {
	return func(t *TestRunProps) error {
		if scriptDir == "" {
			return fmt.Errorf("must specify script directory")
		}
		if t.Name == "" {
			return fmt.Errorf("empty test name provided")
		}
		if t.ExistingScript != nil {
			return fmt.Errorf("script already exists: %s", t.ExistingScript.Name)
		}
		scriptMap := newTestFileConfigMap(scope, t.Namespace, t.Name, scriptDir, libDir, t.Name, libFiles)
		t.scriptMap = *scriptMap.Metadata().Name()

		return nil
	}
}

func WithNewAccount(scope constructs.Construct) TestRunOpt {
	return func(t *TestRunProps) error {
		if t.ExistingAccount != nil {
			return fmt.Errorf("a service account has already been specified")
		}
		_, account := NewTestRunRBAC(scope, t.Namespace)
		t.account = account
		return nil
	}
}

func WithSchedule(schedule *JobSchedule) TestRunOpt {
	return func(t *TestRunProps) error {
		if schedule == nil {
			return fmt.Errorf("schedule is nil")
		}
		t.Schedule = schedule
		return nil
	}
}

func WithScheduleFields(min, hour, dom, mon, dow string) TestRunOpt {
	return func(t *TestRunProps) error {
		if min == "" && hour == "" && dom == "" && dow == "" {
			return fmt.Errorf("must specify at least one field of the schedule")
		}
		schedule := &JobSchedule{
			Minute:     "*",
			Hour:       "*",
			DayOfMonth: "*",
			Month:      "*",
			DayOfWeek:  "*",
		}
		if min != "" {
			schedule.Minute = min
		}
		if hour != "" {
			schedule.Hour = hour
		}
		if dom != "" {
			schedule.DayOfMonth = dom
		}
		if mon != "" {
			schedule.Month = mon
		}
		if dow != "" {
			schedule.DayOfWeek = dow
		}
		t.Schedule = schedule
		return nil
	}
}

func WithDefaultSchedule() TestRunOpt {
	return func(t *TestRunProps) error {
		if t.Schedule != nil {
			return fmt.Errorf("default schedule: a schedule has already been specified")
		}
		schedule := &JobSchedule{
			Minute:     "*/5",
			Hour:       "*",
			DayOfMonth: "*",
			Month:      "*",
			DayOfWeek:  "*",
		}
		t.Schedule = schedule
		return nil
	}
}

func RunOnce() TestRunOpt {
	return func(t *TestRunProps) error {
		t.RunOnce = true
		return nil
	}
}

func WithArgs(args []string) TestRunOpt {
	return func(t *TestRunProps) error {
		if args == nil {
			return fmt.Errorf("args is nil")
		}
		var result string
		for _, arg := range args {
			result = strings.Join(strings.Fields(arg), " ")
		}
		result = strings.TrimSpace(result)
		t.Args = result
		return nil
	}
}

func WithOutput(output string) TestRunOpt {
	return func(t *TestRunProps) error {
		if output == "" {
			return nil
		}
		out := fmt.Sprintf("-o %s", output)
		result := strings.Join([]string{t.Args, out}, " ")
		t.Args = strings.TrimSpace(result)
		return nil
	}
}

func WithOtelOutput(addr string, prefix string) TestRunOpt {
	return func(t *TestRunProps) error {
		if addr == "" {
			return fmt.Errorf("output address is empty")
		}
		addOtel := true
		for _, arg := range strings.Split(t.Args, "") {
			if strings.Contains(arg, "-o") {
			}
			addOtel = false
		}
		if addOtel {
			metricPrefix := defaultMetricPrefix
			if prefix != "" {
				metricPrefix = prefix
			}
			result := t.Args
			result = strings.Join([]string{result, k6OtelInsecure}, " ")
			prefixStr := fmt.Sprintf("%s=%s", k6MetricPrefixArg, metricPrefix)
			exportStr := fmt.Sprintf("%s=%s", otelGrpcEndpoint, addr)
			outputStr := fmt.Sprintf("-o %s", otelOutput)
			result = strings.Join([]string{result, prefixStr, exportStr, outputStr}, " ")
			t.Args = strings.TrimSpace(result)
			return nil
		}
		if prefix != "" {
			prefixStr := fmt.Sprintf("%s=%s", k6MetricPrefixArg, prefix)
			result := strings.Join([]string{t.Args, prefixStr}, " ")
			t.Args = strings.TrimSpace(result)
			return nil
		}
		return nil
	}
}

func newJobTemplate(labels, annotations map[string]string) *chaos.TestRunJobSpecJobTemplate {
	jobLabels := stringMapToK8s(labels)
	jobAnnotations := stringMapToK8s(annotations)
	return &chaos.TestRunJobSpecJobTemplate{
		Labels:      &jobLabels,
		Annotations: &jobAnnotations,
	}
}

func stringMapToK8s(source map[string]string) map[string]*string {
	result := map[string]*string{}
	for k, v := range source {
		result[k] = &v
	}
	return result
}

func envToK8s(env map[string]string) []*chaos.TestRunJobSpecEnv {
	if env == nil {
		return nil
	}
	var result []*chaos.TestRunJobSpecEnv
	for k, v := range env {
		result = append(result, &chaos.TestRunJobSpecEnv{
			Name:  jsii.String(k),
			Value: jsii.String(v),
		})
	}
	return result
}
