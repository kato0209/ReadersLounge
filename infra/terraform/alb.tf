#########
#  ALB  #
#########
resource "aws_lb" "readerslounge" {
  name                       = "readerslounge-alb"
  load_balancer_type         = "application"
  internal                   = false
  idle_timeout               = 60
  enable_deletion_protection = false

  subnets = [
    aws_subnet.public_1a.id,
    aws_subnet.public_1c.id
  ]

  security_groups = [
    aws_security_group.alb.id
  ]

  tags = {
    Name = "readerslounge-alb"
  }
}

###############
# TargetGroup #
###############
resource "aws_lb_target_group" "front" {
  name        = "readerslounge-alb-front-tg"
  target_type = "ip"
  vpc_id      = aws_vpc.readerslounge.id
  port        = 80
  protocol    = "HTTP"

  health_check {
    enabled             = true
    path                = "/"
    healthy_threshold   = 2
    unhealthy_threshold = 2
    timeout             = 120
    interval            = 150
    matcher             = 200
    port                = 80
    protocol            = "HTTP"
  }

  depends_on = [
    aws_lb.readerslounge
  ]
}
resource "aws_lb_target_group" "api" {
  name        = "readerslounge-alb-api-tg"
  target_type = "ip"
  vpc_id      = aws_vpc.readerslounge.id
  port        = 8080
  protocol    = "HTTP"

  health_check {
    enabled             = true
    interval            = 60
    path                = "/"
    port                = 8080
    protocol            = "HTTP"
    matcher             = 200
    timeout             = 50
    healthy_threshold   = 5
    unhealthy_threshold = 2
  }

  depends_on = [
    aws_lb.readerslounge
  ]
}

############
# Listener #
############
resource "aws_lb_listener" "http" {
  load_balancer_arn = aws_lb.readerslounge.arn
  port              = "80"
  protocol          = "HTTP"

  default_action {
    type = "redirect"

    redirect {
      port        = "443"
      protocol    = "HTTPS"
      status_code = "HTTP_301"
    }
  }
  depends_on = [
    aws_lb.readerslounge
  ]
}
resource "aws_lb_listener" "https" {
  load_balancer_arn = aws_lb.readerslounge.arn
  port              = "443"
  protocol          = "HTTPS"
  certificate_arn   = aws_acm_certificate_validation.readerslounge.certificate_arn

  default_action {
    target_group_arn = aws_lb_target_group.front.arn
    type             = "forward"
  }
  depends_on = [
    aws_lb.readerslounge,
    aws_lb_target_group.front
  ]
}
resource "aws_lb_listener" "api" {
  load_balancer_arn = aws_lb.readerslounge.arn
  port              = "8080"
  protocol          = "HTTPS"
  certificate_arn   = aws_acm_certificate_validation.readerslounge.certificate_arn

  default_action {
    target_group_arn = aws_lb_target_group.api.arn
    type             = "forward"
  }
  depends_on = [
    aws_lb.readerslounge,
    aws_lb_target_group.api
  ]
}
