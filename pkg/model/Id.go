package model

import (
	"fmt"
	"github.com/google/uuid"
)

// Pk is the primary Id of an entity.
type Pk struct {
	Pk string `json:"pk"`
}

func (v *Pk) Validate() error {
	if v.Pk == "" {
		v.Pk = uuid.New().String()
	}
	return nil
}

// Fk is a CYA in case we need to traverse upwards to the root (parent entity) in our domain model.
// It works like a traditional Relational Data Foreign Key by defining n-n relationships between entities.
type Fk struct {
	Fk     string `json:"fk"`     // The key to join
	Table  string `json:"table"`  // The table to scan
	Column string `json:"column"` // The column of the key
}

func (e *Fk) Validate() error {
	if e.Table == "" {
		return fmt.Errorf("no table\n")
	} else if e.Fk == "" {
		return fmt.Errorf("no column\n")
	} else if e.Fk == "" {
		return fmt.Errorf("no fk\n")
	}
	return nil
}
