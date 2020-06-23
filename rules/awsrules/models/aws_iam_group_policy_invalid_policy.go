// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsIAMGroupPolicyInvalidPolicyRule checks the pattern is valid
type AwsIAMGroupPolicyInvalidPolicyRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsIAMGroupPolicyInvalidPolicyRule returns new rule with default attributes
func NewAwsIAMGroupPolicyInvalidPolicyRule() *AwsIAMGroupPolicyInvalidPolicyRule {
	return &AwsIAMGroupPolicyInvalidPolicyRule{
		resourceType:  "aws_iam_group_policy",
		attributeName: "policy",
		max:           131072,
		min:           1,
		pattern:       regexp.MustCompile(`^[\x{0009}\x{000A}\x{000D}\x{0020}-\x{00FF}]+$`),
	}
}

// Name returns the rule name
func (r *AwsIAMGroupPolicyInvalidPolicyRule) Name() string {
	return "aws_iam_group_policy_invalid_policy"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsIAMGroupPolicyInvalidPolicyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsIAMGroupPolicyInvalidPolicyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsIAMGroupPolicyInvalidPolicyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsIAMGroupPolicyInvalidPolicyRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"policy must be 131072 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"policy must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[\x{0009}\x{000A}\x{000D}\x{0020}-\x{00FF}]+$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
