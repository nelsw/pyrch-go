package model

// Aggregate is an generic struct used  to  embed common fields and functionality into the domain model.
type Aggregate struct {
	Id
	Aggregates
}

type Aggregates struct {
	Ids  []string `json:"ids"`
	Def  string   `json:"def"`
	Size int      `json:"size"`
}

func (v *Aggregate) Validate() error {
	if err := v.Id.Validate(); err != nil {
		return err
	}
	size := len(v.Ids)
	if size != v.Size {
		v.Size = size
	}
	if v.Def == "" {
		v.Def = v.Ids[0]
	}
	return nil
}
