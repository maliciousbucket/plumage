package k8s

// ValidatingAdmissionPolicySpec is the specification of the desired behavior of the AdmissionPolicy.
type ValidatingAdmissionPolicySpecV1Beta1 struct {
	// auditAnnotations contains CEL expressions which are used to produce audit annotations for the audit event of the API request.
	//
	// validations and auditAnnotations may not both be empty; a least one of validations or auditAnnotations is required.
	AuditAnnotations *[]*AuditAnnotationV1Beta1 `field:"optional" json:"auditAnnotations" yaml:"auditAnnotations"`
	// failurePolicy defines how to handle failures for the admission policy.
	//
	// Failures can occur from CEL expression parse errors, type check errors, runtime errors and invalid or mis-configured policy definitions or bindings.
	//
	// A policy is invalid if spec.paramKind refers to a non-existent Kind. A binding is invalid if spec.paramRef.name refers to a non-existent resource.
	//
	// failurePolicy does not define how validations that evaluate to false are handled.
	//
	// When failurePolicy is set to Fail, ValidatingAdmissionPolicyBinding validationActions define how failures are enforced.
	//
	// Allowed values are Ignore or Fail. Defaults to Fail.
	// Default: Fail.
	//
	FailurePolicy *string `field:"optional" json:"failurePolicy" yaml:"failurePolicy"`
	// MatchConditions is a list of conditions that must be met for a request to be validated.
	//
	// Match conditions filter requests that have already been matched by the rules, namespaceSelector, and objectSelector. An empty list of matchConditions matches all requests. There are a maximum of 64 match conditions allowed.
	//
	// If a parameter object is provided, it can be accessed via the `params` handle in the same manner as validation expressions.
	//
	// The exact matching logic is (in order):
	// 1. If ANY matchCondition evaluates to FALSE, the policy is skipped.
	// 2. If ALL matchConditions evaluate to TRUE, the policy is evaluated.
	// 3. If any matchCondition evaluates to an error (but none are FALSE):
	// - If failurePolicy=Fail, reject the request
	// - If failurePolicy=Ignore, the policy is skipped.
	MatchConditions *[]*MatchConditionV1Beta1 `field:"optional" json:"matchConditions" yaml:"matchConditions"`
	// MatchConstraints specifies what resources this policy is designed to validate.
	//
	// The AdmissionPolicy cares about a request if it matches _all_ Constraints. However, in order to prevent clusters from being put into an unstable state that cannot be recovered from via the API ValidatingAdmissionPolicy cannot match ValidatingAdmissionPolicy and ValidatingAdmissionPolicyBinding. Required.
	MatchConstraints *MatchResourcesV1Beta1 `field:"optional" json:"matchConstraints" yaml:"matchConstraints"`
	// ParamKind specifies the kind of resources used to parameterize this policy.
	//
	// If absent, there are no parameters for this policy and the param CEL variable will not be provided to validation expressions. If ParamKind refers to a non-existent kind, this policy definition is mis-configured and the FailurePolicy is applied. If paramKind is specified but paramRef is unset in ValidatingAdmissionPolicyBinding, the params variable will be null.
	ParamKind *ParamKindV1Beta1 `field:"optional" json:"paramKind" yaml:"paramKind"`
	// Validations contain CEL expressions which is used to apply the validation.
	//
	// Validations and AuditAnnotations may not both be empty; a minimum of one Validations or AuditAnnotations is required.
	Validations *[]*ValidationV1Beta1 `field:"optional" json:"validations" yaml:"validations"`
	// Variables contain definitions of variables that can be used in composition of other expressions.
	//
	// Each variable is defined as a named CEL expression. The variables defined here will be available under `variables` in other expressions of the policy except MatchConditions because MatchConditions are evaluated before the rest of the policy.
	//
	// The expression of a variable can refer to other variables defined earlier in the list but not those after. Thus, Variables must be sorted by the order of first appearance and acyclic.
	Variables *[]*VariableV1Beta1 `field:"optional" json:"variables" yaml:"variables"`
}
