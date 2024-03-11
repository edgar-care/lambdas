package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	ID string `json:"id"`
}

func Handler(ctx context.Context, event Event) (string, error) {
	ID := event.ID

	result := fmt.Sprintf("Received parameter1: %s, parameter2: %d", ID)

	return result, nil
}

func main() {
	lambda.Start(Handler)
}
