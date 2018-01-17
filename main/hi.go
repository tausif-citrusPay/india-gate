package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HelloHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	response := "hi" + request.Body
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       response,
		Headers: map[string]string{
			"Content-Type": "text/plain",
		},
	}, nil

}


func main() {
	lambda.Start(HelloHandler)
}