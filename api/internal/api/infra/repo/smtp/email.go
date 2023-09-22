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

// NewEmailService creates a new EmailRepository instance.
//
// It takes a `server` parameter of type `*gomail.Dialer`, which represents the server to be used for sending emails.
// It returns a pointer to an `EmailRepository` object.
func NewEmailService(server *gomail.Dialer) *EmailRepository {
	return &EmailRepository{
		Server: server,
	}
}

// DeliverEmail delivers the given email message.
//
// The function takes a *gomail.Message as a parameter and returns an error.
func (r *EmailRepository) DeliverEmail(message *gomail.Message) error {
	err := r.Server.DialAndSend(message)

	if err != nil {
		return err
	}

	return nil
}
