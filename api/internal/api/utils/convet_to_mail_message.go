package utils

import (
	"github.com/CarlosEduardoAD/jobbo-api/internal/api/domain/email"
	"gopkg.in/gomail.v2"
)

func ConvertToMailMessage(email *email.Email) *gomail.Message {
	message := gomail.NewMessage()
	message.SetHeader("From", email.From)
	message.SetHeader("To", email.To)
	message.SetAddressHeader("Cc", email.From, email.From)
	message.SetHeader("Subject", email.Subject)
	message.SetBody("text/html", email.Body)
	return message
}
