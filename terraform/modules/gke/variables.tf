variable "project_id" {
  type        = string
  description = "GCP's project ID to create the resource"
}

variable "cluster_name" {
  type        = string
  description = "The name of the cluster to be created"
}

variable "vpc_name" {
  type        = string
  description = "The name of the vpc to use"
}

variable "subnet_name" {
  type        = string
  description = "The name of the subnet to use"
}
