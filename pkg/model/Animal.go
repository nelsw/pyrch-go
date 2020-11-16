package model

type Animal struct {
	Url    string   `json:"url"`
	Title  string   `json:"title"`
	Images []string `json:"images"`
}

func (v *Animal) Validate() error {
	return nil
}
