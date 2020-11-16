package model

import "time"

type Moment struct {
	Unix int64 `json:"unix"`
}

func (m *Moment) Validate() error {
	if m.Unix == 0 {
		m.Unix = time.Now().Unix()
	}
	return nil
}
