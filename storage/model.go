package storage

import (
	"errors"
	"net/mail"
	"time"
)

type Members []Member //надбудова над масивами

type Member struct {
	Name             string
	Email            string
	ID               int
	RegistrationDate time.Time
}

func (m Member) Validate() error {

	if !isLetter(m.Name) {
		return errors.New("")
	}
	_, err := mail.ParseAddress(m.Email)
	return err

}

func (m *Member) SetRegistrationDate() {
	m.RegistrationDate = time.Now()
}

//perevirka na bukvy
func isLetter(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}