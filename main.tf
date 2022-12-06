terraform {
    required_providers {
        aws = {
        source  = "hashicorp/aws"
        version = ">= 4.9.0"
        }
        random = {
        source  = "hashicorp/random"
        version = "~> 3.1.0"
        }
        archive = {
        source  = "hashicorp/archive"
        version = "~> 2.2.0"
        }
    }

    required_version = "~> 1.0"
}

provider "aws" {
    region = var.aws_region
    profile = var.aws_profile
}


module "log_group" {
    source  = "terraform-aws-modules/cloudwatch/aws//modules/log-group"
    version = "~> 3.0"

    name              = "gateway"
    retention_in_days = 120
}


module "api_gateway" {
    source = "terraform-aws-modules/apigateway-v2/aws"

    name          = "edgar.care"
    description   = "My awesome HTTP API Gateway"
    protocol_type = "HTTP"

    cors_configuration = {
        allow_headers = ["*"]
        allow_methods = ["*"]
        allow_origins = ["*"]
    }

    create_api_domain_name = false # to control creation of API Gateway Domain Name

    # Access logs
    default_stage_access_log_destination_arn = module.log_group.cloudwatch_log_group_arn
    default_stage_access_log_format = jsonencode({
        requestId               = "$context.requestId"
        sourceIp                = "$context.identity.sourceIp"
        requestTime             = "$context.requestTime"
        protocol                = "$context.protocol"
        httpMethod              = "$context.httpMethod"
        resourcePath            = "$context.resourcePath"
        routeKey                = "$context.routeKey"
        status                  = "$context.status"
        responseLength          = "$context.responseLength"
        integrationErrorMessage = "$context.integrationErrorMessage"
        }
    )

    # Routes and integrations
    integrations = {
        "GET /graphql" = {
            lambda_arn             = module.graphql_lambda.lambda_function_arn
            payload_format_version = "2.0"
            timeout_milliseconds   = 12000
        }

        "POST /graphql" = {
            lambda_arn = module.graphql_lambda.lambda_function_arn
            payload_format_version = "2.0"
            timeout_milliseconds = 12000
        }

        # "GET /ping" = {
        #     lambda_arn             = module.ping.lambda_function_arn
        #     payload_format_version = "2.0"
        #     timeout_milliseconds   = 12000
        # }

        # "GET /some-route-with-authorizer" = {
        #   integration_type = "HTTP_PROXY"
        #   integration_uri  = "some url"
        #   authorizer_key   = "azure"
        # }

        # "$default" = {
        #   lambda_arn = "arn:aws:lambda:eu-west-1:052235179155:function:my-default-function"
        # }
    }

    # authorizers = {
    #   "azure" = {
    #     authorizer_type  = "JWT"
    #     identity_sources = "$request.header.Authorization"
    #     name             = "azure-auth"
    #     audience         = ["d6a38afd-45d6-4874-d1aa-3c5c558aqcc2"]
    #     issuer           = "https://sts.windows.net/aaee026e-8f37-410e-8869-72d9154873e4/"
    #   }
    # }
}

module "graphql_lambda" {
    source = "terraform-aws-modules/lambda/aws"

    function_name = "graphql"
    description   = "Go GraphQL api"
    handler       = "main"
    runtime       = "go1.x"
    publish       = true


    source_path = [
		{
			path = "${path.module}/graphql"
            commands = [ "make install", "make build", ":zip" ]
			patterns = [
				"!.*", // * Exclude everything
				"main",
                ".env",
				"schema.graphql",
			]
		}
	]

    allowed_triggers = {
        AllowExecutionFromAPIGateway = {
            service    = "apigateway"
            source_arn = "${module.api_gateway.apigatewayv2_api_execution_arn}/*/*"
        }
    }
}