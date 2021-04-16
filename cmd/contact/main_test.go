package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"testing"
)

func TestSave200(t *testing.T) {
	b, _ := json.Marshal(Contact{
		Name:    "John Doe",
		Email:   "John.Doe@gmail.com",
		Message: "Hello World!",
	})
	if out, _ := Handle(events.APIGatewayProxyRequest{
		Body: string(b),
	}); out.StatusCode != 200 {
		t.Fail()
	}
}

// for code coverage purposes only
func TestHandleMain(t *testing.T) {
	go main()
}
