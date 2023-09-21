package services

import (
	email_repo "github.com/CarlosEduardoAD/jobbo-api/internal/api/infra/repo/smtp"
)

func DeliverEmail(er *email_repo.EmailRepository) (error, bool) {
	err := er.DeliverEmail()

	if err != nil {
		return err, false
	}

	return nil, true
}
