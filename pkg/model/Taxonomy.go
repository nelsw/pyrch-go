package model

type Taxonomy struct {
	Domain     string `json:"domain"`
	Kingdom    string `json:"kingdom"`
	Phylum     string `json:"phylum"`
	Class      string `json:"class"`
	Order      string `json:"order"`
	Infraorder string `json:"Infraorder"`
	Family     string `json:"family"`
	Tribe      string `json:"tribe"`
	Genus      string `json:"genus"`
	Species    string `json:"species"`
	SubSpecies string `json:"sub_species"`
	Variety    string `json:"variety"`
}

func (m *Taxonomy) Validate() error {
	return nil
}
