variable "project_id" {
  type        = string
  description = "GCP's project ID to create the Terraform backend bucket"
}

variable "backend_bucket_name" {
  type        = string
  description = "Name of the bucket to be created for Terraform backend"
}
