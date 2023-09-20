package unit

import (
	"testing"

	email "github.com/CarlosEduardoAD/jobbo-api/internal/api/domain/email"
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
	emailToBeDelivered := email.NewEmail("karl.devcontato@gmail.com", "karl.devcontato@gmail.com", "Test", "Email body")
	err, sucess := email.DeliverEmail(dialer, emailToBeDelivered)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, sucess)
}
