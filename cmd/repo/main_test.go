package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"pyrch-go/pkg/model"
	data "pyrch-go/test"
	"testing"
)

func TestFindOne200(t *testing.T) {
	if out, _ := Handle(events.APIGatewayProxyRequest{
		Headers: map[string]string{"Authorization": data.TKN},
		Path:    "find-one",
		PathParameters: map[string]string{
			"table": "profile",
			"id":    "test",
		},
	}); out.StatusCode != 200 {
		t.Fail()
	}
}

func TestFindAll200(t *testing.T) {
	if out, _ := Handle(events.APIGatewayProxyRequest{
		Headers: map[string]string{"Authorization": data.TKN},
		Path:    "find-all",
		PathParameters: map[string]string{
			"table": "profile",
		},
	}); out.StatusCode != 200 {
		t.Fail()
	}
}

func TestSaveOne200(t *testing.T) {

	m := model.Moment{}
	_ = m.Validate()

	b, _ := json.Marshal(struct {
		Id     string `json:"id"`
		UserId string `json:"user_id"`
		model.Moment
	}{
		"test",
		"test",
		m,
	})
	if out, _ := Handle(events.APIGatewayProxyRequest{
		Headers: map[string]string{"Authorization": data.TKN},
		Path:    "save",
		PathParameters: map[string]string{
			"table": "profile",
		},
		Body: string(b),
	}); out.StatusCode != 200 {
		t.Fail()
	}
}

func TestHandleMain(t *testing.T) {
	go main()
}
