resource "aws_instance" "qdapi_instance" {
  ami           = var.ami
  instance_type = var.aws_instance_type

  key_name = var.instance_key_name

  tags = {
    Name = var.instance_tag_name
  }

  user_data = file("setup.sh")
}


resource "aws_security_group" "qdapi_sg" {
  name = var.aws_security_group_name

  # SSH
  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  # HTTP
  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  # Custom TCP port
  ingress {
    from_port   = var.custom_tcp_port
    to_port     = var.custom_tcp_port
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}
