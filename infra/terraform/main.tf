terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 4.1.0"
    }
  }
  backend "s3" {
    bucket  = "readerslounge-terraform-state"
    key     = "terraform.tfstate"
    region  = "ap-northeast-1"
    encrypt = true
  }
}

locals {
  app_name = "ReadersLounge"
}

provider "aws" {
  profile = "terraform"
  region  = "ap-northeast-1"
  default_tags {
    tags = {
      application = local.app_name
    }
  }
}
