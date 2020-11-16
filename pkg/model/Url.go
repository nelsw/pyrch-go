package model

import "fmt"

type Url struct {
	Url string `json:"url"`
}

func (e *Url) Validate() error {
	if e.Url == "" {
		return fmt.Errorf("no Url\n")
	}
	return nil
}

type Urls struct {
	Aggregate
}
