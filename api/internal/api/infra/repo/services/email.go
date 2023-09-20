	package email_repo

	import (
		"gopkg.in/gomail.v2"
	)

	type EmailActions interface {
		DeliverEmail() error
	}

	type EmailRepository struct {
		Server  *gomail.Dialer // dependency inversion
		Message *gomail.Message
	}

	func NewEmailService(server *gomail.Dialer, message *gomail.Message) *EmailRepository {
		return &EmailRepository{
			Server:  server,
			Message: message,
		}
	}

	func (r *EmailRepository) DeliverEmail() error {
		err := r.Server.DialAndSend(r.Message)

		if err != nil {
			return err
		}

		return nil
	}
