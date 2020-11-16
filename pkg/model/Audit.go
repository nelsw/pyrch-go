package model

type Audit struct {
	Updated Moment `json:"updated"`
}

func (e *Audit) Validate() error {
	return e.Updated.Validate()
}
