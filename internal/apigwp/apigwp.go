// Response returns an API Gateway Proxy Response with a nil error to provide detailed status codes and response bodies.
// While a status code must be provided, further arguments are recognized with reflection but not required.
package apigwp

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

func NotOk(code int, err error) (events.APIGatewayProxyResponse, error) {
	return response(code, "", err.Error()), nil
}

func Bad(err error) (events.APIGatewayProxyResponse, error) {
	return response(400, "", err.Error()), nil
}

func BadRequest(s string) (events.APIGatewayProxyResponse, error) {
	return response(400, "", s), nil
}

func Unauthorized(err error) (events.APIGatewayProxyResponse, error) {
	return response(401, "", err.Error()), nil
}

func Ok(body interface{}) (events.APIGatewayProxyResponse, error) {
	return OkInterface("", body)
}

func OkVoid(token string) (events.APIGatewayProxyResponse, error) {
	return response(200, token, ""), nil
}

func OkInterface(token string, body interface{}) (events.APIGatewayProxyResponse, error) {
	if body == nil {
		return response(200, token, ""), nil
	}
	b, _ := json.Marshal(body)
	return response(200, token, string(b)), nil
}

func LogRequest(r events.APIGatewayProxyRequest) {
	fmt.Printf("request: {\n"+
		"\tmethod: %s\n"+
		"\tresource: %s\n"+
		"\tpath: %s\n"+
		"\theaders: %v\n"+
		"\tquery_string_parameters: %v\n"+
		"\tbody: %s\n"+
		"\tbase64: %v\n"+
		"}\n", r.HTTPMethod, r.Resource, r.Path, r.Headers, r.QueryStringParameters, r.Body, r.IsBase64Encoded)
}

func response(s int, t, b string) events.APIGatewayProxyResponse {
	fmt.Printf("response: {\n\tcode: %v\n\tbody: %s\n}\n", s, b)
	return events.APIGatewayProxyResponse{
		StatusCode: s,
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Authorization":               t,
		},
		MultiValueHeaders: nil,
		Body:              b,
		IsBase64Encoded:   false,
	}
}
