// https://github.com/terraform-google-modules/terraform-google-kubernetes-engine
data "google_client_config" "default" {}

module "gke" {
  source  		      = "terraform-google-modules/kubernetes-engine/google"
	project_id 	      = var.project_id
  name              = var.cluster_name
  region            = "us-east1"
  network           = var.vpc_name
  subnetwork        = var.subnet_name
  ip_range_pods     = "terratest-gke-pods"
  ip_range_services = "terratest-gke-services"
}

module "gke_auth" {
	source       = "terraform-google-modules/kubernetes-engine/google//modules/auth"
  project_id   = var.project_id
  cluster_name = var.cluster_name
  location     = module.gke.location
  depends_on   = [module.gke]
}

resource "local_file" "kubeconfig" {
	content    = module.gke_auth.kubeconfig_raw
  filename   = "../../kubeconfig" 
  depends_on = [module.gke_auth]
}
