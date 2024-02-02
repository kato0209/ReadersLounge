###########
# Route53 #
###########
data "aws_route53_zone" "readerslounge" {
  name    = "readerslounge-server.com"
}

resource "aws_route53_record" "readerslounge" {
  zone_id = data.aws_route53_zone.readerslounge.zone_id
  name    = data.aws_route53_zone.readerslounge.name
  type    = "A"

  alias {
    name                   = aws_lb.readerslounge.dns_name
    zone_id                = aws_lb.readerslounge.zone_id
    evaluate_target_health = true
  }
}
