data "aws_availability_zones" "available" {}

resource "aws_vpc" "main" {
  cidr_block = var.vpc_cidr

  tags = {
    Name = var.vpc_tag_name
  }
}


resource "aws_internet_gateway" "default" {
  vpc_id = aws_vpc.main.id

  tags = {
    Name = var.igw_tag_name
  }
}

locals {
  vpc_cidr_prefix = "${join(".", slice(split(".", var.vpc_cidr), 0, 2))}"
}

/*
  Public Subnet
*/

resource "aws_subnet" "public_subnet" {
  count                   = length(data.aws_availability_zones.available.names)
  vpc_id                  = aws_vpc.main.id
  cidr_block              = "${local.vpc_cidr_prefix}.${10 + count.index}.0/24"
  availability_zone       = data.aws_availability_zones.available.names[count.index]
  map_public_ip_on_launch = true
  tags = {
    Name = "PublicSubnet-${data.aws_availability_zones.available.names[count.index]}"
  }
}

/*
  Private Subnet
*/
resource "aws_subnet" "private_subnet" {
  count                   = length(data.aws_availability_zones.available.names)
  vpc_id                  = aws_vpc.main.id
  cidr_block              = "${local.vpc_cidr_prefix}.${20 + count.index}.0/24"
  availability_zone       = data.aws_availability_zones.available.names[count.index]
  map_public_ip_on_launch = false
  tags = {
    Name = "PrivateSubnet-${data.aws_availability_zones.available.names[count.index]}"
  }
}

/*
 Public Route Tables
*/

resource "aws_route_table" "rt-public" {
  vpc_id = aws_vpc.main.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.default.id
  }

  tags = {
    Name = "Public-rt"
  }
}

resource "aws_route_table_association" "public" {
  count          = length(aws_subnet.public_subnet.*.id)
  subnet_id      = element(aws_subnet.public_subnet.*.id, count.index)
  route_table_id = aws_route_table.rt-public.id
}

/*
 Private Route Tables
*/

resource "aws_route_table" "rt-private" {
  vpc_id = aws_vpc.main.id

  tags = {
    Name = "Private-rt"
  }
}

resource "aws_route_table_association" "private" {
  count          = length(aws_subnet.private_subnet.*.id)
  subnet_id      = element(aws_subnet.private_subnet.*.id, count.index)
  route_table_id = aws_route_table.rt-private.id
}

output "vpc_id" {
  value = aws_vpc.main.id
}

output "subnet_ids" {
  value = aws_subnet.public_subnet.*.id
}

output "private_subnet_ids" {
  value = aws_subnet.private_subnet.*.id
}
