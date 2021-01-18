package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"pyrch-go/pkg/model"
	"testing"
	"time"
)

func TestFindOne200(t *testing.T) {
	if out, _ := Handle(events.APIGatewayProxyRequest{
		Path:           "find-one",
		PathParameters: map[string]string{"id": "test"},
	}); out.StatusCode != 200 {
		t.Fail()
	}
}

func TestSaveOne200(t *testing.T) {
	b, _ := json.Marshal(&model.User{
		model.Id{"test"},
		model.Moment{time.Now().Unix()},
	})
	if out, _ := Handle(events.APIGatewayProxyRequest{
		Path: "save",
		Body: string(b),
	}); out.StatusCode != 200 {
		t.Fail()
	}
}

// for code coverage purposes only
func TestHandleMain(t *testing.T) {
	go main()
}
