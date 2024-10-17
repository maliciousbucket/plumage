package chaosgalahmonitoringio


// // Specifies the job that will be created when executing a CronJob.
//
// JobTemplate batchv1.JobTemplateSpec `json:"jobTemplate"`
type TestRunJobSpecJobTemplate struct {
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
}

