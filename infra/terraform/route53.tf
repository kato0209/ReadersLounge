###########
# Route53 #
###########
resource "aws_route53_zone" "readerslounge" {
  name    = "ReadersLounge.com"
  comment = "ReadersLounge.com host zone"
}

resource "aws_route53_record" "readerslounge" {
  zone_id = aws_route53_zone.readerslounge.zone_id
  name    = aws_route53_zone.readerslounge.name
  type    = "A"

  alias {
    name                   = aws_lb.readerslounge_alb.dns_name
    zone_id                = aws_lb.readerslounge_alb.zone_id
    evaluate_target_health = true
  }
}
