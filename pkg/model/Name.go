package model

import (
	"fmt"
)

type Name struct {
	Name string `json:"name"`
}

func (e *Name) Validate() error {
	if e.Name == "" {
		return fmt.Errorf("no name\n")
	}
	return nil
}
