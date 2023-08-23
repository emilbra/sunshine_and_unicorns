package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformOutputs(t *testing.T) {

	expected_rg_name := "myRG"
	expected_location := "norwayeast"
	expected_kv_name := "uniqueemiltest"
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		Vars: map[string]interface{}{
			"rg_name":     expected_rg_name,
			"rg_location": expected_location,
			"my_kv_name":  expected_kv_name,
		},
		NoColor: true,
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	actual_kv_uri := terraform.Output(t, terraformOptions, "keyvault_uri" )
	actual_kv_name := terraform.Output(t, terraformOptions, "keyvault_name")
	actual_location := terraform.Output(t, terraformOptions, "keyvault_location")
	actual_rg_name := terraform.Output(t, terraformOptions, "rg_name")

	assert.NotEmpty(t, actual_kv_uri)
	assert.Equal(t, expected_kv_name, actual_kv_name)
	assert.Equal(t, expected_rg_name, actual_rg_name)
	assert.Equal(t, expected_location, actual_location)
}
