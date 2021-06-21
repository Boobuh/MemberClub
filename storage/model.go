package storage

import (
	"errors"
	"net/mail"
	"time"
)

var (
	ErrInvalidName  = errors.New("invalid name")
	ErrInvalidEmail = errors.New("invalid email")
)

type Members []Member //надбудова над масивами

type Member struct {
	Name             string    `json:"name"`
	Email            string    `json:"email"`
	RegistrationDate time.Time `json:"registration_date"`
}

func (m Member) Validate() error {

	if !isLetter(m.Name) {
		return ErrInvalidName
	}
	_, err := mail.ParseAddress(m.Email)
	if err != nil {
		return ErrInvalidEmail
	}
	return nil

}

func (m *Member) SetRegistrationDate() {
	m.RegistrationDate = time.Now()
}

func isLetter(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}
