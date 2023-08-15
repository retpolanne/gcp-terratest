resource "google_container_cluster" "gke_terratest_demo" {
  name                = var.cluster_name
  location            = "us-east1"
  initial_node_count  = 1
}
