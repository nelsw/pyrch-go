package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"pyrch-go/internal/apigwp"
	"strings"
)

const (
	ValidateApi    = "http://production.shippingapis.com/ShippingAPI.dll?API=Verify&XML="
	RateRequestApi = "http://production.shippingapis.com/ShippingAPI.dll?API=RateV4&XML="
)

var uid = os.Getenv("USPS_USER_ID")

type Address struct {
	Id     string `json:"id" xml:"ID,attr"`
	Unit   string `json:"unit" xml:"Address1"`
	Street string `json:"street" xml:"Address2"`
	City   string `json:"city" xml:"City"`
	State  string `json:"state" xml:"State"`
	Zip5   string `json:"zip_5" xml:"Zip5"`
	Zip4   string `json:"zip_4" xml:"Zip4"`
}

type Request struct {
	Address  Address          `json:"address"`
	Packages []PackageRequest `json:"packages"`
}

type AddressValidateRequest struct {
	UserId   string  `xml:"USERID,attr"`
	Revision string  `xml:"Revision"`
	Address  Address `xml:"Address"`
}

type AddressValidateResponse struct {
	Address Address `xml:"Address"`
}

type RateV4Request struct {
	UserId   string           `xml:"USERID,attr"`
	Revision string           `xml:"Revision"`
	Packages []PackageRequest `xml:"Package"`
}

type RateV4Response struct {
	Packages []PackageResponse `xml:"Package"`
}

type PackageRequest struct {
	XMLName    xml.Name `xml:"Package"`
	Id         string   `xml:"ID,attr" json:"id"` // product id
	Service    string   `xml:"Service" json:"-"`
	ZipFrom    string   `xml:"ZipOrigination" json:"zip_from"`
	ZipTo      string   `xml:"ZipDestination" json:"zip_to"`
	Pounds     int      `xml:"Pounds" json:"pounds"`
	Ounces     float32  `xml:"Ounces" json:"ounces"`
	Container  string   `xml:"Container" json:"-"`
	Machinable string   `xml:"Machinable" json:"-"`
}

type PackageResponse struct {
	PackageRequest
	Postage Postage `xml:"Postage" json:"postage"`
}

type Postage struct {
	Id    string `xml:"CLASSID,attr" json:"id"`
	Type  string `xml:"MailService" json:"type"`
	Price string `xml:"Rate" json:"price"`
}

func getXML(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return data, err
}

// USPS Handler can verify and validation (entity) or perform a rate request.
func Handle(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	apigwp.LogRequest(r)

	switch r.QueryStringParameters["path"] {

	case "validate":

		var a Address
		err := json.Unmarshal([]byte(r.Body), &a)
		if err != nil {
			return apigwp.Bad(err)
		}

		inBytes, err := xml.Marshal(&AddressValidateRequest{uid, "1", a})
		if err != nil {
			return apigwp.Bad(err)
		}

		outBytes, err := getXML(ValidateApi + url.PathEscape(string(inBytes)))
		if err != nil {
			return apigwp.Bad(err)
		}

		var out = AddressValidateResponse{}
		err = xml.Unmarshal(outBytes, &out)
		if err != nil {
			return apigwp.Bad(err)
		}

		fmt.Println(out)

		fmt.Println(out.Address)

		return apigwp.Ok(out.Address)

	case "rate":
		var pp []PackageRequest
		_ = json.Unmarshal([]byte(r.Body), &pp)

		for i, p := range pp {
			p.Service = "PRIORITY"
			p.Container = "LG FLAT RATE BOX"
			p.Machinable = "TRUE"
			pp[i] = p
		}

		inBytes, _ := xml.Marshal(&RateV4Request{uid, "2", pp})
		outBytes, _ := getXML(RateRequestApi + url.PathEscape(string(inBytes)))

		var out RateV4Response
		_ = xml.Unmarshal(outBytes, &out)

		rates := map[string]map[string]map[string]string{}
		for _, p := range out.Packages {
			rates[p.Id] = map[string]map[string]string{"USPS": {strings.Split(p.Postage.Type, "&")[0]: p.Postage.Price}}
		}
		return apigwp.Ok(rates)
	}

	return apigwp.Bad(fmt.Errorf("nothing returned for [%v].\n", r))
}

func main() {
	lambda.Start(Handle)
}
