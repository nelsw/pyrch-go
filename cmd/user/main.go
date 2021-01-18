package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"pyrch-go/internal/apigwp"
	"pyrch-go/internal/faas"
	"pyrch-go/pkg/model"
)

func Handle(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	apigwp.LogRequest(r)

	token := r.Headers["Authorize"]

	if res := faas.CallIt("token", "verify", r.Headers); res.StatusCode != 200 {
		return apigwp.NotOk(res.StatusCode, errors.New(res.Body))
	} else {
		token = res.Headers["Authorize"]
	}

	if r.Path == "find-one" {
		if id, ok := r.PathParameters["id"]; !ok {
			return apigwp.Bad(fmt.Errorf("no id"))
		} else {
			b, _ := json.Marshal(&model.User{})
			return faas.InvokeIt("repo", events.APIGatewayProxyRequest{
				Headers: map[string]string{"Authorize": token},
				Path:    r.Path,
				PathParameters: map[string]string{
					"table": "user",
					"id":    id,
				},
				Body: string(b),
			}), nil
		}
	}

	if r.Path == "save" {
		return faas.InvokeIt("repo", events.APIGatewayProxyRequest{
			Headers: map[string]string{"Authorize": token},
			Path:    r.Path,
			PathParameters: map[string]string{
				"table": "user",
			},
			Body: r.Body,
		}), nil
	}

	return apigwp.OkVoid(token)
}

func main() {
	lambda.Start(Handle)
}
