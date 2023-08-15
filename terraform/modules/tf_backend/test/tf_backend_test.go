package test

import (
	"flag"
	"testing"

	"github.com/gruntwork-io/terratest/modules/gcp"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

var destroy = flag.Bool("destroy", false, "In case you need to destroy the resources")

func TestMain(t *testing.T) {
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		VarFiles:     []string{"./terraform.tfvars"},
	})
	if *destroy {
		defer terraform.Destroy(t, terraformOptions)
	}
	terraform.InitAndApply(t, terraformOptions)
	// We usually defer terraform Destroy, but since we may want to ensure things
	// instead of doing e2e tests, we can just comment this
	//defer terraform.Destroy(t, terraformOptions)
	TestTFBackendStorageBucketExists(t)
}

func TestTFBackendStorageBucketExists(t *testing.T) {
	gcp.AssertStorageBucketExists(t, "annemacedo-tf-backend-demo")
}
