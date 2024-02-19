output "instance_public_ip_address" {
  value = aws_instance.qdapi_instance.public_ip
}

output "instance_private_ip_address" {
  value = aws_instance.qdapi_instance.private_ip
}

output "instance_arn" {
  value = aws_instance.qdapi_instance.arn
}