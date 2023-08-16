# VPC Module

Please set these environment variables: 

```sh
export GOOGLE_PROJECT=gcp-project-id
export TF_BACKEND_BUCKET_NAME=backend-bucket
```

Testing:

```sh
go test -v ./...
```

If you need to destroy after testing: 

```sh
go test -v ./... -destroy
```
