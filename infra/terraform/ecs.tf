###########
#   ECS   #
###########
resource "aws_ecs_cluster" "readerslounge" {
  name = "readerslounge-ecs-cluster"
}
resource "aws_ecs_task_definition" "front" {
  family                   = "readerslounge-front-task"
  cpu                      = "512"
  memory                   = "1024"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  container_definitions    = file("./tasks/front_definition.json")
  execution_role_arn       = aws_iam_role.ecs_task_execution.arn
}
resource "aws_ecs_service" "front" {
  name                              = "readerslounge-front-ecs-service"
  cluster                           = aws_ecs_cluster.readerslounge.arn
  task_definition                   = aws_ecs_task_definition.front.arn
  desired_count                     = 1
  launch_type                       = "FARGATE"
  health_check_grace_period_seconds = 600

  network_configuration {
    assign_public_ip = true
    security_groups = [
      aws_security_group.ecs_front.id
    ]
    subnets = [
      aws_subnet.public_1a.id,
      aws_subnet.public_1c.id
    ]
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.front.arn
    container_name   = "front-container"
    container_port   = "80"
  }
}

resource "aws_ecs_task_definition" "api" {
  family                   = "readerslounge-api-task"
  cpu                      = "256"
  memory                   = "512"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  container_definitions    = file("./tasks/api_definition.json")
  execution_role_arn       = aws_iam_role.ecs_task_execution.arn
}
resource "aws_ecs_service" "api" {
  name            = "readerslounge-api-ecs-service"
  cluster         = aws_ecs_cluster.readerslounge.arn
  task_definition = aws_ecs_task_definition.api.arn
  desired_count   = 1
  launch_type     = "FARGATE"

  network_configuration {
    assign_public_ip = true
    security_groups = [
      aws_security_group.ecs_api.id
    ]
    subnets = [
      aws_subnet.public_1a.id,
      aws_subnet.public_1c.id
    ]
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.api.arn
    container_name   = "api-container"
    container_port   = "8080"
  }
}

##########
#  権限   #
##########
resource "aws_iam_role" "ecs_task_execution" {
  name = "ecs_task_execution"

  assume_role_policy = jsonencode({
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid    = ""
        Principal = {
          Service = "ecs-tasks.amazonaws.com"
        }
      },
    ]
  })
}
