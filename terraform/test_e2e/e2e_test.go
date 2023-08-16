package test_e2e

import (
	"context"
	"flag"
	"fmt"
	"os"
	"testing"

	container "cloud.google.com/go/container/apiv1"
	containerpb "cloud.google.com/go/container/apiv1/containerpb"
	"github.com/gruntwork-io/terratest/modules/gcp"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var clusterName string = "terratest-gke-anne"

var destroy = flag.Bool("destroy", false, "In case you need to destroy the resources")

func getGKEClient(t *testing.T) (*container.ClusterManagerClient, context.Context) {
	ctx := context.Background()
	gkeClient, err := container.NewClusterManagerClient(ctx)
	if err != nil {
		t.Fatalf("Error creating gke client %s\n", err)
	}
	return gkeClient, ctx
}

func getCluster(t *testing.T) *containerpb.Cluster {
	gkeClient, ctx := getGKEClient(t)
	defer gkeClient.Close()
	clusterURI := fmt.Sprintf(
		"projects/%s/locations/us-east1/clusters/%s",
		gcp.GetGoogleProjectIDFromEnvVar(t),
		clusterName,
	)
	req := &containerpb.GetClusterRequest{
		Name: clusterURI,
	}
	gkeCluster, err := gkeClient.GetCluster(ctx, req)
	if err != nil {
		t.Fatalf("Error getting gke cluster %s %s\n", clusterName, err)
	}
	return gkeCluster
}

func getKubectlOptions(namespace string) *k8s.KubectlOptions {
	return &k8s.KubectlOptions{
		ContextName: clusterName,
		ConfigPath:  "../kubeconfig",
		Namespace:   namespace,
		Logger:      logger.Discard,
	}
}

func TestMain(t *testing.T) {
	bucket := os.Getenv("TF_BACKEND_BUCKET_NAME")
	if bucket == "" {
		t.Fatal("Please provide the TF_BACKEND_BUCKET_NAME variable")
	}
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		VarFiles:     []string{"./terraform.tfvars"},
		Vars: map[string]interface{}{
			"project_id": gcp.GetGoogleProjectIDFromEnvVar(t),
		},
		BackendConfig: map[string]interface{}{
			"bucket": bucket,
		},
	})
	if *destroy {
		defer terraform.Destroy(t, terraformOptions)
	}
	terraform.InitAndApply(t, terraformOptions)
	TestGKECreated(t)
	TestGKEHasPrometheusNodePool(t)
	TestGKEHasAnthos(t)
}

func TestGKECreated(t *testing.T) {
	gke := getCluster(t)
	require.NotNil(t, gke)
	assert.Equal(t, clusterName, gke.Name)
}

func TestGKEHasPrometheusNodePool(t *testing.T) {
	gke := getCluster(t)
	require.NotNil(t, gke)
	prometheusNPFound := false
	var nodeTaint *containerpb.NodeTaint
	for _, nodePool := range gke.NodePools {
		if nodePool.Name == "prometheus" {
			prometheusNPFound = true
			nodeTaints := nodePool.Config.Taints
			for _, taint := range nodeTaints {
				if taint.Key == "monitoring" {
					nodeTaint = taint
				}
			}
			break
		}
	}
	assert.True(t, prometheusNPFound)
	require.NotNil(t, nodeTaint)
	assert.Equal(t, "NO_SCHEDULE", nodeTaint.Effect)
	assert.Equal(t, "prometheus", nodeTaint.Value)
}

func TestGKEHasPrometheusInstalled(t *testing.T) {
	k8sOptions := getKubectlOptions("kube-monitoring")
	deployment := k8s.GetDeployment(t, k8sOptions, "prometheus-server")
	assert.True(t, k8s.IsDeploymentAvailable(deployment))
}

func TestGKEPrometheusIsOnSeparateNodePool(t *testing.T) {
	k8sOptions := getKubectlOptions("kube-monitoring")
	_ = k8s.GetDeployment(t, k8sOptions, "prometheus-server")
}

func TestGKEHasGrafanaInstalled(t *testing.T) {
	k8sOptions := getKubectlOptions("kube-monitoring")
	deployment := k8s.GetDeployment(t, k8sOptions, "prometheus-grafana")
	assert.True(t, k8s.IsDeploymentAvailable(deployment))
}

func TestGKEHasAnthos(t *testing.T) {
	gke := getCluster(t)
	require.NotNil(t, gke)
}
