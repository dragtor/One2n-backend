
module "vpc" {
  source       = "./../../../modules/vpc"
  vpc_cidr     = var.vpc_cidr
  vpc_tag_name = var.vpc_tag_name
  igw_tag_name = var.igw_tag_name
}


output "vpc_id" {
  value = module.vpc.vpc_id
}

output "subnet_ids" {
  value = module.vpc.subnet_ids
}

output "private_subnet_ids" {
  value = module.vpc.private_subnet_ids
}
