package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"pyrch-go/pkg/model"
	"testing"
	"time"
)

var pp = map[string]string{
	"table": "user",
	"id":    "test",
}

var body = ""

func init() {

	id := model.Id{"test"}
	moment := model.Moment{time.Now().Unix()}

	m := model.User{id, moment}

	b, _ := json.Marshal(&m)
	body = string(b)
}

func TestFindOne200(t *testing.T) {

	r := events.APIGatewayProxyRequest{
		Path:           "find-one",
		PathParameters: pp,
	}

	if out, _ := Handle(r); out.StatusCode != 200 {
		t.Fail()
	}
}

func TestSaveOne200(t *testing.T) {

	r := events.APIGatewayProxyRequest{
		Path:           "save",
		PathParameters: pp,
		Body:           body,
	}

	if out, _ := Handle(r); out.StatusCode != 200 {
		t.Fail()
	}
}

// for code coverage purposes only
func TestHandleMain(t *testing.T) {
	go main()
}
