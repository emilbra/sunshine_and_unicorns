package test

// Importing the packages to use, mainly the terratest and assert packages.
import (
	"testing"

	// Made by terratest, allows go to interact with terraform
	"github.com/gruntwork-io/terratest/modules/terraform"

	// Testing tools to complement the built-in ones. Usefull for comparison-operations, for example.
	"github.com/stretchr/testify/assert"
)

// The function containg the test.
func TestTerraformOutputs(t *testing.T) {

	// First, we shoul define a couple of values that we want to test against.
	// If an "actual" value does not match an expected value, something is awry

	// You can hard code these
	expected_rg_name := "myRG"

	// or use the GetVariableAsStringFromVarFile - function to get inputs from a varfile
	// the third option in the function corresponds to a specific variable name as defined in the file.
	// this is the preferred option.
	
	expected_location := terraform.GetVariableAsStringFromVarFile(t, "../test.tfvars", "rg_location")
	expected_kv_name := terraform.GetVariableAsStringFromVarFile(t, "../test.tfvars", "kv_name")

	// WithDefaultRetryableErros copies the Options object in the terraform package, but adds sensible defaults for errors commonly resolved by retries
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Here, we are defining additional options as part of the .Options-object
		TerraformDir: "../",
		// Set what variables to pass while running terraform, this is equivalent to setting -var flags while running by command line.
		Vars: map[string]interface{}{
			"rg_name":     expected_rg_name,
			"rg_location": expected_location,
			"my_kv_name":  expected_kv_name,
		},
		NoColor: true,
	})
	// Defer is evaluated immediately, but is not executed until everything else in this function returns. Basically, Destroy will run at the end, no matter other outcomes.
	defer terraform.Destroy(t, terraformOptions)

	// Runs terraform init and apply - just like a basic workflow.
	terraform.InitAndApply(t, terraformOptions)

	// Sets some values to use for comparisons, these are all using terraform.Output to get values from the outputs in outputs.tf
	actual_kv_uri := terraform.Output(t, terraformOptions, "keyvault_uri" )
	actual_kv_name := terraform.Output(t, terraformOptions, "keyvault_name")
	actual_location := terraform.Output(t, terraformOptions, "keyvault_location")
	actual_rg_name := terraform.Output(t, terraformOptions, "rg_name")

	// Assert that an keyvault_uri is actually returned. Remember, this is a value we ourselves never made, but which azure creates automatically.
	assert.NotEmpty(t, actual_kv_uri)

	// Assert that actual values match expected values
	assert.Equal(t, expected_kv_name, actual_kv_name)
	assert.Equal(t, expected_rg_name, actual_rg_name)
	assert.Equal(t, expected_location, actual_location)
}
