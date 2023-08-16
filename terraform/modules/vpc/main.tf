// https://registry.terraform.io/modules/terraform-google-modules/network/google/latest
data "google_client_config" "default" {}

module "vpc" {
  source       = "terraform-google-modules/network/google"
  project_id   = var.project_id
  network_name = var.vpc_name
  routing_mode = "GLOBAL"

  subnets = [
    {
      subnet_name   = "terratest-gke-subnet"
      subnet_ip     = "10.10.10.0/24"
      subnet_region = "us-east1"
    }
  ]

  secondary_ranges = {
    terratest-gke-subnet = [
      {
        range_name    = "terratest-gke-pods"
        ip_cidr_range = "10.0.0.0/16"
      },
      {
        range_name    = "terratest-gke-services"
        ip_cidr_range = "192.168.64.0/24"
      },
    ]
  }

  routes = [
    {
      name              = "interwebs"
      description       = "route through igw to internet"
      destination_range = "0.0.0.0/0"
      tags              = "interwebs-egress"
      next_hop_internet = true
    }
  ]
}
