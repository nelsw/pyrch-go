package model

type Fish struct {
	Pk
	Animal
}

func (e *Fish) Validate() error {
	if err := e.Pk.Validate(); err != nil {
		return err
	} else if err = e.Animal.Validate(); err != nil {
		return err
	}
	return nil
}
