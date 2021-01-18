package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"pyrch-go/internal/apigwp"
)

func HandleRequest(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	apigwp.LogRequest(r)

	return apigwp.Bad(fmt.Errorf("no path [%s]\n", r.Path))

}

func main() {
	lambda.Start(HandleRequest)
}
