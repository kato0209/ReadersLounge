#########
#  RDS  #
#########
resource "aws_db_instance" "readerslounge" {
  db_name                 = var.DB_NAME
  allocated_storage       = 20
  instance_class          = "db.t3.micro"
  engine                  = "postgres"
  engine_version          = "15.3"
  storage_type            = "gp2"
  username                = var.DB_USERNAME
  password                = var.DB_PASSWORD
  backup_retention_period = 7
  copy_tags_to_snapshot   = true
  max_allocated_storage   = 200
  skip_final_snapshot     = true
  vpc_security_group_ids  = [aws_security_group.rds.id]
  db_subnet_group_name    = aws_db_subnet_group.rds.name

  parameter_group_name = aws_db_parameter_group.readerslounge.name
  apply_immediately    = true

  lifecycle {
    prevent_destroy = false
  }

}

resource "aws_db_parameter_group" "readerslounge" {
  name   = "readerslounge-aws-db-parameter-group"
  family = "postgres15"

  parameter {
    name  = "timezone"
    value = "Asia/Tokyo"
  }

  parameter {
    name  = "rds.force_ssl"
    value = "0"
  }
}
