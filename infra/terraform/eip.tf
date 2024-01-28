#########
#  eip  #
#########
resource "aws_eip" "readerslounge_ngw" {
  domain     = "vpc"
  depends_on = [aws_internet_gateway.readerslounge]

  tags = {
    Name = "readerslounge-ngw-eip"
  }
}
