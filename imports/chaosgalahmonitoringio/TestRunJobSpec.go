package chaosgalahmonitoringio

// TestRunJobSpec defines the desired state of TestRunJob.
type TestRunJobSpec struct {
	Args *string `field:"required" json:"args" yaml:"args"`
	// // Specifies the job that will be created when executing a CronJob.
	//
	// JobTemplate batchv1.JobTemplateSpec `json:"jobTemplate"`
	JobTemplate    *TestRunJobSpecJobTemplate `field:"required" json:"jobTemplate" yaml:"jobTemplate"`
	Schedule       *TestRunJobSpecSchedule    `field:"required" json:"schedule" yaml:"schedule"`
	ServiceAccount *string                    `field:"required" json:"serviceAccount" yaml:"serviceAccount"`
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run "make" to regenerate code after modifying this file.
	TestName                   *string               `field:"required" json:"testName" yaml:"testName"`
	Env                        *[]*TestRunJobSpecEnv `field:"optional" json:"env" yaml:"env"`
	EnvConfigMap               *string               `field:"optional" json:"envConfigMap" yaml:"envConfigMap"`
	FailedJobsHistoryLimit     *float64              `field:"optional" json:"failedJobsHistoryLimit" yaml:"failedJobsHistoryLimit"`
	Image                      *string               `field:"optional" json:"image" yaml:"image"`
	RunOnce                    *bool                 `field:"optional" json:"runOnce" yaml:"runOnce"`
	ScriptConfigMap            *string               `field:"optional" json:"scriptConfigMap" yaml:"scriptConfigMap"`
	StartingDeadlineSeconds    *float64              `field:"optional" json:"startingDeadlineSeconds" yaml:"startingDeadlineSeconds"`
	SuccessfulJobsHistoryLimit *float64              `field:"optional" json:"successfulJobsHistoryLimit" yaml:"successfulJobsHistoryLimit"`
	Suspend                    *bool                 `field:"optional" json:"suspend" yaml:"suspend"`
	TestRunJobHistoryLimit     *float64              `field:"optional" json:"testRunJobHistoryLimit" yaml:"testRunJobHistoryLimit"`
}
