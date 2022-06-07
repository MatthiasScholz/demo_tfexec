output "secret" {
  value       = aws_secretsmanager_secret.secret.name
  description = "Name of the created secret store."
}

output "secret_arn" {
  value       = aws_secretsmanager_secret.secret.arn
  description = "ARN of the created secret store."
}
