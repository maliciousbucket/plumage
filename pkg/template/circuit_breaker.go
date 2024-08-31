package template

import (
	"fmt"
)

const (
	ExpressionLatency           = "LatencyAtQuantileMS"
	ExpressionNetworkErrRatio   = "NetworkErrorRatio()"
	ExpressionResponseCodeRatio = "ResponseCodeRatio"
)

type CircuitBreakerConfig struct {
	Expressions             []ExpressionList `yaml:"expressionGroups"`
	CircuitCheckPeriod      string           `yaml:"checkPeriod"`
	CircuitFallbackDuration string           `yaml:"fallbackDuration"`
	CircuitRecoveryDuration string           `yaml:"recoveryDuration"`
}

func (c *CircuitBreakerConfig) Validate() error {
	return nil
}

func (c *CircuitBreakerConfig) CircuitBreakerExpression() string {
	var expressions []string
	for _, expr := range c.Expressions {
		var innerExpressions string
		ExpressionWithAndArgs(expr.Expressions)(innerExpressions)
		expressions = append(expressions, innerExpressions)

	}
	var circuitBreakerExpression string
	if len(expressions) > 0 {
		circuitBreakerExpression = expressions[0]
		for _, expr := range expressions[1:] {
			circuitBreakerExpression = fmt.Sprintf("%s || %s", circuitBreakerExpression, expr)
		}
	}
	return circuitBreakerExpression
}

func (c *CircuitBreakerConfig) CheckPeriod() string {
	return c.CircuitCheckPeriod
}

func (c *CircuitBreakerConfig) FallbackDuration() string {
	return c.CircuitFallbackDuration
}

func (c *CircuitBreakerConfig) RecoveryDuration() string {
	return c.CircuitRecoveryDuration
}

type CircuitBreakerExpression interface {
	Expression() string
	Validate() error
}

type ExpressionList struct {
	Expressions []CircuitBreakerExpression `yaml:"expressions"`
}

type Latency struct {
	Quantile  float64 `yaml:"quantile"`
	Operator  string  `yaml:"operator"`
	Parameter float64 `yaml:"parameter"`
}

func (l *Latency) Expression() string {
	exp := fmt.Sprintf("%s(%f) %s %f", ExpressionLatency, l.Quantile, l.Operator, l.Parameter)
	return exp
}

func (l *Latency) Validate() error {
	if l.Quantile > 100 {
		return fmt.Errorf("quantile %f cannot be greater than 100", l.Quantile)
	}

	if l.Quantile < 0 {
		return fmt.Errorf("quantile %f cannot be less than 0", l.Quantile)
	}

	opErr := validateOperator(l.Operator)
	if opErr != nil {
		return opErr
	}

	if l.Parameter < 0 {
		return fmt.Errorf("parameter %f cannot be less than 0", l.Parameter)
	}
	return nil
}

type ResponseCodeRatio struct {
	From       int     `yaml:"from"`
	To         int     `yaml:"to"`
	DivideFrom int     `yaml:"divideFrom"`
	DivideTo   int     `yaml:"divideTo"`
	Operator   string  `yaml:"operator"`
	Parameter  float64 `yaml:"parameter"`
}

func (r *ResponseCodeRatio) Expression() string {
	exp := fmt.Sprintf("%s(%d, %d, %d, %d) %s %f", ExpressionResponseCodeRatio,
		r.From, r.To, r.DivideFrom, r.DivideTo, r.Operator, r.Parameter)
	return exp
}

func (r *ResponseCodeRatio) Validate() error {
	if r.From > r.To {
		return fmt.Errorf("from value must be less than to value. From: %d, To: %d", r.From, r.To)
	}

	if r.To < 100 {
		return fmt.Errorf("to value must be greater than 100. To: %d", r.To)
	}

	if r.To > 600 {
		return fmt.Errorf("to value must be less than or equal to 600. To: %d", r.To)
	}

	if r.DivideFrom > r.DivideTo {
		return fmt.Errorf("divideFrom value must be less than divideTo. From: %d, To: %d", r.From, r.To)
	}

	if r.DivideTo < 100 {
		return fmt.Errorf("divideTo value must be greater than 100. To: %d", r.DivideTo)
	}

	if r.DivideTo > 600 {
		return fmt.Errorf("divideTo value must be less than or equal to 600. To: %d", r.DivideTo)
	}

	if r.Parameter > 1 {
		return fmt.Errorf("parameter %f cannot be greater than 1", r.Parameter)
	}

	if r.Parameter < 0 {
		return fmt.Errorf("parameter %f cannot be less than 0", r.Parameter)
	}

	oppErr := validateOperator(r.Operator)
	if oppErr != nil {
		return oppErr
	}

	return nil
}

type NetworkErrorRatio struct {
	Operator  string  `yaml:"operator"`
	Parameter float64 `yaml:"parameter"`
}

func (n *NetworkErrorRatio) Expression() string {
	exp := fmt.Sprintf("%s %s %f", ExpressionNetworkErrRatio, n.Operator, n.Parameter)
	return exp
}

func (n *NetworkErrorRatio) Validate() error {
	opErr := validateOperator(n.Operator)
	if opErr != nil {
		return opErr
	}
	if n.Parameter < 0 {
		return fmt.Errorf("parameter %f cannot be less than 0", n.Parameter)
	}
	if n.Parameter > 1 {
		return fmt.Errorf("parameter %f cannot be greater than 1", n.Parameter)
	}
	return nil
}

//TODO: Change to slice

func WithAndArg(args ...CircuitBreakerExpression) func(s string) {
	return func(s string) {
		for _, arg := range args {
			s = fmt.Sprintf("%s && %s", s, arg.Expression())
		}
	}
}

func ExpressionWithAndArgs(args []CircuitBreakerExpression) func(s string) {
	return func(s string) {
		if s == "" {
			s = args[0].Expression()
			for _, arg := range args[1:] {
				s = fmt.Sprintf("%s && %s", s, arg.Expression())
			}
		} else {
			for _, arg := range args {
				s = fmt.Sprintf("%s && %s", s, arg.Expression())
			}
		}

	}
}

func WithOrArgs(args ...CircuitBreakerExpression) func(s string) {
	return func(s string) {
		for _, arg := range args {
			s = fmt.Sprintf("%s || %s", s, arg.Expression())
		}
	}
}

func ExpressionWithOrArgs(args []CircuitBreakerExpression) func(s string) {
	return func(s string) {
		if s == "" {
			s = args[0].Expression()
			for _, arg := range args[1:] {
				s = fmt.Sprintf("%s || %s", s, arg.Expression())
			}
		} else {
			for _, arg := range args {
				s = fmt.Sprintf("%s || %s", s, arg.Expression())
			}
		}

	}
}

func validateOperator(op string) error {
	switch op {
	case ">":
		return nil
	case ">=":
		return nil
	case "<":
		return nil
	case "<=":
		return nil
	case "==":
		return nil
	case "!=":
		return nil
	default:
		return fmt.Errorf("invalid operator '%s'", op)
	}
}
