package email_repo

import (
	"gopkg.in/gomail.v2"
)

type EmailActions interface {
	DeliverEmail() error
}

type EmailRepository struct {
	Server *gomail.Dialer // dependency inversion
}

func NewEmailService(server *gomail.Dialer) *EmailRepository {
	return &EmailRepository{
		Server: server,
	}
}

func (r *EmailRepository) DeliverEmail(message *gomail.Message) error {
	err := r.Server.DialAndSend(message)

	if err != nil {
		return err
	}

	return nil
}
