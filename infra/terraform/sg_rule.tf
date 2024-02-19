#########################
# SecurityGroupRule RDS #
#########################
resource "aws_security_group_rule" "rds1" {
  description       = "readerslounge rds sg rule1"
  type              = "ingress"
  from_port         = 5432
  to_port           = 5432
  protocol          = "tcp"
  cidr_blocks       = ["0.0.0.0/0"]
  security_group_id = aws_security_group.rds.id
}

resource "aws_security_group_rule" "rds2" {
  description       = "readerslounge rds sg rule2"
  type              = "egress"
  from_port         = 0
  to_port           = 0
  protocol          = "-1"
  cidr_blocks       = ["0.0.0.0/0"]
  security_group_id = aws_security_group.rds.id
}

#########################
# SecurityGroupRule ALB #
#########################
resource "aws_security_group_rule" "alb1" {
  description       = "readerslounge-alb-sg-rule1"
  type              = "ingress"
  from_port         = 80
  to_port           = 80
  protocol          = "tcp"
  cidr_blocks       = ["0.0.0.0/0"]
  security_group_id = aws_security_group.alb.id
}
resource "aws_security_group_rule" "alb2" {
  description       = "readerslounge-alb-sg-rule2"
  type              = "ingress"
  from_port         = 443
  to_port           = 443
  protocol          = "tcp"
  cidr_blocks       = ["0.0.0.0/0"]
  security_group_id = aws_security_group.alb.id
}
resource "aws_security_group_rule" "alb3" {
  description       = "readerslounge-alb-sg-rule3"
  type              = "egress"
  from_port         = 0
  to_port           = 0
  protocol          = "-1"
  cidr_blocks       = ["0.0.0.0/0"]
  security_group_id = aws_security_group.alb.id
}
resource "aws_security_group_rule" "alb4" {
  description       = "readerslounge-alb-sg-rule4"
  type              = "ingress"
  from_port         = 8080
  to_port           = 8080
  protocol          = "tcp"
  cidr_blocks       = ["0.0.0.0/0"]
  security_group_id = aws_security_group.alb.id
}

#########################
# SecurityGroupRule ECS #
#########################

#__________  front  __________#
resource "aws_security_group_rule" "ecs_front1" {
  description       = "readerslounge-ecs-front-sg-rule1"
  type              = "egress"
  from_port         = 0
  to_port           = 0
  protocol          = "-1"
  cidr_blocks       = ["0.0.0.0/0"]
  security_group_id = aws_security_group.ecs_front.id
}
resource "aws_security_group_rule" "ecs_front2" {
  description              = "readerslounge-ecs-front-sg-rule2"
  type                     = "ingress"
  from_port                = 80
  to_port                  = 80
  protocol                 = "tcp"
  source_security_group_id = aws_security_group.alb.id
  security_group_id        = aws_security_group.ecs_front.id
}
resource "aws_security_group_rule" "ecs_front3" {
  description              = "readerslounge-ecs-front-sg-rule3"
  type                     = "ingress"
  from_port                = 443
  to_port                  = 443
  protocol                 = "tcp"
  source_security_group_id = aws_security_group.alb.id
  security_group_id        = aws_security_group.ecs_front.id
}

#__________  api  __________#
resource "aws_security_group_rule" "ecs_api1" {
  description              = "readerslounge-ecs-api-sg-rule1"
  type                     = "ingress"
  from_port                = 8080
  to_port                  = 8080
  protocol                 = "tcp"
  source_security_group_id = aws_security_group.alb.id
  security_group_id        = aws_security_group.ecs_api.id
}
resource "aws_security_group_rule" "ecs_api2" {
  description       = "readerslounge-ecs-api-sg-rule2"
  type              = "egress"
  from_port         = 0
  to_port           = 0
  protocol          = "-1"
  cidr_blocks       = ["0.0.0.0/0"]
  security_group_id = aws_security_group.ecs_api.id
}
