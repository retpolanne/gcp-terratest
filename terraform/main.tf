module "gke_vpc" {
  source     = "./modules/vpc"
  project_id = var.project_id
}

module "terratest_gke" {
  source       = "./modules/gke"
  project_id   = var.project_id
  cluster_name = var.cluster_name
  vpc_name     = module.gke_vpc.vpc_name
  subnet_name  = module.gke_vpc.subnet_name
}
