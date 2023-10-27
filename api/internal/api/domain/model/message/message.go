package message

import (
	"errors"
	"time"

	"github.com/CarlosEduardoAD/jobbo-api/internal/api/domain/model/server"
	"github.com/google/uuid"
)

// Message is a struct that represents a message
type Message struct {
	ID             string `gorm:"primaryKey"`
	UserID         string `json:"userId" query:"userId"`
	OrganizationId string `json:"organizationId" query:"organizationId"`
	CampaignId     string `json:"campaignId" query:"campaignId"`
	From           string `json:"from" query:"from"`
	To             string `json:"to" query:"to"`
	Subject        string `json:"subject" query:"subject"`
	Body           string `json:"body" query:"body"`
	SentAt         time.Time
	ServerID       string        `json:"serverId" query:"serverId"` // Chave estrangeira
	Server         server.Server `gorm:"foreignKey:ServerID"`
}

func NewMessage(userID string, orgId string, campaignId string, from string, to string, subject string, body string, sentAt time.Time, severId string, server server.Server) (*Message, error) {

	return &Message{
		ID:             uuid.New().String(),
		UserID:         userID,
		OrganizationId: orgId,
		CampaignId:     campaignId,
		From:           from,
		To:             to,
		Subject:        subject,
		Body:           body,
		SentAt:         sentAt,
		ServerID:       severId,
		Server:         server,
	}, nil
}

func (m *Message) Validate() error {
	if m.UserID == "" {
		return errors.New("user id is empty")
	}

	if m.OrganizationId == "" {
		return errors.New("organization id is empty")
	}

	if m.CampaignId == "" {
		return errors.New("campaign id is empty")
	}

	if m.From == "" {
		return errors.New("from is empty")
	}

	if m.To == "" {
		return errors.New("to is empty")
	}

	if m.Subject == "" {
		return errors.New("subject is empty")
	}

	if m.Body == "" {
		return errors.New("body is empty")
	}

	if m.ServerID == "" {
		return errors.New("server id is empty")
	}

	return nil
}
