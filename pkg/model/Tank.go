package model

type Tank struct {
	Id
	Water
	Animals []string `json:"animals"`
}
