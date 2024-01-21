##########
# Subnet #
##########
resource "aws_subnet" "public_1a" {
  vpc_id                  = aws_vpc.readerslounge.id
  cidr_block              = "10.0.0.0/24"
  availability_zone       = "ap-northeast-1a"
  map_public_ip_on_launch = true

  tags = {
    Name = "readerslounge-public-1a"
  }
}
resource "aws_subnet" "public_1c" {
  vpc_id                  = aws_vpc.readerslounge.id
  cidr_block              = "10.0.1.0/24"
  availability_zone       = "ap-northeast-1c"
  map_public_ip_on_launch = true

  tags = {
    Name = "readerslounge-public-1c"
  }
}
resource "aws_subnet" "private_1a" {
  vpc_id                  = aws_vpc.readerslounge.id
  cidr_block              = "10.0.2.0/24"
  availability_zone       = "ap-northeast-1a"
  map_public_ip_on_launch = true

  tags = {
    Name = "readerslounge-private-1a"
  }
}
resource "aws_subnet" "private_1c" {
  vpc_id                  = aws_vpc.readerslounge.id
  cidr_block              = "10.0.3.0/24"
  availability_zone       = "ap-northeast-1c"
  map_public_ip_on_launch = true

  tags = {
    Name = "readerslounge-private-1c"
  }
}

###############
# SubnetGroup #
###############
resource "aws_db_subnet_group" "rds" {
  name        = "readerslounge-rds-subnet-group"
  description = "rds subnet for readerslounge"
  subnet_ids  = [aws_subnet.private_1a.id, aws_subnet.private_1c.id]
}
