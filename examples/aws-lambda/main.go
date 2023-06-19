package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ApiResponse := events.APIGatewayProxyResponse{}
	switch request.HTTPMethod {
	case "GET":
		ApiResponse = events.APIGatewayProxyResponse{Body: "HTTP GET", StatusCode: 200}

	case "POST":
		ApiResponse = events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 201}

	}
	return ApiResponse, nil

}

func main() {
	lambda.Start(HandleRequest)
}
