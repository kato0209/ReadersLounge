#########
#  eip  #
#########
resource "aws_eip" "readerslounge_ngw" {
  vpc        = true
  depends_on = [aws_internet_gateway.readerslounge]

  tags = {
    Name = "readerslounge-ngw-eip"
  }
}
