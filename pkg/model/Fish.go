package model

type Fish struct {
	Id
	Animal
}

func (e *Fish) Validate() error {
	if err := e.Id.Validate(); err != nil {
		return err
	} else if err = e.Animal.Validate(); err != nil {
		return err
	}
	return nil
}
