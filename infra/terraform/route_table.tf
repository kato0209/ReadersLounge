##############
# RouteTable #
##############
resource "aws_route_table" "public_rtb" {
  vpc_id = aws_vpc.readerslounge.id

  route {
    gateway_id = aws_internet_gateway.readerslounge.id
    cidr_block = "0.0.0.0/0"
  }

  tags = {
    Name = "readerslounge-public-route"
  }
}
resource "aws_route_table" "private_rtb" {
  vpc_id = aws_vpc.readerslounge.id

  route {
    nat_gateway_id = aws_nat_gateway.readerslounge.id
    cidr_block     = "0.0.0.0/0"
  }

  tags = {
    Name = "readerslounge-private-route"
  }
}

###############
# association #
###############
resource "aws_route_table_association" "public_rtb_1a" {
  subnet_id      = aws_subnet.public_1a.id
  route_table_id = aws_route_table.public_rtb.id
}
resource "aws_route_table_association" "public_rtb_1c" {
  subnet_id      = aws_subnet.public_1c.id
  route_table_id = aws_route_table.public_rtb.id
}
resource "aws_route_table_association" "private_rtb_1a" {
  subnet_id      = aws_subnet.private_1a.id
  route_table_id = aws_route_table.private_rtb.id
}
resource "aws_route_table_association" "private_rtb_1c" {
  subnet_id      = aws_subnet.private_1c.id
  route_table_id = aws_route_table.private_rtb.id
}
