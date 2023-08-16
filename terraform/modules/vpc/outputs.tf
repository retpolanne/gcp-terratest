output "vpc_name" {
  value = module.vpc.network_name
}

output "subnet_name" {
  value = module.vpc.subnets_names[0]
}
