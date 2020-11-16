package model

type Tank struct {
	Pk
	Fk
	Water
	Animals []string `json:"animals"`
}
