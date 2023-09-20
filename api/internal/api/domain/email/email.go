package email

import (
	"errors"

	email_repo "github.com/CarlosEduardoAD/jobbo-api/internal/api/infra/repo/services"
	"gopkg.in/gomail.v2"
)

type Email struct {
	From    string `json:"from" query:"from"`
	To      string `json:"to" query:"to"`
	Subject string `json:"subject" query:"subject"`
	Body    string `json:"body" query:"body"`
}

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

// This works as a constructor

func NewEmail(from string, to string, subject string, body string) *Email {
	return &Email{
		From:    from,
		To:      to,
		Subject: subject,
		Body:    body}
}

func DeliverEmail(dialer *gomail.Dialer, e *Email) (error, bool) {
	var err error

	if err = e.Validate(); err != nil {
		return err, false
	}

	message := gomail.NewMessage()
	message.SetHeader("From", e.From)
	message.SetHeader("To", e.To)
	message.SetAddressHeader("Cc", e.From, e.From)
	message.SetHeader("Subject", e.Subject)
	message.SetBody("text/html", e.Body)

	emailRepo := email_repo.NewEmailService(dialer, message)
	err = emailRepo.DeliverEmail()

	if err != nil {
		return err, false
	}

	return nil, true
}
