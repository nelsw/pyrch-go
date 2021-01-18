package model

import (
	"fmt"
	"unicode"
)

type Credential struct {
	Id
	Password string `json:"password"`
}

func (v *Credential) Validate() error {
	if err := validatePassword(v.Password); err != nil {
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
