resource "google_storage_bucket" "tf_backend" {
  name          = var.backend_bucket_name
  location      = "US"
  force_destroy = true

  uniform_bucket_level_access = true
}
