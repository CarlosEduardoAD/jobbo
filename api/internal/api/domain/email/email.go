package email

import (
	"errors"
)

// Email entity
// 	- From : string
// 	- To : string
// 	- Subject : string
// 	- Body : string
type Email struct {
	From    string `json:"from" query:"from"`
	To      string `json:"to" query:"to"`
	Subject string `json:"subject" query:"subject"`
	Body    string `json:"body" query:"body"`
}

// Validate email integrity
func (e *Email) Validate() error {
	if e.From == "" {
		return errors.New("from is required")
	}
	if e.To == "" {
		return errors.New("to is required")
	}
	if e.Subject == "" {
		return errors.New("subject is required")
	}
	if e.Body == "" {
		return errors.New("body is required")
	}

	return nil
}

// Init new Email entity
func NewEmail(from string, to string, subject string, body string) *Email {
	return &Email{
		From:    from,
		To:      to,
		Subject: subject,
		Body:    body}
}
