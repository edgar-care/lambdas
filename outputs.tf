output "base_url" {
    description = "Base URL for API Gateway stage."

    value = module.api_gateway.default_apigatewayv2_stage_invoke_url
}

# output "routes" {
#     description = "Deployed routes"

#     value = {
#         "/hello" = module.hello-world.lambda_function_name
#         "/ping"  = module.ping.lambda_function_name
#     }
# }