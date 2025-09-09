package test

import (
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerragruntExample(t *testing.T) {
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/terragrunt-example",
		TerraformBinary: "terragrunt",
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.Apply(t, terraformOptions)

	output := terraform.Output(t, terraformOptions, "output")
	assert.Equal(t, "one input another input", output)
}

func TestTerragruntConsole(t *testing.T) {
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/terragrunt-example",
		TerraformBinary: "terragrunt",
		Stdin: strings.NewReader("local.mylocal"),
	})

	defer terraform.Destroy(t, terraformOptions)

	out := terraform.RunTerraformCommand(t, terraformOptions, "console")
	assert.Contains(t, out, `"local variable named mylocal"`)
}
