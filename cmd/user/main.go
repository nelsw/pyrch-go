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
	"regexp"
)

var paths = regexp.MustCompile(`save|find-one`)

func Handle(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	apigwp.LogRequest(r)

	token := r.Headers["Authorize"]

	// is path valid?
	if !paths.Match([]byte(r.Path)) {
		return apigwp.NotFound(fmt.Errorf("path not found [%s]", r.Path))
	}

	// is token valid?
	res := faas.CallIt("token", "verify", r.Headers)
	if res.StatusCode != 200 {
		return apigwp.NotOk(res.StatusCode, errors.New(res.Body))
	}
	token = res.Headers["Authorize"]

	if r.Path == "find-one" {
		if pk, ok := r.PathParameters["pk"]; !ok {
			return apigwp.Bad(fmt.Errorf("no pk"))
		} else {
			b, _ := json.Marshal(&model.User{})
			return faas.InvokeIt("repo", events.APIGatewayProxyRequest{
				Headers: map[string]string{"Authorize": token},
				Path:    r.Path,
				PathParameters: map[string]string{
					"table": "user",
					"pk":    pk,
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
