package chaosgalahmonitoringio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// TestRunJob is the Schema for the testrunjobs API.
type TestRunJobProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// TestRunJobSpec defines the desired state of TestRunJob.
	Spec *TestRunJobSpec `field:"optional" json:"spec" yaml:"spec"`
}
