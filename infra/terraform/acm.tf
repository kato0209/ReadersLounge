##########
#  ACM   #
##########

##SSL証明書の定義
resource "aws_acm_certificate" "readerslounge" {
  domain_name               = aws_route53_record.readerslounge.name
  subject_alternative_names = []
  validation_method         = "DNS"
  lifecycle {
    create_before_destroy = true
  }
}
##SSL検証
resource "aws_route53_record" "readerslounge_certificate" {
  name    = tolist(aws_acm_certificate.readerslounge.domain_validation_options)[0].resource_record_name
  type    = tolist(aws_acm_certificate.readerslounge.domain_validation_options)[0].resource_record_type
  records = [tolist(aws_acm_certificate.readerslounge.domain_validation_options)[0].resource_record_value]
  zone_id = aws_route53_zone.readerslounge.id
  ttl     = 60
}
##検証待機
resource "aws_acm_certificate_validation" "readerslounge" {
  certificate_arn         = aws_acm_certificate.readerslounge.arn
  validation_record_fqdns = [aws_route53_record.readerslounge_certificate.fqdn]
}
