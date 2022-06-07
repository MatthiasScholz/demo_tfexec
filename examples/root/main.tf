# FIXME: Use customer managed key instead of AWS managed key, see:
#        https://aquasecurity.github.io/tfsec/v1.1.5/checks/aws/ssm/secret-use-customer-key/
# tfsec:ignore:aws-ssm-secret-use-customer-key
resource "aws_secretsmanager_secret" "secret" {
  name = "${var.environment}/${var.team}/${var.stack}/${var.name}"

  tags = {
    environment = var.environment
    iac         = var.repository
    stack       = var.stack
    team        = var.team
  }
}
