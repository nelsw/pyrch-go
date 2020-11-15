package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"sam-app/pkg/client/faas/client"
	"sam-app/pkg/factory/apigwp"
	"strings"
)

func Handle(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	apigwp.LogRequest(r)
	if r.Path != "add" && r.Path != "delete" {
		return apigwp.Response(400, fmt.Errorf("bad path [%s]", r.Path))
	} else if _, ok := r.Headers["Authorize"]; !ok {
		return apigwp.Response(400, fmt.Errorf("no token"))
	} else if csv, ok := r.QueryStringParameters["ids"]; !ok {
		return apigwp.Response(400, fmt.Errorf("no ids"))
	} else if col, ok := r.QueryStringParameters["col"]; !ok {
		return apigwp.Response(400, fmt.Errorf("no col"))
	} else {

		authenticate := events.APIGatewayProxyRequest{Path: "authenticate", Headers: r.Headers}
		authResponse := client.Invoke("tokenHandler", authenticate)
		if authResponse.StatusCode != 200 {
			return apigwp.Response(401, authResponse.Body)
		} else {
			r.Headers = authResponse.Headers
		}

		claims := map[string]string{}
		_ = json.Unmarshal([]byte(authResponse.Body), &claims)

		keyword := r.Path + " " + col
		ids := strings.Split(csv, ",")
		m := map[string]interface{}{"table": "user", "id": claims["jti"], "ids": ids, "keyword": keyword, "type": "*user.Entity"}
		code, body := client.CallIt(&m, "repoHandler")
		return apigwp.ProxyResponse(code, r.Headers, body)
	}
}

func main() {
	lambda.Start(Handle)
}
