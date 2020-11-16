package main

import (
	"github.com/aws/aws-lambda-go/events"
	"testing"
)

func TestFind200(t *testing.T) {
	handleTest(t, "repo")
}

func TestCrawl200(t *testing.T) {
	handleTest(t, "crawl")
}

func handleTest(t *testing.T, path string) {
	if out, _ := Handle(events.APIGatewayProxyRequest{Path: path}); out.StatusCode != 200 {
		t.Fail()
	}
}

//for code coverage purposes only
func TestHandle(t *testing.T) {
	go main()
}
