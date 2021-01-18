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
		Path: "find-one",
		PathParameters: map[string]string{
			"table": "user",
			"id":    "test",
		},
	}); out.StatusCode != 200 {
		t.Fail()
	}
}

func TestFindAll200(t *testing.T) {
	if out, _ := Handle(events.APIGatewayProxyRequest{
		Path: "find-all",
		PathParameters: map[string]string{
			"table": "fish",
		},
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
		PathParameters: map[string]string{
			"table": "user",
		},
		Body: string(b),
	}); out.StatusCode != 200 {
		t.Fail()
	}
}

func TestHandleMain(t *testing.T) {
	go main()
}
