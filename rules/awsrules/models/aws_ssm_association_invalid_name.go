// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsSsmAssociationInvalidNameRule checks the pattern is valid
type AwsSsmAssociationInvalidNameRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsSsmAssociationInvalidNameRule returns new rule with default attributes
func NewAwsSsmAssociationInvalidNameRule() *AwsSsmAssociationInvalidNameRule {
	return &AwsSsmAssociationInvalidNameRule{
		resourceType:  "aws_ssm_association",
		attributeName: "name",
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9_\-.:/]{3,128}$`),
	}
}

// Name returns the rule name
func (r *AwsSsmAssociationInvalidNameRule) Name() string {
	return "aws_ssm_association_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsmAssociationInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsmAssociationInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsmAssociationInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsmAssociationInvalidNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z0-9_\-.:/]{3,128}$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
