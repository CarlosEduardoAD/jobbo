package unit

import (
	"testing"

	email "github.com/CarlosEduardoAD/jobbo-api/internal/api/domain/model/email"
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
	email := email.NewEmail("test@example.com", "test2@example.com", "Test", "Email body", "5aa4b806-f8c4-4180-9034-e5e9ec4d7dc3")
	assert.Equal(t, nil, email.Validate())
}

// func TestEmailDeliver(t *testing.T) {
// 	dialer := utils.ConnectSMTP("smtp.gmail.com", 587, "karl.devcontato@gmail.com", "ehuf hvxx funu frov")
// 	emailToBeDelivered := utils.ConvertToMailMessage(email.NewEmail("carlosgoalfy@gmail.com", "karl.devcontato@gmail.com", "Test", "Email body"))
// 	emailRepository := email_repo.NewEmailService(dialer)
// 	emailHandler := email_service.NewEmailHandler(emailRepository)
// 	err, sucess := emailHandler.DeliverEmail(emailToBeDelivered)
// 	assert.Equal(t, nil, err)
// 	assert.Equal(t, true, sucess)
// }
