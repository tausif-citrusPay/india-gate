package main

import (
	"github.com/aws/aws-lambda-go/events"
)

func helloHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	response := "hi" + request.Body
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       response,
		Headers: map[string]string{
			"Content-Type": "text/plain",
		},
	}, nil

}


