package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
	"os"
	"pyrch-go/internal/apigwp"
	"strconv"
)

func init() {
	stripe.Key = os.Getenv("STRIPE_SK_" + os.Getenv("STAGE"))
}

func Handle(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	apigwp.LogRequest(r)

	s := r.QueryStringParameters["path"]
	if s == "key" {
		return apigwp.Ok(os.Getenv("STRIPE_PK_" + os.Getenv("STAGE")))
	}

	i64, _ := strconv.ParseInt(r.QueryStringParameters["amount"], 10, 64)
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(i64),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		PaymentMethodTypes: []*string{
			stripe.String("card"),
		},
	}
	pi, err := paymentintent.New(params)
	if err != nil {
		return apigwp.Bad(err)
	}

	return apigwp.Ok(pi)
}

func main() {
	lambda.Start(Handle)
}
