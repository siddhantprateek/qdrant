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

resource "aws_instance" "qdapi_instance" {
  ami           = "ami-0f5ee92e2d63afc18"
  instance_type = "t2.micro"
  key_name      = "qdapi"

  tags = {
    Name = "qdapi-ec2"
  }

  user_data = <<-EOF
              #!/bin/bash
              sudo apt-get update
              sudo snap install docker
              git clone https://github.com/siddhantprateek/qdrant.git
              cd qdrant
              sudo docker-compose -f docker-compose.prod.yaml up -d
              EOF
}
