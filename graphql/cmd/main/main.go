package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	edgar_gql "github.com/edgar-care/edgarlib/graphql/server"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ginLambda *ginadapter.GinLambdaV2

func connect(dbUrl string) *edgar_gql.DB {
	log.Println("Connecting to " + dbUrl)
	client, err := mongo.NewClient(options.Client().ApplyURI(dbUrl))
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	tmp := edgar_gql.DB{Client: client}

	return &tmp
}

func graphqlHandler() gin.HandlerFunc {
	db := connect(os.Getenv("DATABASE_URL"))
	h := handler.NewDefaultServer(edgar_gql.NewExecutableSchema(edgar_gql.Config{Resolvers: &edgar_gql.Resolver{Db: db}}))

	return func(c *gin.Context) {

		if c.Request.Header.Get(os.Getenv("API_KEY")) != os.Getenv("API_KEY_VALUE") {
			c.JSON(http.StatusForbidden, gin.H{"error": "unauthorized"})
			return
		}
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	if ginLambda == nil {
		log.Printf("Gin cold start")
		r := gin.Default()
		r.GET("/graphql/playground", playgroundHandler())
		r.POST("/graphql/query", graphqlHandler())
		r.POST("/dev/graphql/query", graphqlHandler())
		r.POST("/demo/graphql/query", graphqlHandler())
		r.GET("/dev/graphql/playground", playgroundHandler())
		r.GET("/demo/graphql/playground", playgroundHandler())

		ginLambda = ginadapter.NewV2(r)
	}

	spew.Dump(req)

	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
