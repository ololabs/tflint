// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsGameliftBuildInvalidVersionRule checks the pattern is valid
type AwsGameliftBuildInvalidVersionRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsGameliftBuildInvalidVersionRule returns new rule with default attributes
func NewAwsGameliftBuildInvalidVersionRule() *AwsGameliftBuildInvalidVersionRule {
	return &AwsGameliftBuildInvalidVersionRule{
		resourceType:  "aws_gamelift_build",
		attributeName: "version",
		max:           1024,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsGameliftBuildInvalidVersionRule) Name() string {
	return "aws_gamelift_build_invalid_version"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsGameliftBuildInvalidVersionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsGameliftBuildInvalidVersionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsGameliftBuildInvalidVersionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsGameliftBuildInvalidVersionRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"version must be 1024 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"version must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
