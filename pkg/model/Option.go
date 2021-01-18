package model

type Option struct {
	Id
	Price   int64  `json:"price"`   // 7900 = $79.00, stripe thinks it makes cents
	Weight  int    `json:"weight"`  // 170 = 1.7, to avoid decimals entirely
	Label   string `json:"label"`   // oz, lb, kilo, ton, w/e
	Stock   int    `json:"stock"`   // quantity available
	Address string `json:"address"` // shipping departure location
}

func (e *Option) Validate() error {
	if err := e.Id.Validate(); err != nil {
		return err
	}
	return nil
}

// Options is an aggregate value of Option ids.
type Options struct {
	Aggregate
}
