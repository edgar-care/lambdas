variable "aws_region" {
    description = "AWS region for all resources."

    type    = string
    default = "eu-west-3"
}

variable "aws_profile" {
    description = "AWS Profile to use"

    type = string
    default = "edgar.care"
}
