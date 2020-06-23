// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint/tflint"
)

func Test_AwsDirectoryServiceDirectoryInvalidDescriptionRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_directory_service_directory" "foo" {
	description = "@example"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsDirectoryServiceDirectoryInvalidDescriptionRule(),
					Message: `"@example" does not match valid pattern ^([a-zA-Z0-9_])[\\a-zA-Z0-9_@#%*+=:?./!\s-]*$`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_directory_service_directory" "foo" {
	description = "example"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsDirectoryServiceDirectoryInvalidDescriptionRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
