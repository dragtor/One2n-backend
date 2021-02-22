resource "aws_security_group" "sshaccess" {
  name        = "sshaccess"
  description = "Allow SSH from Anywhere"
  vpc_id      = module.vpc.vpc_id
  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  tags = {
    Name = "ssh access"
  }
}

resource "aws_security_group" "httpaccess" {
  name        = "httpaccess"
  description = "Allow HTTP from Anywhere"
  vpc_id      = module.vpc.vpc_id
  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  tags = {
    Name = "http access"
  }
}

resource "aws_security_group" "httpsaccess" {
  name        = "httpsaccess"
  description = "Allow HTTPS from Anywhere"
  vpc_id      = module.vpc.vpc_id
  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  tags = {
    Name = "http access"
  }
}

output "sg-ssh-access" {
  value = aws_security_group.sshaccess.id
}

output "sg-http-access" {
  value = aws_security_group.httpaccess.id
}

output "sg-https-access" {
  value = aws_security_group.httpsaccess.id
}
