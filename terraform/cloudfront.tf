locals {
  # api_endpoint comes with https:// prefix, which needs to be removed for CloudFront origin domain
  api_gateway_domain = replace(aws_apigatewayv2_api.this.api_endpoint, "/^https?:///", "")
}

resource "aws_cloudfront_distribution" "this" {
  enabled = true
  comment = "${var.project_name} API distribution"

  origin {
    domain_name = local.api_gateway_domain
    origin_id   = "APIGateway"

    custom_origin_config {
      http_port              = 80
      https_port             = 443
      origin_protocol_policy = "https-only"
      origin_ssl_protocols   = ["TLSv1.2"]
    }
  }

  # PriceClass_100 includes: US, Canada, Europe
  # This is the most cost-effective class
  price_class = "PriceClass_100"

  web_acl_id = aws_wafv2_web_acl.cloudfront.arn

  default_cache_behavior {
    allowed_methods  = ["GET", "HEAD", "OPTIONS", "PUT", "POST", "PATCH", "DELETE"]
    cached_methods   = ["GET", "HEAD"]
    target_origin_id = "APIGateway"

    # CachingDisabled Managed Policy
    cache_policy_id = "4135ea2d-6df8-44a3-9db3-43551f2be087"

    viewer_protocol_policy = "redirect-to-https"
  }

  # Cache behavior for /list
  ordered_cache_behavior {
    path_pattern     = "/list"
    allowed_methods  = ["GET", "HEAD", "OPTIONS"]
    cached_methods   = ["GET", "HEAD"]
    target_origin_id = "APIGateway"

    # CachingOptimized Managed Policy
    cache_policy_id = "658327ea-f89d-4fab-a63d-7e88639e58f6"

    viewer_protocol_policy = "redirect-to-https"
  }

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }

  viewer_certificate {
    cloudfront_default_certificate = true
  }
}
