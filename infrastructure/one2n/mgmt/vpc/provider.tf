provider "aws" {
  version                 = "~> 3.0"
  region                  = "ap-south-1"
  shared_credentials_file = "~/.aws/credentials"
  profile                 = "dragtor"
}

terraform {
  backend "s3" {
    bucket                  = "dragtor-terraform-remote-state"
    key                     = "infra/one2n/init/terraform.tfstate"
    region                  = "ap-south-1"
    shared_credentials_file = "~/.aws/credentials"
    profile                 = "dragtor"
  }
}


