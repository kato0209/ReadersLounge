###################
# InternetGateway #
###################
resource "aws_internet_gateway" "readerslounge" {
  vpc_id = aws_vpc.readerslounge.id

  tags = {
    Name = "readerslounge-igw"
  }
}
