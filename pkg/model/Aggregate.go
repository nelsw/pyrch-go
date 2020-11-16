package model

// Aggregate is an generic struct used  to  embed common fields and functionality into the domain model.
type Aggregate struct {
	Pk
	Fk
	Aggregates
}

type Aggregates struct {
	Pks  []string `json:"pks"`
	Def  string   `json:"def"`
	Size int      `json:"size"`
}

func (v *Aggregate) Validate() error {
	if err := v.Pk.Validate(); err != nil {
		return err
	} else if err := v.Fk.Validate(); err != nil {
		return err
	}
	size := len(v.Pks)
	if size != v.Size {
		v.Size = size
	}
	if v.Def == "" {
		v.Def = v.Pks[0]
	}
	return nil
}
