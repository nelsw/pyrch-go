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
	allPaths = regexp.MustCompile(`claims|create|id|verify`)
	jwtPaths = regexp.MustCompile(`claims|id|verify`)
	jwtKey   = []byte(os.Getenv("JWT_KEY"))
)

func Handle(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	apigwp.LogRequest(r)

	// is the path valid?
	if !allPaths.Match([]byte(r.Path)) {
		return apigwp.Bad(fmt.Errorf("path not found [%s]", r.Path))
	}

	// is security validation required?
	if jwtPaths.Match([]byte(r.Path)) {

		// is there a security provided?
		token, ok := r.Headers["Authorization"]
		if !ok {
			return apigwp.Bad(fmt.Errorf(`authorization not found`))
		}

		// strip the "Bearer" prefix
		token = strings.ReplaceAll(token, `Bearer `, ``)

		// is the security valid?
		if jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(_ *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		}); err != nil {
			return apigwp.Unauthorized(err)
		} else if !jwtToken.Valid {
			return apigwp.Unauthorized(fmt.Errorf(`bad token, invalid segments`))
		}
	}

	// create claims
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Id:        r.Body,
		IssuedAt:  time.Now().Unix(),
	}

	// create security
	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)
	token = "Bearer " + token

	// return response
	if r.Path == "id" {
		return apigwp.OkInterface(token, claims.Id)
	} else if r.Path == "claims" {
		return apigwp.OkInterface(token, &claims)
	}

	return apigwp.OkVoid(token)
}

func main() {
	lambda.Start(Handle)
}
