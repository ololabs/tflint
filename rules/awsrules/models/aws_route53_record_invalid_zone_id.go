// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsRoute53RecordInvalidZoneIDRule checks the pattern is valid
type AwsRoute53RecordInvalidZoneIDRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsRoute53RecordInvalidZoneIDRule returns new rule with default attributes
func NewAwsRoute53RecordInvalidZoneIDRule() *AwsRoute53RecordInvalidZoneIDRule {
	return &AwsRoute53RecordInvalidZoneIDRule{
		resourceType:  "aws_route53_record",
		attributeName: "zone_id",
		max:           32,
	}
}

// Name returns the rule name
func (r *AwsRoute53RecordInvalidZoneIDRule) Name() string {
	return "aws_route53_record_invalid_zone_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRoute53RecordInvalidZoneIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRoute53RecordInvalidZoneIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRoute53RecordInvalidZoneIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRoute53RecordInvalidZoneIDRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"zone_id must be 32 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
