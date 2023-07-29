terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = "ap-south-1"
}

resource "aws_instance" "ec2_instance" {
  ami           = "ami-0f5ee92e2d63afc18"
  instance_type = "t2.micro"
  key_name      = "qdrant-key-pair"

  tags = {
    Name = "qdrantEC2"
  }
}

resource "aws_security_group" "ec2_security_group" {
  name_prefix = "ec2-security-group"
}


resource "aws_security_group_rule" "ec2_ssh_inbound" {
  type              = "ingress"
  from_port         = 22
  to_port           = 22
  protocol          = "tcp"
  cidr_blocks       = ["0.0.0.0/0"]
  security_group_id = aws_security_group.ec2_security_group.id
}

resource "aws_security_group_rule" "ec2_https_inbound" {
  type              = "ingress"
  from_port         = 433
  to_port           = 433
  protocol          = "tcp"
  cidr_blocks       = ["0.0.0.0/0"]
  security_group_id = aws_security_group.ec2_security_group.id
}

resource "aws_security_group_rule" "ec2_http_inbound" {
  type              = "ingress"
  from_port         = 80
  to_port           = 80
  protocol          = "tcp"
  cidr_blocks       = ["0.0.0.0/0"]
  security_group_id = aws_security_group.ec2_security_group.id
}

# resource "aws_security_group" "allow_http_https" {
#   name_prefix = "allow_http_https"
#   ingress {
#     from_port   = 80
#     to_port     = 80
#     protocol    = "tcp"
#     cidr_blocks = ["0.0.0.0/0"]
#   }
#   ingress {
#     from_port   = 443
#     to_port     = 443
#     protocol    = "tcp"
#     cidr_blocks = ["0.0.0.0/0"]
#   }
# }
