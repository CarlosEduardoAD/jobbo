package unit

import (
	"testing"

	email_service "github.com/CarlosEduardoAD/jobbo-api/internal/api/domain/handlers"
	email "github.com/CarlosEduardoAD/jobbo-api/internal/api/domain/model/email"
	email_repo "github.com/CarlosEduardoAD/jobbo-api/internal/api/infra/repo/smtp"
	"github.com/CarlosEduardoAD/jobbo-api/internal/api/utils"
	"github.com/stretchr/testify/assert"
)

func TestEmailCreationEmpty(t *testing.T) {
	email := &email.Email{}
	assert.Error(t, email.Validate(), "from is required")
}

func TestEmailCreationWithoutTo(t *testing.T) {
	email := &email.Email{From: "test@example.com"}
	assert.Error(t, email.Validate(), "to is required")
}

func TestEmailCreationWithoutSubject(t *testing.T) {
	email := &email.Email{From: "test@example.com", To: "test2@example.com"}
	assert.Error(t, email.Validate(), "subject is required")
}

func TestEmailCreationWithoutBody(t *testing.T) {
	email := &email.Email{From: "test@example.com", To: "test2@example.com", Subject: "Test"}
	assert.Error(t, email.Validate(), "body is required")
}

func TestEmailCreation(t *testing.T) {
	email := email.NewEmail("test@example.com", "test2@example.com", "Test", "Email body")
	assert.Equal(t, nil, email.Validate())
}

func TestEmailDeliver(t *testing.T) {
	dialer := utils.ConnectSMTP("smtp.gmail.com", 587, "karl.devcontato@gmail.com", "ehuf hvxx funu frov")
	emailToBeDelivered := utils.ConvertToMailMessage(email.NewEmail("carlosgoalfy@gmail.com", "karl.devcontato@gmail.com", "Test", "Email body"))
	emailRepository := email_repo.NewEmailService(dialer)
	emailHandler := email_service.NewEmailHandler(emailRepository)
	err, sucess := emailHandler.DeliverEmail(emailToBeDelivered)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, sucess)
}
