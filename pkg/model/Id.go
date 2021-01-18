package model

import (
	"github.com/google/uuid"
)

// Id is the primary key of an entity.
type Id struct {
	Id string `json:"id"`
}

func (v *Id) Validate() error {
	if v.Id == "" {
		v.Id = uuid.New().String()
	}
	return nil
}
