package services

import (
	email_repo "github.com/CarlosEduardoAD/jobbo-api/internal/api/infra/repo/smtp"
	"gopkg.in/gomail.v2"
)

type EmailHandler struct {
	Repository *email_repo.EmailRepository
}

func NewEmailHandler(er *email_repo.EmailRepository) *EmailHandler {
	return &EmailHandler{
		Repository: er,
	}
}

func (eh *EmailHandler) DeliverEmail(message *gomail.Message) (error, bool) {
	err := eh.Repository.DeliverEmail(message)

	if err != nil {
		return err, false
	}

	return nil, true
}
