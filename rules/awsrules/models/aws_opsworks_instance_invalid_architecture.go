// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsOpsworksInstanceInvalidArchitectureRule checks the pattern is valid
type AwsOpsworksInstanceInvalidArchitectureRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsOpsworksInstanceInvalidArchitectureRule returns new rule with default attributes
func NewAwsOpsworksInstanceInvalidArchitectureRule() *AwsOpsworksInstanceInvalidArchitectureRule {
	return &AwsOpsworksInstanceInvalidArchitectureRule{
		resourceType:  "aws_opsworks_instance",
		attributeName: "architecture",
		enum: []string{
			"x86_64",
			"i386",
		},
	}
}

// Name returns the rule name
func (r *AwsOpsworksInstanceInvalidArchitectureRule) Name() string {
	return "aws_opsworks_instance_invalid_architecture"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsOpsworksInstanceInvalidArchitectureRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsOpsworksInstanceInvalidArchitectureRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsOpsworksInstanceInvalidArchitectureRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsOpsworksInstanceInvalidArchitectureRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as architecture`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
