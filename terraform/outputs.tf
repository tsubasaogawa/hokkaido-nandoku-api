output "api_gateway_invoke_url" {
  description = "The invoke URL for the API Gateway stage."
  value       = aws_apigatewayv2_stage.this.invoke_url
}

output "cloudfront_distribution_domain_name" {
  description = "The domain name of the CloudFront distribution."
  value       = aws_cloudfront_distribution.this.domain_name
}

