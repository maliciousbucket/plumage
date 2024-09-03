package kubernetes

import (
	"fmt"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"testing"
)

type MockRetrySpec struct {
	Attempts          int
	InitialIntervalMS int
}

func (m *MockRetrySpec) RetryAttempts() int {
	return m.Attempts
}

func (m *MockRetrySpec) IntervalMS() string {
	return fmt.Sprintf("%dms", m.InitialIntervalMS)
}

func TestMiddlewareRetrySpec(t *testing.T) {
	//app := cdk8s.Testing_App(nil)
	chart := cdk8s.Testing_Chart()
	//
	//	spec := MockRetrySpec{
	//		Attempts:          1,
	//		InitialIntervalMS: 500,
	//	}
	//
	//	NewRetryMiddleware(chart, "a-service", &spec)
	cdk8s.Testing_Synth(chart)
}
