#################
# CloudWatchLog #
#################
resource "aws_cloudwatch_log_group" "api" {
  name              = "/ecs/api"
  retention_in_days = 180
}
resource "aws_cloudwatch_log_group" "front" {
  name              = "/ecs/front"
  retention_in_days = 180
}
resource "aws_cloudwatch_log_group" "db-create" {
  name              = "/ecs/db-create"
  retention_in_days = 180
}
resource "aws_cloudwatch_log_group" "db-migrate" {
  name              = "/ecs/db-migrate"
  retention_in_days = 180
}
resource "aws_cloudwatch_log_group" "db-migrate-reset" {
  name              = "/ecs/db-migrate-reset"
  retention_in_days = 180
}
