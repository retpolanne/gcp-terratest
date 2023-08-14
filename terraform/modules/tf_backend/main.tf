resource "google_storage_bucket" "tf_backend" {
  name          = "annemacedo-tf-backend-demo"
  location      = "US"
  force_destroy = true

  uniform_bucket_level_access = true
}
