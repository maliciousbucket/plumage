package resilience

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	kplus "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2"
	"time"
)

const (
	balancedADurationSeconds = 90
	balancedBDurationSeconds = 180

	conservativeADurationSeconds = 120
	conservativeBDurationSeconds = 210

	AggressiveADurationSeconds = 30
	AggressiveBDurationSeconds = 60
)

func AutoScalingUpRules(p scalingPolicy, s stabilisationPolicy, maxReplicas int) *kplus.ScalingRules {
	var rules kplus.ScalingRules
	duration := StabilisationDuration(s)

	switch p {
	case scalingPolicyBalanced:
		policyA := kplus.ScalingPolicy{
			Replicas: kplus.Replicas_Percent(jsii.Number(75)),
			Duration: cdk8s.Duration_Seconds(jsii.Number(balancedADurationSeconds)),
		}
		policyB := kplus.ScalingPolicy{
			Replicas: kplus.Replicas_Absolute(jsii.Number(maxReplicas)),
			Duration: cdk8s.Duration_Seconds(jsii.Number(balancedBDurationSeconds)),
		}
		rules.Policies = &[]*kplus.ScalingPolicy{&policyA, &policyB}
		rules.StabilizationWindow = cdk8s.Duration_Parse(jsii.String(duration.String()))
		rules.Strategy = kplus.ScalingStrategy_MIN_CHANGE

	case scalingPolicyConservative:
		policyA := kplus.ScalingPolicy{
			Replicas: kplus.Replicas_Percent(jsii.Number(65)),
			Duration: cdk8s.Duration_Seconds(jsii.Number(conservativeADurationSeconds)),
		}
		policyB := kplus.ScalingPolicy{
			Replicas: kplus.Replicas_Percent(jsii.Number(85)),
			Duration: cdk8s.Duration_Seconds(jsii.Number(conservativeBDurationSeconds)),
		}
		rules.Policies = &[]*kplus.ScalingPolicy{&policyA, &policyB}
		rules.StabilizationWindow = cdk8s.Duration_Parse(jsii.String(duration.String()))
		rules.Strategy = kplus.ScalingStrategy_MIN_CHANGE

	case scalingPolicyAggressive:
		policyA := kplus.ScalingPolicy{
			Replicas: kplus.Replicas_Percent(jsii.Number(90)),
			Duration: cdk8s.Duration_Seconds(jsii.Number(AggressiveADurationSeconds)),
		}

		policyB := kplus.ScalingPolicy{
			Replicas: kplus.Replicas_Absolute(jsii.Number(maxReplicas)),
			Duration: cdk8s.Duration_Seconds(jsii.Number(AggressiveBDurationSeconds)),
		}
		rules.Policies = &[]*kplus.ScalingPolicy{&policyA, &policyB}
		rules.StabilizationWindow = cdk8s.Duration_Parse(jsii.String(duration.String()))
		rules.Strategy = kplus.ScalingStrategy_MAX_CHANGE

	case scalingPolicyNone:

	}
	return &rules
}

func AutoScalingDownRules(p scalingPolicy, s stabilisationPolicy, minReplicas int) *kplus.ScalingRules {
	rules := kplus.ScalingRules{}
	duration := StabilisationDuration(s)

	switch p {
	case scalingPolicyBalanced:
		policyA := kplus.ScalingPolicy{
			Replicas: kplus.Replicas_Percent(jsii.Number(50)),
			Duration: cdk8s.Duration_Seconds(jsii.Number(balancedADurationSeconds)),
		}
		policyB := kplus.ScalingPolicy{
			Replicas: kplus.Replicas_Percent(jsii.Number(20)),
			Duration: cdk8s.Duration_Seconds(jsii.Number(balancedBDurationSeconds)),
		}
		rules.Policies = &[]*kplus.ScalingPolicy{&policyA, &policyB}
		rules.StabilizationWindow = cdk8s.Duration_Parse(jsii.String(duration.String()))
		rules.Strategy = kplus.ScalingStrategy_MIN_CHANGE
	case scalingPolicyConservative:
		policyA := kplus.ScalingPolicy{
			Replicas: kplus.Replicas_Percent(jsii.Number(30)),
			Duration: cdk8s.Duration_Seconds(jsii.Number(conservativeADurationSeconds)),
		}

		policyB := kplus.ScalingPolicy{
			Replicas: kplus.Replicas_Absolute(jsii.Number(minReplicas)),
			Duration: cdk8s.Duration_Seconds(jsii.Number(conservativeBDurationSeconds)),
		}
		rules.Policies = &[]*kplus.ScalingPolicy{&policyA, &policyB}
		rules.StabilizationWindow = cdk8s.Duration_Parse(jsii.String(duration.String()))
		rules.Strategy = kplus.ScalingStrategy_MIN_CHANGE
	case scalingPolicyAggressive:
		policyA := kplus.ScalingPolicy{
			Replicas: kplus.Replicas_Percent(jsii.Number(50)),
			Duration: cdk8s.Duration_Seconds(jsii.Number(AggressiveADurationSeconds)),
		}
		policyB := kplus.ScalingPolicy{
			Replicas: kplus.Replicas_Percent(jsii.Number(30)),
			Duration: cdk8s.Duration_Seconds(jsii.Number(AggressiveBDurationSeconds)),
		}
		rules.Policies = &[]*kplus.ScalingPolicy{&policyA, &policyB}
		rules.StabilizationWindow = cdk8s.Duration_Parse(jsii.String(duration.String()))
		rules.Strategy = kplus.ScalingStrategy_MAX_CHANGE
	case scalingPolicyNone:
	}
	return &rules
}

func AutoScalerMetrics(r []resource, p resourcePolicy) []*kplus.Metric {
	return nil
}

func StabilisationDuration(s stabilisationPolicy) *time.Duration {
	var duration time.Duration
	switch s {
	case stabilisationPolicyBalanced:
		duration = time.Second * time.Duration(120)

	case stabilisationPolicyAggressive:
		duration = time.Second * time.Duration(60)

	case stabilisationPolicyConservative:
		duration = time.Second * time.Duration(300)

	case stabilisationPolicyNone:
		duration = time.Second * time.Duration(0)
	}
	return &duration

}
