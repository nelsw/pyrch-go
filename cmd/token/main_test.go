package main

import (
	"github.com/aws/aws-lambda-go/events"
	"testing"
)

var headers map[string]string

func init() {
	out, _ := Handle(events.APIGatewayProxyRequest{Path: "create", Body: "test"})
	headers = out.Headers
}

func TestHandleCreate200(t *testing.T) {
	if out, _ := Handle(events.APIGatewayProxyRequest{Path: "create", Body: "test"}); out.StatusCode != 200 {
		t.Fail()
	}
}

func TestHandleVerify200(t *testing.T) {
	if out, _ := Handle(events.APIGatewayProxyRequest{Path: "verify", Headers: headers}); out.StatusCode != 200 {
		t.Fail()
	}
}

func TestHandleClaims200(t *testing.T) {
	if out, _ := Handle(events.APIGatewayProxyRequest{Path: "claims", Headers: headers}); out.StatusCode != 200 {
		t.Fail()
	}
}

// for code coverage purposes only
func TestHandle(t *testing.T) {
	go main()
}
