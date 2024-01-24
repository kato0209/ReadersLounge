##############
# NATGateway #
##############
resource "aws_nat_gateway" "readerslounge" {
  allocation_id = aws_eip.readerslounge_ngw.id
  subnet_id     = aws_subnet.public_1a.id
  depends_on    = [aws_internet_gateway.readerslounge]

  tags = {
    Name = "readerslounge-ngw"
  }
}
