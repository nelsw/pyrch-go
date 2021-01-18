package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"pyrch-go/internal/apigwp"
	"pyrch-go/internal/faas"
)

func Handle(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	apigwp.LogRequest(r)

	if r.Path == "find-one" {
		return faas.InvokeIt("repo", events.APIGatewayProxyRequest{
			Path: "find-one",
			PathParameters: map[string]string{
				"table": "credential",
				"id":    r.PathParameters["id"],
			},
		}), nil
	}

	if r.Path == "save" {
		return faas.InvokeIt("repo", r), nil
	}

	return apigwp.Bad(fmt.Errorf("no path [%s]\n", r.Path))
}

func main() {
	lambda.Start(Handle)
}
