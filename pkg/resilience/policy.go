package resilience

import (
	"fmt"
	"github.com/maliciousbucket/plumage/pkg/template"
	"slices"
	"strings"
)

const (
	resourceCpu    = resource("cpu")
	resourceMemory = resource("memory")

	scalingPolicyAggressive   = scalingPolicy("aggressive")
	scalingPolicyBalanced     = scalingPolicy("balanced")
	scalingPolicyConservative = scalingPolicy("conservative")
	scalingPolicyNone         = scalingPolicy("none")

	stabilisationPolicyAggressive   = stabilisationPolicy("aggressive")
	stabilisationPolicyBalanced     = stabilisationPolicy("balanced")
	stabilisationPolicyConservative = stabilisationPolicy("conservative")
	stabilisationPolicyNone         = stabilisationPolicy("none")

	resourcePolicyAggressive   = resourcePolicy("aggressive")
	resourcePolicyBalanced     = resourcePolicy("balanced")
	resourcePolicyConservative = resourcePolicy("conservative")
)

type resource string

type stabilisationPolicy string

type resourcePolicy string

type scalingPolicy string

type Policy struct {
	MinReplicas     int
	MaxReplicas     int
	Resources       []resource
	ResourcePolicy  *resourcePolicy
	ScaleUpPolicy   *scalingPolicy
	ScaleDownPolicy *scalingPolicy
}

func ParseScalingTemplate(a *template.AutoScalingTemplate) (*Policy, error) {

	if len(a.Policy.ScaleOn) == 0 {
		return nil, fmt.Errorf("must select at least one resource to scale on")
	}

	var resources []resource
	for _, res := range a.Policy.ScaleOn {
		lowRes := strings.ToLower(res)
		var policyRes resource
		switch lowRes {
		case "cpu":

			policyRes = resourceCpu
			if !slices.Contains(resources, policyRes) {
				resources = append(resources, policyRes)
			}
		case "memory":
			policyRes = resourceMemory
			if !slices.Contains(resources, policyRes) {
				resources = append(resources, policyRes)
			}
		default:
			return nil, fmt.Errorf("unknown resource %s", lowRes)

		}
	}

	var resPolicy resourcePolicy
	if a.Policy.ResourcePolicy == "" {
		resPolicy = resourcePolicyBalanced
	}
	lowResPolicy := strings.ToLower(a.Policy.ResourcePolicy)
	switch lowResPolicy {
	case "aggressive":
		resPolicy = resourcePolicyAggressive
	case "balanced":
		resPolicy = resourcePolicyBalanced
	case "conservative":
		resPolicy = resourcePolicyConservative
	default:
		return nil, fmt.Errorf("unknown resource policy %s", a.Policy.ResourcePolicy)
	}

	var scaleUpPolicy scalingPolicy
	if a.Policy.ScaleUpPolicy == "" {
		scaleUpPolicy = scalingPolicyNone
	}

	lowScaleUp := strings.ToLower(a.Policy.ScaleUpPolicy)
	switch lowScaleUp {
	case "none":
		scaleUpPolicy = scalingPolicyNone
	case "conservative":
		scaleUpPolicy = scalingPolicyConservative
	case "aggressive":
		scaleUpPolicy = scalingPolicyAggressive
	case "balanced":
		scaleUpPolicy = scalingPolicyBalanced
	default:
		return nil, fmt.Errorf("unknown policy scale_up %s", a.Policy.ScaleUpPolicy)

	}

	var scaleDownPolicy scalingPolicy
	if a.Policy.ScaleDownPolicy == "" {
		scaleDownPolicy = scalingPolicyNone
	}

	lowScaleDown := strings.ToLower(a.Policy.ScaleDownPolicy)
	switch lowScaleDown {
	case "none":
		scaleDownPolicy = scalingPolicyNone
	case "conservative":
		scaleDownPolicy = scalingPolicyConservative
	case "aggressive":
		scaleDownPolicy = scalingPolicyAggressive
	case "balanced":
		scaleDownPolicy = scalingPolicyBalanced
	default:
		return nil, fmt.Errorf("unknown policy scale_down %s", a.Policy.ScaleDownPolicy)
	}

	var minReplicas = 1
	var maxReplicas = 2

	if a.MinReplicas != 0 && a.MaxReplicas != 0 {
		if a.MinReplicas > a.MaxReplicas {
			return nil, fmt.Errorf("minimum number of replicas exceeds maximum number of replicas. Min: %d, Max: %d", a.MinReplicas, a.MaxReplicas)
		}
	}

	if a.MinReplicas != 0 {
		if a.MaxReplicas == 0 {
			maxReplicas = a.MinReplicas + 1
		}
		minReplicas = a.MinReplicas
	}

	return &Policy{
		MinReplicas:     minReplicas,
		MaxReplicas:     maxReplicas,
		Resources:       resources,
		ResourcePolicy:  &resPolicy,
		ScaleUpPolicy:   &scaleUpPolicy,
		ScaleDownPolicy: &scaleDownPolicy,
	}, nil

}

type CircuitBreakerPolicy struct{}

type RetryPolicy struct {
}

type RateLimitPolicy struct{}
