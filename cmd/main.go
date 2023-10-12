package main

import (
	"go-serverless/pkg/handlers"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var (
	dynaClient dynamodbiface.DynamoDBAPI
)

func main() {
	region := os.Getenv("AWS_REGION")
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region)})

	if err != nil {
		return
	}
	dynaClient = dynamodb.New(awsSession)
	lambda.Start(handler)
}

const tableName = "go-serverless"

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "GET":
		switch req.Path {
		case "/":
			return handlers.GetUser(req, tableName, dynaClient)
		case "/todos":
			return handlers.FetchTODOItemsByUser(req, tableName, dynaClient)
		default:
			return handlers.UnhandledMethod()
		}
	case "POST":
		switch req.Path {
		case "/":
			return handlers.CreateUser(req, tableName, dynaClient)
		case "/todos":
			return handlers.CreateOrUpdateTODOList(req, tableName, dynaClient)
		default:
			return handlers.UnhandledMethod()
		}
	case "PUT":
		return handlers.UpdateUser(req, tableName, dynaClient)
	case "DELETE":
		return handlers.DeleteUser(req, tableName, dynaClient)
	default:
		return handlers.UnhandledMethod()
	}
}
