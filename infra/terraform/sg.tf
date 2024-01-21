#####################
# SecurityGroup RDS #
#####################
resource "aws_security_group" "rds" {
  name        = "readerslounge-rds-sg"
  description = "RDS security group for readerslounge"
  vpc_id      = aws_vpc.readerslounge.id
}

#####################
# SecurityGroup ALB #
#####################
resource "aws_security_group" "alb" {
  name        = "readerslounge-alb-sg"
  description = "ALB security group for readerslounge"
  vpc_id      = aws_vpc.readerslounge.id
}

#####################
# SecurityGroup ECS #
#####################
resource "aws_security_group" "ecs_front" {
  name        = "readerslounge-ecs-front-sg"
  description = "ECS service security group for front"
  vpc_id      = aws_vpc.readerslounge.id
}

resource "aws_security_group" "ecs_api" {
  name        = "readerslounge-ecs-api-sg"
  description = "ECS service security group for api"
  vpc_id      = aws_vpc.readerslounge.id
}
