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
  container_definitions    = file("./tasks/gadget_front_definition.json")
  execution_role_arn       = module.ecs_task_execution_role.iam_role_arn
}
resource "aws_ecs_service" "front" {
  name             = "readerslounge-front-ecs-service"
  cluster          = aws_ecs_cluster.readerslounge.arn
  task_definition  = aws_ecs_task_definition.front.arn
  desired_count    = 1
  launch_type      = "FARGATE"
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
  container_definitions    = file("./tasks/gadget_api_definition.json")
  execution_role_arn       = module.ecs_task_execution_role.iam_role_arn
}
resource "aws_ecs_service" "api" {
  name             = "readerslounge-api-ecs-service"
  cluster          = aws_ecs_cluster.readerslounge.arn
  task_definition  = aws_ecs_task_definition.api.arn
  desired_count    = 1
  launch_type      = "FARGATE"

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

resource "aws_ecs_task_definition" "db-create" {
  family                   = "gadget-db-create"
  container_definitions    = file("./tasks/gadget_db_create_definition.json")
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  cpu                      = "256"
  memory                   = "512"
  execution_role_arn       = module.ecs_task_execution_role.iam_role_arn
}
resource "aws_ecs_task_definition" "db-migrate" {
  family                   = "gadget-db-migrate"
  container_definitions    = file("./tasks/gadget_db_migrate_definition.json")
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  cpu                      = "256"
  memory                   = "512"
  execution_role_arn       = module.ecs_task_execution_role.iam_role_arn
}
resource "aws_ecs_task_definition" "db-migrate-reset" {
  family                   = "gadget-db-migrate-reset"
  container_definitions    = file("./tasks/gadget_db_migrate_reset_definition.json")
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  cpu                      = "256"
  memory                   = "512"
  execution_role_arn       = module.ecs_task_execution_role.iam_role_arn
}

##########
#  権限   #
##########
data "aws_iam_policy" "ecs_task_execution_role_policy" {
  arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

data "aws_iam_policy_document" "ecs_task_execution" {
  source_json = data.aws_iam_policy.ecs_task_execution_role_policy.policy

  statement {
    effect    = "Allow"
    actions   = ["ssm:GetParameters", "kms:Decrypt"]
    resources = ["*"]
  }
}

module "ecs_task_execution_role" {
  source     = "./modules/iam_role"
  name       = "ecs-task-execution"
  identifier = "ecs-tasks.amazonaws.com"
  policy     = data.aws_iam_policy_document.ecs_task_execution.json
}
