package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

func ping() (string, error) {
	return ".", nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(ping)
}
