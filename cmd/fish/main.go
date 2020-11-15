package main

import (
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gocolly/colly"
	"pyrch-go/internal/apigwp"
	"pyrch-go/internal/client/faas/client"
	"pyrch-go/pkg/model"
	"regexp"
)

type Entity struct {
	Id         string   `json:"id"`
	Owner      string   `json:"owner"`
	Name       string   `json:"name"`
	Common     string   `json:"common"`
	Scientific string   `json:"scientific"`
	Summary    string   `json:"summary"`
	Images     []Image  `json:"images"`
	Options    []Option `json:"options"`
	model.Taxonomy
}

type Image struct {
	Id       string `json:"id"`
	EntityId string `json:"entity_id"`
	Url      string `json:"url"`
	Name     string `json:"name"`
}

type Option struct {
	Id       string   `json:"id"`
	EntityId string   `json:"entity_id"`
	Price    int64    `json:"price"`   // 7900 = $79.00, stripe thinks it makes cents
	Weight   int      `json:"weight"`  // 170 = 1.7, to avoid decimals entirely
	Label    string   `json:"label"`   // oz, lb, kilo, ton, w/e
	Stock    int      `json:"stock"`   // quantity available
	Address  string   `json:"address"` // shipping departure location
	Images   []string `json:"images"`  // urls
}

func (e *Entity) Validate() error {
	return nil
}

var (
	allPaths = regexp.MustCompile(`crawl|find`)
	jwtPaths = regexp.MustCompile(`crawl`)
)

func Handle(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	apigwp.LogRequest(r)

	token := r.Headers["Authorize"]

	// is the path valid?
	if !allPaths.Match([]byte(r.Path)) {
		return apigwp.BadRequest(fmt.Errorf("path not found [%s]", r.Path))
	}

	// is token validation required?
	if jwtPaths.Match([]byte(r.Path)) {
		res := client.CallIt("token", "verify", r.Headers)
		if res.StatusCode != 200 {
			return apigwp.NotOk(res.StatusCode, errors.New(res.Body))
		}
		token = res.Headers["Authorize"]
	}

	if r.Path == "crawl" {
		// Instantiate default collector
		c := colly.NewCollector(
			// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
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

		// Start scraping on https://hackerspaces.org
		c.Visit("https://tampabaycichlids.com/collections/mbuna")

		return apigwp.Ok()
	}

	// else path == "find"

	return apigwp.OkVoid(token)
}

func main() {
	lambda.Start(Handle)
}
