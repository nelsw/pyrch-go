package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"pyrch-go/pkg/model"
	data "pyrch-go/test"
	"testing"
)

func TestSave200(t *testing.T) {
	b, _ := json.Marshal(model.Credential{
		Id:       model.Id{"test"},
		Password: "test",
		UserId:   "test",
	})
	if out, _ := Handle(events.APIGatewayProxyRequest{
		Headers: map[string]string{"Authorization": data.TKN},
		Path:    "save",
		PathParameters: map[string]string{
			"table": "credential",
		},
		Body: string(b),
	}); out.StatusCode != 200 {
		t.Fail()
	}
}

func TestFindOne200(t *testing.T) {
	if out, _ := Handle(events.APIGatewayProxyRequest{
		Headers: map[string]string{"Authorization": data.TKN},
		Path:    "find-one",
		PathParameters: map[string]string{
			"table": "credential",
			"id":    "test",
		},
	}); out.StatusCode != 200 {
		t.Fail()
	}
}

// for code coverage purposes only
func TestHandleMain(t *testing.T) {
	go main()
}
