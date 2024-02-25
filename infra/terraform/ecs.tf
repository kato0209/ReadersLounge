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
  container_definitions = jsonencode([
    {
      name      = "front-container"
      image     = "620958051842.dkr.ecr.ap-northeast-1.amazonaws.com/readerslounge-front:latest"
      essential = true
      logConfiguration = {
        logDriver = "awslogs"
        options = {
          "awslogs-region"        = "ap-northeast-1"
          "awslogs-stream-prefix" = "front"
          "awslogs-group"         = "/ecs/front"
        }
      }
      portMappings = [
        {
          protocol      = "tcp"
          containerPort = 80
        }
      ]
    }
  ])
  execution_role_arn = aws_iam_role.ecs_task_execution.arn
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

  container_definitions = jsonencode([
    {
      name      = "api-container"
      image     = "620958051842.dkr.ecr.ap-northeast-1.amazonaws.com/readerslounge-api:latest"
      essential = true
      logConfiguration = {
        logDriver = "awslogs"
        options = {
          "awslogs-region"        = "ap-northeast-1"
          "awslogs-stream-prefix" = "api"
          "awslogs-group"         = "/ecs/api"
        }
      }
      portMappings = [
        {
          protocol      = "tcp"
          containerPort = 8080
          hostPort      = 8080
        }
      ]
      environment = [
        {
          name  = "GO_ENV",
          value = "prod"
        },
        {
          name  = "PGHOST",
          value = aws_db_instance.readerslounge.address
        },
        {
          name  = "PGSSLMODE",
          value = "disable"
        },
        {
          name  = "PGPORT",
          value = "5432"
        },
        {
          name  = "API_PROTOCOL",
          value = "https"
        },
        {
          name  = "API_DOMAIN",
          value = "readerslounge-server.com"
        },
        {
          name  = "API_PORT",
          value = "8080"
        },
        {
          name  = "FE_URL",
          value = "https://readerslounge-server.com"
        },
        {
          name  = "GOOGLE_OAUTH_PATH",
          value = "oauth/google/callback"
        },
        {
          name  = "GOOGLE_OAUTH_USER_INFO_EMAIL_URL",
          value = "https://www.googleapis.com/auth/userinfo.email"
        },
        {
          name  = "GOOGLE_OAUTH_USER_INFO_PROFILE_URL",
          value = "https://www.googleapis.com/auth/userinfo.profile"
        },
        {
          name  = "RAKUTEN_BOOKS_API_URL",
          value = "https://app.rakuten.co.jp/services/api/BooksBook/Search/20170404"
        },
        {
          name  = "RAKUTEN_BOOKS_GENRE_API_URL",
          value = "https://app.rakuten.co.jp/services/api/BooksGenre/Search/20121128"
        },
        {
          name  = "PGDATABASE"
          value = var.PGDATABASE
        },
        {
          name  = "PGUSER"
          value = var.PGUSER
        },
        {
          name  = "PGPASSWORD"
          value = var.PGPASSWORD
        },
        {
          name  = "JWT_SECRET"
          value = var.JWT_SECRET
        },
        {
          name  = "GOOGLE_CLIENT_ID"
          value = var.GOOGLE_CLIENT_ID
        },
        {
          name  = "GOOGLE_CLIENT_SECRET"
          value = var.GOOGLE_CLIENT_SECRET
        },
        {
          name  = "RAKUTEN_APPLICATION_ID"
          value = var.RAKUTEN_APPLICATION_ID
        },
        {
          name  = "AWS_REGION"
          value = var.AWS_REGION
        },
        {
          name  = "AWS_ACCESS_KEY_ID"
          value = var.AWS_ACCESS_KEY_ID
        },
        {
          name  = "AWS_SECRET_ACCESS_KEY"
          value = var.AWS_SECRET_ACCESS_KEY
        },
        {
          name  = "S3_BUCKET_NAME"
          value = var.S3_BUCKET_NAME
        }
      ]
    }
  ])
  execution_role_arn = aws_iam_role.ecs_task_execution.arn
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

######################################
#  マイグレーション task definition   #
######################################

resource "aws_ecs_task_definition" "goose_migration" {
  family                   = "readerslounge-goose-migration-task"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = "256"
  memory                   = "512"

  container_definitions = jsonencode([
    {
      name      = "goose-migration-container"
      image     = "620958051842.dkr.ecr.ap-northeast-1.amazonaws.com/readerslounge-migration:latest"
      essential = true
      cpu       = 256
      memory    = 256

      environment = [
        {
          name  = "GOOSE_DRIVER",
          value = "postgres"
        },
        {
          name  = "GOOSE_DBSTRING",
          value = "host=${aws_db_instance.readerslounge.address} user=${var.PGUSER} dbname=${var.PGDATABASE} password=${var.PGPASSWORD}"
        }
      ]

      logConfiguration = {
        logDriver = "awslogs"
        options = {
          "awslogs-region"        = "ap-northeast-1"
          "awslogs-stream-prefix" = "migration"
          "awslogs-group"         = "/ecs/migration"
        }
      }

      command = ["up"]
    }
  ])
  execution_role_arn = aws_iam_role.ecs_task_execution.arn
  task_role_arn      = aws_iam_role.ecs_task_role.arn
}

##########
#  権限   #
##########
resource "aws_iam_role" "ecs_task_execution" {
  name                = "ecs_task_execution"
  managed_policy_arns = [aws_iam_policy.ecs_task_execution.arn]

  assume_role_policy = jsonencode({
    Statement = [
      {
        Effect = "Allow"
        "Action" : [
          "sts:AssumeRole",
        ],
        Sid = ""
        Principal = {
          Service = "ecs-tasks.amazonaws.com"
        }
      },
    ]
  })
}

resource "aws_iam_policy" "ecs_task_execution" {
  name = "ecs_task_execution_policy"
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "ecr:GetAuthorizationToken",
          "ecr:BatchCheckLayerAvailability",
          "ecr:GetDownloadUrlForLayer",
          "ecr:BatchGetImage",
          "logs:CreateLogStream",
          "logs:PutLogEvents",
          "ecr:GetAuthorizationToken",
          "ecr:BatchCheckLayerAvailability",
          "ecr:GetDownloadUrlForLayer",
          "ecr:GetRepositoryPolicy",
          "ecr:DescribeRepositories",
          "ecr:ListImages",
          "ecr:DescribeImages",
          "ecr:BatchGetImage",
          "ecr:GetLifecyclePolicy",
          "ecr:GetLifecyclePolicyPreview",
          "ecr:ListTagsForResource",
          "ecr:DescribeImageScanFindings"
        ]
        Effect   = "Allow"
        Resource = "*"
      },
    ]
  })
}

resource "aws_iam_role" "ecs_task_role" {
  name                = "ecs_task_role"
  managed_policy_arns = [aws_iam_policy.ecs_task_role.arn]
  assume_role_policy = jsonencode({
    Statement = [
      {
        Effect = "Allow"
        "Action" : [
          "sts:AssumeRole",
        ],
        Sid = ""
        Principal = {
          Service = "ecs-tasks.amazonaws.com"
        }
      },
    ]
  })
}

resource "aws_iam_policy" "ecs_task_role" {
  name = "ecs_task_role_policy"
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "ssmmessages:CreateDataChannel",
          "ssmmessages:OpenDataChannel",
          "ssmmessages:OpenControlChannel",
          "ssmmessages:CreateControlChannel"
        ]
        Effect   = "Allow"
        Resource = "*"
      },
    ]
  })
}
