// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsIAMPolicyInvalidPathRule checks the pattern is valid
type AwsIAMPolicyInvalidPathRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsIAMPolicyInvalidPathRule returns new rule with default attributes
func NewAwsIAMPolicyInvalidPathRule() *AwsIAMPolicyInvalidPathRule {
	return &AwsIAMPolicyInvalidPathRule{
		resourceType:  "aws_iam_policy",
		attributeName: "path",
		max:           512,
		min:           1,
		pattern:       regexp.MustCompile(`^((/[A-Za-z0-9\.,\+@=_-]+)*)/$`),
	}
}

// Name returns the rule name
func (r *AwsIAMPolicyInvalidPathRule) Name() string {
	return "aws_iam_policy_invalid_path"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsIAMPolicyInvalidPathRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsIAMPolicyInvalidPathRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsIAMPolicyInvalidPathRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsIAMPolicyInvalidPathRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"path must be 512 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"path must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^((/[A-Za-z0-9\.,\+@=_-]+)*)/$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
