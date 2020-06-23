// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsElasticBeanstalkApplicationVersionInvalidKeyRule checks the pattern is valid
type AwsElasticBeanstalkApplicationVersionInvalidKeyRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsElasticBeanstalkApplicationVersionInvalidKeyRule returns new rule with default attributes
func NewAwsElasticBeanstalkApplicationVersionInvalidKeyRule() *AwsElasticBeanstalkApplicationVersionInvalidKeyRule {
	return &AwsElasticBeanstalkApplicationVersionInvalidKeyRule{
		resourceType:  "aws_elastic_beanstalk_application_version",
		attributeName: "key",
		max:           1024,
	}
}

// Name returns the rule name
func (r *AwsElasticBeanstalkApplicationVersionInvalidKeyRule) Name() string {
	return "aws_elastic_beanstalk_application_version_invalid_key"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsElasticBeanstalkApplicationVersionInvalidKeyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsElasticBeanstalkApplicationVersionInvalidKeyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsElasticBeanstalkApplicationVersionInvalidKeyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsElasticBeanstalkApplicationVersionInvalidKeyRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"key must be 1024 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
