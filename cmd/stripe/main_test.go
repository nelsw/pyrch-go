package main

import (
	"github.com/aws/aws-lambda-go/events"
	"testing"
)

func TestHandleAmount(t *testing.T) {
	if out, _ := Handle(events.APIGatewayProxyRequest{
		QueryStringParameters: map[string]string{"amount": "2789"},
	}); out.StatusCode != 200 {
		t.Fatal(out)
	} else {
		t.Log(out)
	}
}

func TestHandle(t *testing.T) {
	go main()
}
