package message_test

import (
	"testing"
	"time"

	"github.com/CarlosEduardoAD/jobbo-api/internal/api/domain/model/message"
	"github.com/CarlosEduardoAD/jobbo-api/internal/api/domain/model/server"
	"github.com/google/uuid"
)

func TestNewMessage(t *testing.T) {
	userID := uuid.New().String()
	orgID := uuid.New().String()
	campaignID := uuid.New().String()
	from := "sender@example.com"
	to := "recipient@example.com"
	subject := "Test message"
	body := "This is a test message"
	sentAt := time.Now()
	serverId := uuid.New().String()
	server := server.Server{
		ID:   uuid.New().String(),
		Name: "Test server",
	}

	msg, err := message.NewMessage(userID, orgID, campaignID, from, to, subject, body, sentAt, serverId, server)

	if err != nil {
		t.Errorf("NewMessage returned an error: %v", err)
	}

	if msg.ID == "" {
		t.Error("NewMessage did not generate an ID")
	}

	if msg.UserID != userID {
		t.Errorf("NewMessage returned a message with the wrong user ID. Expected %s, got %s", userID, msg.UserID)
	}

	if msg.OrganizationId != orgID {
		t.Errorf("NewMessage returned a message with the wrong organization ID. Expected %s, got %s", orgID, msg.OrganizationId)
	}

	if msg.CampaignId != campaignID {
		t.Errorf("NewMessage returned a message with the wrong campaign ID. Expected %s, got %s", campaignID, msg.CampaignId)
	}

	if msg.From != from {
		t.Errorf("NewMessage returned a message with the wrong sender. Expected %s, got %s", from, msg.From)
	}

	if msg.To != to {
		t.Errorf("NewMessage returned a message with the wrong recipient. Expected %s, got %s", to, msg.To)
	}

	if msg.Subject != subject {
		t.Errorf("NewMessage returned a message with the wrong subject. Expected %s, got %s", subject, msg.Subject)
	}

	if msg.Body != body {
		t.Errorf("NewMessage returned a message with the wrong body. Expected %s, got %s", body, msg.Body)
	}

	if !msg.SentAt.Equal(sentAt) {
		t.Errorf("NewMessage returned a message with the wrong sentAt time. Expected %v, got %v", sentAt, msg.SentAt)
	}

	if msg.Server.ID != server.ID {
		t.Errorf("NewMessage returned a message with the wrong server ID. Expected %s, got %s", server.ID, msg.Server.ID)
	}

	if msg.Server.Name != server.Name {
		t.Errorf("NewMessage returned a message with the wrong server name. Expected %s, got %s", server.Name, msg.Server.Name)
	}
}

func TestMessage_Validate(t *testing.T) {
	msg := message.Message{
		UserID:         uuid.New().String(),
		OrganizationId: uuid.New().String(),
		CampaignId:     uuid.New().String(),
		From:           "sender@example.com",
		To:             "recipient@example.com",
		Subject:        "Test message",
		Body:           "This is a test message",
	}

	err := msg.Validate()

	if err != nil {
		t.Errorf("Validate returned an error for a valid message: %v", err)
	}

	msg.UserID = ""

	err = msg.Validate()

	if err == nil {
		t.Error("Validate did not return an error for a message with an empty user ID")
	}

	if err.Error() != "user id is empty" {
		t.Errorf("Validate returned the wrong error for a message with an empty user ID. Expected 'user id is empty', got '%v'", err)
	}

	msg.UserID = uuid.New().String()
	msg.OrganizationId = ""

	err = msg.Validate()

	if err == nil {
		t.Error("Validate did not return an error for a message with an empty organization ID")
	}

	if err.Error() != "organization id is empty" {
		t.Errorf("Validate returned the wrong error for a message with an empty organization ID. Expected 'organization id is empty', got '%v'", err)
	}

	msg.OrganizationId = uuid.New().String()
	msg.CampaignId = ""

	err = msg.Validate()

	if err == nil {
		t.Error("Validate did not return an error for a message with an empty campaign ID")
	}

	if err.Error() != "campaign id is empty" {
		t.Errorf("Validate returned the wrong error for a message with an empty campaign ID. Expected 'campaign id is empty', got '%v'", err)
	}

	msg.CampaignId = uuid.New().String()
	msg.From = ""

	err = msg.Validate()

	if err == nil {
		t.Error("Validate did not return an error for a message with an empty sender")
	}

	if err.Error() != "from is empty" {
		t.Errorf("Validate returned the wrong error for a message with an empty sender. Expected 'from is empty', got '%v'", err)
	}

	msg.From = "arlos"
	msg.To = ""

	err = msg.Validate()

	if err == nil {
		t.Error("Validate did not return an error for a message with an empty recipient")
	}

	if err.Error() != "to is empty" {
		t.Errorf("Validate returned the wrong error for a message with an empty recipient. Expected 'to is empty', got '%v'", err)
	}

	msg.To = "emerson"
	msg.Subject = ""

	err = msg.Validate()

	if err == nil {
		t.Error("Validate did not return an error for a message with an empty subject")
	}

	if err.Error() != "subject is empty" {
		t.Errorf("Validate returned the wrong error for a message with an empty subject. Expected 'subject is empty', got '%v'", err)
	}

	msg.Subject = "Test message"
	msg.Body = ""

	err = msg.Validate()

	if err == nil {
		t.Error("Validate did not return an error for a message with an empty body")
	}

	if err.Error() != "body is empty" {
		t.Errorf("Validate returned the wrong error for a message with an empty body. Expected 'body is empty', got '%v'", err)
	}

	// Repeat the above pattern for each field that needs to be validated
}
