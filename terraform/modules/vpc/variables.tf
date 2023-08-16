variable "project_id" {
  type        = string
  description = "GCP's project ID to create the resource"
}

variable "vpc_name" {
  type        = string
  description = "The name of the vpc to be created"
  default     = "main-vpc"
}
