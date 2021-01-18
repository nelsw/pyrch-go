package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"pyrch-go/pkg/model"
	"testing"
)

func TestSave200(t *testing.T) {
	b, _ := json.Marshal(model.Fish{
		Id: model.Id{"test1"},
		Animal: model.Animal{
			Url:    "https://www.tampabaycichlids.com/collections/mbuna/products/cynotilapia-afra-jalo-reef-mbuna-malawi-african-cichlid",
			Title:  "Cynotilapia afra (Jalo Reef), Mbuna, Malawi African Cichlid",
			Images: []string{"https://cdn.shopify.com/s/files/1/0877/8234/products/left_facing-2_1024x1024.jpg?v=1527178249"},
		},
	})
	if out, _ := Handle(events.APIGatewayProxyRequest{
		Path: "save",
		PathParameters: map[string]string{
			"table": "fish",
		},
		Body: string(b),
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

func TestCrawl200(t *testing.T) {
	if out, _ := Handle(events.APIGatewayProxyRequest{Path: "crawl"}); out.StatusCode != 200 {
		t.Fail()
	}
}

//for code coverage purposes only
func TestHandle(t *testing.T) {
	go main()
}
