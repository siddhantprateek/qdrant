variable "aws_region" {
  description = "Default region for provider"
  type        = string
  default     = "ap-south-1"
}

variable "ami" {
  description = "Amazon machine image to use on EC2 instance"
  type        = string
  default     = "ami-0fa377108253bf620" // Ubuntu 22.04
}

variable "aws_instance_type" {
  description = "Amazon EC2 instance type"
  type        = string
  default     = "t2.micro"
}

variable "custom_tcp_port" {
  description = "Custom TCP port"
  type        = number
  default     = 8080
}

variable "instance_key_name" {
  description = "Instance Key name"
  type        = string
  default     = "qdapi"
}

variable "instance_tag_name" {
  description = "Instance tag name"
  type        = string
  default     = "QDAPI_INSTANCE"
}


variable "aws_security_group_name" {
  description = "Security group for the instance"
  type        = string
  default     = "qdapi_sg"
}
