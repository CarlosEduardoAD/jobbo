package services

import (
	email_repo "github.com/CarlosEduardoAD/jobbo-api/internal/api/infra/repo/smtp"
	"gopkg.in/gomail.v2"
)

type EmailHandler struct {
	Repository *email_repo.EmailRepository
}

// NewEmailHandler creates a new instance of the EmailHandler struct.
//
// Parameters:
// - er: a pointer to an EmailRepository object.
//
// Returns:
// - a pointer to an EmailHandler object.
func NewEmailHandler(er *email_repo.EmailRepository) *EmailHandler {
	return &EmailHandler{
		Repository: er,
	}
}

// DeliverEmail delivers an email message.
//
// It takes a *gomail.Message object as a parameter and returns an error and a boolean value.
func (eh *EmailHandler) DeliverEmail(message *gomail.Message) (error, bool) {
	err := eh.Repository.DeliverEmail(message)

	if err != nil {
		return err, false
	}

	return nil, true
}
