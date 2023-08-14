# GCP Terratest

This is an example repo for creating stuff on GCP using Terratest. 

## Installing needed tools

If you're on a Mac, use `brew bundle` and this will install:

- Terraform
- Golang
- Google Cloud SDK

## Logging in with gcloud

```sh
gcloud auth login
gcloud auth application-default login
```

## Running tests

```sh
go test ./...
```
