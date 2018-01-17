package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func helloHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {


	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(request.Body),
		Headers: map[string]string{
			"Content-Type": "text/html",
		},
	}, nil

}

func main() {
	lambda.Start(helloHandler)
}

