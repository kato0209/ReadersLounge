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
      image     = "620958051842.dkr.ecr.ap-northeast-1.amazonaws.com/front:latest"
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
      command = [
        "npm", "run", "build"
      ]
      environment = [
        {
          name  = "VITE_API_URL"
          value = "https://readerslounge-server.com:8080"
        },
        {
          name  = "VITE_GOOGLE_OAUTH_REDIRECT_PATH"
          value = "/oauth/google/callback"
        },
        {
          name  = "VITE_GOOGLE_OAUTH_USER_INFO_EMAIL_URL"
          value = "https://www.googleapis.com/auth/userinfo.email"
        },
        {
          name  = "VITE_GOOGLE_OAUTH_USER_INFO_PROFILE_URL"
          value = "https://www.googleapis.com/auth/userinfo.profile"
        },
        {
          name  = "VITE_WEBSOCKET_URL"
          value = "wss://readerslounge-server.com:8080"
        }
      ]
      secrets = [
        {
          name      = "VITE_GOOGLE_OAUTH_CLIENT_ID"
          valueFrom = "vite-google-oauth-client-id"
        },
        {
          name      = "VITE_GOOGLE_CLIENT_SECRET"
          valueFrom = "vite-google-client-secret"
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
  depends_on = [
    aws_lb.readerslounge,
    aws_lb_target_group.front,
  ]
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
      image     = "620958051842.dkr.ecr.ap-northeast-1.amazonaws.com/api:latest"
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
          value = aws_db_instance.readerslounge.endpoint
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
        }
      ]
      secrets = [
        {
          name      = "PGDATABASE",
          valueFrom = "pgdatabase"
        },
        {
          name      = "PGUSER",
          valueFrom = "pguser"
        },
        {
          name      = "PGPASSWORD",
          valueFrom = "pgpassword"
        },
        {
          name      = "JWT_SECRET",
          valueFrom = "jwt-secret"
        },
        {
          name      = "GOOGLE_CLIENT_ID",
          valueFrom = "google-client-id"
        },
        {
          name      = "GOOGLE_CLIENT_SECRET",
          valueFrom = "google-client-secret"
        },
        {
          name      = "RAKUTEN_APPLICATION_ID",
          valueFrom = "rakuten-application-id"
        },
        {
          name      = "AWS_DEFAULT_REGION",
          valueFrom = "aws_default_region"
        },
        {
          name      = "AWS_ACCESS_KEY_ID",
          valueFrom = "aws_access_key_id"
        },
        {
          name      = "AWS_SECRET_ACCESS_KEY",
          valueFrom = "aws_secret_access_key"
        },
        {
          name      = "S3_BUCKET_NAME",
          valueFrom = "s3_bucket_name"
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
  depends_on = [
    aws_lb.readerslounge,
    aws_lb_target_group.api,
  ]
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
