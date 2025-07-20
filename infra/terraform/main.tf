terraform {
  required_providers {
    aws = { source = "hashicorp/aws" }
  }
}

provider "aws" {
  region = "eu-central-1"
}

resource "aws_apigatewayv2_api" "termin" {
  name          = "terminjetzt-${var.stage}"
  protocol_type = "HTTP"
}

output "api_url" {
  value = aws_apigatewayv2_api.termin.api_endpoint
}
