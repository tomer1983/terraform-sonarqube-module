package test

import (
	"testing"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestSonarqubeModule(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/complete",
		Vars: map[string]interface{}{
			"sonarqube_url":   "http://localhost:9000",
			"sonarqube_token": "admin",
		},
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	// Test project creation
	projectIds := terraform.Output(t, terraformOptions, "project_ids")
	assert.NotEmpty(t, projectIds)

	// Test quality gate creation
	qualityGateIds := terraform.Output(t, terraformOptions, "quality_gate_ids")
	assert.NotEmpty(t, qualityGateIds)
}
