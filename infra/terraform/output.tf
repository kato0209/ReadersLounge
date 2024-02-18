output "subnet_1a_id" {
  value = aws_subnet.public_1a.id
}

output "subnet_1c_id" {
  value = aws_subnet.public_1c.id
}

output "security_group_id" {
  value = aws_security_group.ecs_api.id
}
