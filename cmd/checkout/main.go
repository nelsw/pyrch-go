package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/checkout/session"
	"os"
	"pyrch-go/internal/apigwp"
)

var successUrl = os.Getenv("SUCCESS_URL")
var failureUrl = os.Getenv("FAILURE_URL")

func init() {
	stripe.Key = os.Getenv("STRIPE_KEY")
}

func HandleRequest(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	apigwp.LogRequest(r)

	priceId, ok := r.QueryStringParameters["priceId"]
	if !ok {
		return apigwp.BadRequest("missing priceId")
	}

	params := &stripe.CheckoutSessionParams{
		SuccessURL: &successUrl,
		CancelURL:  &failureUrl,
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		Mode: stripe.String(string(stripe.CheckoutSessionModeSubscription)),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(priceId),
				Quantity: stripe.Int64(1),
			},
		},
	}

	s, err := session.New(params)
	if err != nil {
		return apigwp.BadRequest(err.Error())
	}

	return apigwp.Ok(struct {
		SessionID string `json:"sessionId"`
	}{
		SessionID: s.ID,
	})
}

func main() {
	lambda.Start(HandleRequest)
}
