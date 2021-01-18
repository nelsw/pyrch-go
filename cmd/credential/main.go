package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"pyrch-go/internal/apigwp"
	"pyrch-go/internal/faas"
)

func Handle(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	apigwp.LogRequest(r)

	return faas.InvokeIt("repo", r), nil
}

func main() {
	lambda.Start(Handle)
}
