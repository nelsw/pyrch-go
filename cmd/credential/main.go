package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"pyrch-go/internal/apigwp"
	"pyrch-go/internal/faas"
	"unicode"
)

type Credential struct {
	Id       string `json:"id"`
	Password string `json:"password"`
	UserId   string `json:"user_id"`
}

func (v *Credential) Validate() error {
	if v.Id == "" {
		return fmt.Errorf(`id is blank`)
	} else if err := validatePassword(v.Password); err != nil {
		return err
	} else {
		return nil
	}
}

// Thanks https://stackoverflow.com/a/25840157.
func validatePassword(s string) error {
	var number, upper, special bool
	length := 0
	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			number = true
			length++
		case unicode.IsUpper(c):
			upper = true
			length++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
			length++
		case unicode.IsLetter(c) || c == ' ':
			length++
		}
	}
	if length < 8 || length > 24 {
		return fmt.Errorf("bad password, must contain 8-24 characters")
	} else if !number {
		return fmt.Errorf("bad password, must contain at least 1 number")
	} else if !upper {
		return fmt.Errorf("bad password, must contain at least 1 uppercase letter")
	} else if !special {
		return fmt.Errorf("bad password, must contain at least 1 special character")
	} else {
		return nil
	}
}

func Handle(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	apigwp.LogRequest(r)

	if r.Path == "save" {
		c := Credential{}
		if err := json.Unmarshal([]byte(r.Body), &c); err != nil {
			return apigwp.Bad(err)
		} else if err = c.Validate(); err != nil {
			return apigwp.Bad(err)
		} else {
			// fall through to repo handler
		}
	}

	return faas.InvokeIt("repo", r), nil
}

func main() {
	lambda.Start(Handle)
}
