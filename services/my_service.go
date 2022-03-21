package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(context.Context, events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Print("This is printed as the http response :(")
	// fmt.Println("This wouldn't work too :(")

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       `{"message": "Hello serverless world"}`,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
