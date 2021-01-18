package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gocolly/colly"
	"pyrch-go/internal/apigwp"
	"pyrch-go/internal/faas"
	"pyrch-go/pkg/model"
)

func Handle(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	apigwp.LogRequest(r)

	if r.Path == "crawl" {

		token := r.Headers["Authorization"]

		if res := faas.CallIt("token", "verify", r.Headers); res.StatusCode != 200 {
			return apigwp.NotOk(res.StatusCode, errors.New(res.Body))
		} else {
			token = res.Headers["Authorization"]
		}

		c := colly.NewCollector(
			colly.AllowedDomains("tampabaycichlids.com"),
		)

		c.OnHTML("div.small--one-half:nth-child(n)", func(e *colly.HTMLElement) {
			e1 := e.DOM.Get(0)
			id := "https://tampabaycichlids.com" + e1.FirstChild.NextSibling.Attr[0].Val
			fish := model.Fish{model.Id{id}, model.Animal{
				Url:    id,
				Title:  e1.FirstChild.NextSibling.FirstChild.NextSibling.Attr[1].Val,
				Images: []string{"https://" + e1.FirstChild.NextSibling.FirstChild.NextSibling.Attr[0].Val}},
			}
			b, _ := json.Marshal(&fish)
			faas.InvokeIt("repo", events.APIGatewayProxyRequest{
				Resource:                        "",
				Path:                            "save",
				HTTPMethod:                      "",
				Headers:                         r.Headers,
				MultiValueHeaders:               nil,
				QueryStringParameters:           nil,
				MultiValueQueryStringParameters: nil,
				PathParameters:                  map[string]string{"table": "fish"},
				StageVariables:                  nil,
				RequestContext:                  events.APIGatewayProxyRequestContext{},
				Body:                            string(b),
				IsBase64Encoded:                 false,
			})
		})

		// Before making a request print "Visiting ..."
		c.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting", r.URL.String())
		})

		c.OnResponse(func(f *colly.Response) {
			fmt.Println("Visited", f.Request.URL.String())
		})

		_ = c.Visit("https://tampabaycichlids.com/collections/mbuna")

		return apigwp.OkVoid(token)
	}

	return faas.InvokeIt("repo", r), nil

}

func main() {
	lambda.Start(Handle)
}
