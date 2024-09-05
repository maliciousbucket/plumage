package resilience

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"os"
	"reflect"
	"testing"
)

func TestValidRateLimitConfig_UnmarshalYAML(t *testing.T) {
	tests := []struct {
		name     string
		file     string
		expected *RateLimitConfig
	}{
		{
			name: "Valid IpDepth Rate Limit Config",
			file: "../../testdata/template/retry-strategy/valid-ipDepth-ratelimit-strategy-config.yaml",
			expected: &RateLimitConfig{
				Average:  10,
				Burst:    5,
				Period:   "10s",
				Strategy: &IpDepthStrategy{Depth: 3},
			},
		},

		{
			name: "Multiple LimitStrategy config",
			file: "../../testdata/template/retry-strategy/multiple-ratelimit-strategy-config.yaml",
			expected: &RateLimitConfig{
				Average: 10,
				Burst:   5,
				Period:  "10s",
				Strategy: &IpDepthStrategy{
					Depth: 3,
				},
			},
		},
		{
			name: "Valid Host LimitStrategy config",
			file: "../../testdata/template/retry-strategy/valid-host-strategy-config.yaml",
			expected: &RateLimitConfig{
				Average: 10,
				Burst:   5,
				Period:  "10s",
				Strategy: &HostStrategy{
					Host: true,
				},
			},
		},
		{
			name: "Valid Header LimitStrategy config",
			file: "../../testdata/template/retry-strategy/valid-header-strategy-config.yaml",
			expected: &RateLimitConfig{
				Average:  10,
				Burst:    5,
				Period:   "10s",
				Strategy: &RequestHeaderStrategy{HeaderName: "host"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := os.ReadFile(tt.file)
			if err != nil {
				t.Fatalf("failed to read template file: %v", err)
			}
			var config RateLimitConfig
			if err := yaml.Unmarshal(data, &config); err != nil {
				t.Fatalf("failed to unmarshal YAML: %v", err)
			}
			assertRateLimitStrategyEqual(t, tt.expected, &config)
		})
	}
}

func TestInvalidRateLimitConfig_UnmarshalYAML(t *testing.T) {
	tests := []struct {
		name string
		file string
	}{
		{
			name: "Invalid Header LimitStrategy config",
			file: "../../testdata/template/retry-strategy/invalid-header-strategy-config.yaml",
		},
		{
			name: "Invalid Host LimitStrategy config",
			file: "../../testdata/template/retry-strategy/invalid-host-strategy-config.yaml",
		},
		{
			name: "Invalid IpDepth LimitStrategy config",
			file: "../../testdata/template/retry-strategy/invalid-ipdepth-strategy-config.yaml",
		},
		{
			name: "Unknown strategy",
			file: "../../testdata/template/retry-strategy/unknown-ratelimit-strategy-config.yaml",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := os.ReadFile(tt.file)
			if err != nil {
				t.Fatalf("failed to read template file: %v", err)
			}
			var config RateLimitConfig
			err = yaml.Unmarshal(data, &config)
			assert.Error(t, err)
		})
	}
}

func assertRateLimitStrategyEqual(t testing.TB, expected, actual *RateLimitConfig) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected:\n%v\nActual:\n%v", expected, actual)
	}
}

//{
//	name: "Invalid Retry Config",
//	file: "../../testdata/template/retry-strategy/unknown-ratelimit-strategy-config.yaml",
//	expected: &RateLimitConfig{
//		AverageRequests:  10,
//		BurstRequests:    5,
//		RatePeriod:   "10s",
//		LimitStrategy: nil,
//	},
//},
