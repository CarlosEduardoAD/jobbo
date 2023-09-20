package services

import (
	email "github.com/CarlosEduardoAD/jobbo-api/internal/api/domain/email"
	email_repo "github.com/CarlosEduardoAD/jobbo-api/internal/api/infra/repo/services/email"
	"gopkg.in/gomail.v2"
)

func DeliverEmail(dialer *gomail.Dialer, e *email.Email) (error, bool) {
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
