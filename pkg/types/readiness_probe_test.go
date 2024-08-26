package types

import (
	compose "github.com/compose-spec/compose-go/v2/types"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func TestParseHealthCheck(t *testing.T) {
	timeout := compose.Duration(10 * time.Second)
	interval := compose.Duration(30 * time.Second)
	retries := uint64(5)
	startPeriod := compose.Duration(1 * time.Minute)
	startInterval := compose.Duration(10 * time.Second)

	var validHealthCheck = compose.HealthCheckConfig{
		Test: compose.HealthCheckTest{
			"CMD",
			"wget",
			"-o",
			"/dev/null",
			"-q",
			"http://ai-service:5001/health",
		},
		Timeout:       &timeout,
		Interval:      &interval,
		Retries:       &retries,
		StartPeriod:   &startPeriod,
		StartInterval: &startInterval,
		Disable:       false,
		Extensions:    nil,
	}

	var invalidHealthCheck = compose.HealthCheckConfig{
		Test:          compose.HealthCheckTest{},
		Timeout:       &timeout,
		Interval:      &interval,
		Retries:       &retries,
		StartPeriod:   &startPeriod,
		StartInterval: &startInterval,
		Disable:       false,
		Extensions:    nil,
	}

	cmd1 := "CMD"
	cmd2 := "wget"
	cmd3 := "-o"
	cmd4 := "/dev/null"
	cmd5 := "-q"
	cmd6 := "http://ai-service:5001/health"
	expectedDelay := "1m0s"
	expectedTimeout := "10s"
	expectedInterval := "30s"
	expectedRetries := float64(5)

	var expected = &ReadinessProbe{
		Commands: []*string{
			&cmd1,
			&cmd2,
			&cmd3,
			&cmd4,
			&cmd5,
			&cmd6,
		},
		Delay:    &expectedDelay,
		Timeout:  &expectedTimeout,
		Interval: &expectedInterval,
		Retries:  &expectedRetries,
	}

	t.Run("test valid health check", func(t *testing.T) {
		probe, err := ParseHealthCheck(&validHealthCheck)
		assert.NoError(t, err)
		assertProbeEquals(t, probe, expected)

	})

	t.Run("test invalid health check", func(t *testing.T) {
		_, err := ParseHealthCheck(&invalidHealthCheck)
		assert.Error(t, err)
	})
}

func TestConvertReadinessProbe(t *testing.T) {
	t.Run("test valid readiness probe", func(t *testing.T) {

	})
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func assertError(t testing.TB, err error) {
	t.Helper()
	if err == nil {
		t.Fatalf("expected error, got none")
	}
}

func assertProbeEquals(t testing.TB, got, want *ReadinessProbe) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
