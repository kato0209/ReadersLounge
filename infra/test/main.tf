terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 4.1.0"
    }
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


##########
#  ACM   #
##########

##SSL証明書の定義
resource "aws_acm_certificate" "readerslounge" {
  domain_name       = "readerslounge.com"
  validation_method = "DNS"
  lifecycle {
    create_before_destroy = true
  }
}
##SSL検証
resource "aws_route53_record" "readerslounge_certificate" {
  for_each = {
    for dvo in aws_acm_certificate.readerslounge.domain_validation_options : dvo.domain_name => {
      name   = dvo.resource_record_name
      record = dvo.resource_record_value
      type   = dvo.resource_record_type
    }
  }

  allow_overwrite = true
  name            = each.value.name
  records         = [each.value.record]
  ttl             = 60
  type            = each.value.type
  zone_id         = aws_route53_zone.readerslounge.zone_id
}
##検証待機
resource "aws_acm_certificate_validation" "readerslounge" {
  certificate_arn         = aws_acm_certificate.readerslounge.arn
  validation_record_fqdns = [for record in aws_route53_record.readerslounge_certificate : record.fqdn]
}

###########
# Route53 #
###########
resource "aws_route53_zone" "readerslounge" {
  name = "readerslounge.com"
}


