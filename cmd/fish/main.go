package main

import (
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gocolly/colly"
	"pyrch-go/internal/apigwp"
	"pyrch-go/internal/faas"
)

func Handle(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	apigwp.LogRequest(r)

	token := r.Headers["Authorize"]

	if r.Path == "crawl" {

		res := faas.CallIt("token", "verify", r.Headers)
		if res.StatusCode != 200 {
			return apigwp.NotOk(res.StatusCode, errors.New(res.Body))
		}
		token = res.Headers["Authorize"]

		c := colly.NewCollector(
			colly.AllowedDomains("tampabaycichlids.com"),
		)

		c.OnHTML("div.small--one-half:nth-child(n)", func(e *colly.HTMLElement) {

			e1 := e.DOM.Get(0)

			fmt.Println("https://tampabaycichlids.com" + e1.FirstChild.NextSibling.Attr[0].Val)
			fmt.Println("https://" + e1.FirstChild.NextSibling.FirstChild.NextSibling.Attr[0].Val)
			fmt.Println(e1.FirstChild.NextSibling.FirstChild.NextSibling.Attr[1].Val)

		})

		// Before making a request print "Visiting ..."
		c.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting", r.URL.String())
		})

		c.OnResponse(func(f *colly.Response) {
			fmt.Println("Visited", f.Request.URL.String())
		})

		c.Visit("https://tampabaycichlids.com/collections/mbuna")

		return apigwp.OkVoid(token)
	}

	// else path == "repo"

	return apigwp.Bad(fmt.Errorf("path not found [%s]", r.Path))
}

func main() {
	lambda.Start(Handle)
}
