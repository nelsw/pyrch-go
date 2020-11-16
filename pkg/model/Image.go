package model

type Image struct {
	Pk
	Fk
	Url
	Name
}

func (e *Image) Validate() error {
	if err := e.Pk.Validate(); err != nil {
		return err
	} else if err = e.Fk.Validate(); err != nil {
		return err
	} else if err = e.Name.Validate(); err != nil {
		return err
	} else if err = e.Url.Validate(); err != nil {
		return err
	}
	return nil
}

// Images is an aggregate value of Image ids.
type Images struct {
	Aggregate
}
