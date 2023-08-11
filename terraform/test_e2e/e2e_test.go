package test

import (
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

var terraformOptions *terraform.Options

func setup() {
	terraformOptions = terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		VarFiles:     []string{"./terraform.tfvars", "./secret.tfvars"},
	})
}

func teardown() {
	terraform.Destroy(t, terraformOptions)
}

func TestMain(m *testing.M) {
	setup()
	defer teardown()
	code := m.Run()
	os.Exit(code)
}

func TestHelloWorld(t *testing.T) {
	t.Fatal("not implemented")
}
