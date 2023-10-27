package email

import (
	"errors"
)

// Email entity
//   - From : string
//   - To : string
//   - Subject : string
//   - Body : string
type Email struct {
	From    string `json:"from" query:"from"`
	To      string `json:"to" query:"to"`
	Subject string `json:"subject" query:"subject"`
	Body    string `json:"body" query:"body"`
}

// NewEmail initializes a new Email struct with the given parameters.
//
// Parameters
//   - from: sender email address
//   - to: receiver email address
//   - subject:email subject
//   - body: email body
//
// It returns a pointer to a new Email struct initialized with the given values.
func NewEmail(from string, to string, subject string, body string) *Email {
	return &Email{
		From:    from,
		To:      to,
		Subject: subject,
		Body:    body}
}

// Validate validates the integrity of the Email receiver.
// It returns an error if the Email is invalid.
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
