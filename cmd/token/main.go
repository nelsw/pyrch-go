package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dgrijalva/jwt-go"
	"os"
	"pyrch-go/internal/apigwp"
	"regexp"
	"strings"
	"time"
)

var (
	allPaths = regexp.MustCompile(`claims|create|verify`)
	jwtPaths = regexp.MustCompile(`claims|verify`)
	jwtKey   = []byte(os.Getenv("JWT_KEY"))
)

func Handle(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	apigwp.LogRequest(r)

	// is the path valid?
	if !allPaths.Match([]byte(r.Path)) {
		return apigwp.BadRequest(fmt.Errorf("path not found [%s]", r.Path))
	}

	// is token validation required?
	if jwtPaths.Match([]byte(r.Path)) {

		// is there a token provided?
		token, ok := r.Headers["Authorize"]
		if !ok {
			return apigwp.BadRequest(fmt.Errorf(`token not found`))
		}

		// strip the "Bearer" prefix
		token = strings.ReplaceAll(token, `Bearer `, ``)

		// is the token valid?
		if jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(_ *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		}); err != nil || !jwtToken.Valid {
			return apigwp.Unauthorized(fmt.Errorf(`bad token, invalid segments or expired`))
		}
	}

	// create claims
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Id:        r.Body,
		IssuedAt:  time.Now().Unix(),
	}

	// create token
	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)
	token = "Bearer " + token

	// return response
	if r.Path == "claims" {
		return apigwp.OkInterface(token, &claims)
	}
	return apigwp.OkVoid(token)
}

func main() {
	lambda.Start(Handle)
}
