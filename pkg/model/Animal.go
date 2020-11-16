package model

import "fmt"

type Animal struct {
	Url
	Name
	Handle string   `json:"handle"`
	Latin  string   `json:"latin"`
	Detail string   `json:"detail"`
	Images []string `json:"images"`
}

func (v *Animal) Validate() error {
	if err := v.Name.Validate(); err != nil {
		return err
	} else if v.Handle == "" && v.Latin == "" {
		return fmt.Errorf("animal is empty\n")
	}
	return v.Url.Validate()
}
