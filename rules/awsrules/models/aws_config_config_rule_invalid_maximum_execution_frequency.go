// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsConfigConfigRuleInvalidMaximumExecutionFrequencyRule checks the pattern is valid
type AwsConfigConfigRuleInvalidMaximumExecutionFrequencyRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsConfigConfigRuleInvalidMaximumExecutionFrequencyRule returns new rule with default attributes
func NewAwsConfigConfigRuleInvalidMaximumExecutionFrequencyRule() *AwsConfigConfigRuleInvalidMaximumExecutionFrequencyRule {
	return &AwsConfigConfigRuleInvalidMaximumExecutionFrequencyRule{
		resourceType:  "aws_config_config_rule",
		attributeName: "maximum_execution_frequency",
		enum: []string{
			"One_Hour",
			"Three_Hours",
			"Six_Hours",
			"Twelve_Hours",
			"TwentyFour_Hours",
		},
	}
}

// Name returns the rule name
func (r *AwsConfigConfigRuleInvalidMaximumExecutionFrequencyRule) Name() string {
	return "aws_config_config_rule_invalid_maximum_execution_frequency"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsConfigConfigRuleInvalidMaximumExecutionFrequencyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsConfigConfigRuleInvalidMaximumExecutionFrequencyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsConfigConfigRuleInvalidMaximumExecutionFrequencyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsConfigConfigRuleInvalidMaximumExecutionFrequencyRule) Check(runner *tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as maximum_execution_frequency`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
