package model

type Fish struct {
	Pk
	Fk
	Animal
}

func (e *Fish) Validate() error {
	if err := e.Pk.Validate(); err != nil {
		return err
	} else if err = e.Fk.Validate(); err != nil {
		return err
	} else if err = e.Animal.Validate(); err != nil {
		return err
	}
	return nil
}
