package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/edgar-care/graphql/cmd/main/database"
	edGraphql "github.com/edgar-care/graphql/cmd/main/graphql"
	"github.com/edgar-care/graphql/cmd/main/lib"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/joho/godotenv"
)

var schema *graphql.Schema

var (
	ErrQueryNameNotProvided = errors.New("no query was provided in the HTTP body")
)

func Handler(context context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)

	if len(request.Body) < 1 {
		return events.APIGatewayProxyResponse{
			Body:       string(lib.ErrorMarshal("No query was provided in the HTTP body")),
			StatusCode: 200,
		}, nil
	}

	var params struct {
		Query         string                 `json:"query"`
		OperationName string                 `json:"operationName"`
		Variables     map[string]interface{} `json:"variables"`
	}

	if err := json.Unmarshal([]byte(request.Body), &params); err != nil {
		log.Print("Could not decode body", err)
	}

	response := schema.Exec(context, params.Query, params.OperationName, params.Variables)
	responseJSON, err := json.Marshal(response)
	if err != nil {
		log.Print("Could not decode body")
	}

	return events.APIGatewayProxyResponse{
		Body:       string(responseJSON),
		StatusCode: 200,
	}, nil

}

func init() {
	err := godotenv.Load()
	lib.CheckError(err)

	rawSchema, err := os.ReadFile("schema.graphql")
	lib.CheckError(err)

	schema = graphql.MustParseSchema(string(rawSchema), &edGraphql.Resolver{})
}

func main() {
	_, present := os.LookupEnv("AWS_LAMBDA_FUNCTION_NAME")
	db := database.Connect(os.Getenv("DATABASE_URL"))
	edGraphql.Init(db)

	if !present {
		http.Handle("/graphql", &relay.Handler{Schema: schema})
		log.Print("Starting to listen 8080...")
		log.Fatal(http.ListenAndServe(":8080", nil))
	} else {
		lambda.Start(Handler)
	}
}
