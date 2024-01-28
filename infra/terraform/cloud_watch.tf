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
