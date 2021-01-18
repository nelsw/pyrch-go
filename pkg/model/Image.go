package model

type Image struct {
	Id
	Url
	Name
}

func (e *Image) Validate() error {
	if err := e.Id.Validate(); err != nil {
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
