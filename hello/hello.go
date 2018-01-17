package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// HelloHandler is handler func
func HelloHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	response := "Hi " + request.Body
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       response,
		Headers: map[string]string{
			"Content-Type": "text/html",
		},
	}, nil

}

func main() {
	lambda.Start(HelloHandler)
}
