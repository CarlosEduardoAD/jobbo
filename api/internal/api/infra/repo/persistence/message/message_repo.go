package message

import (
	"log"

	"github.com/CarlosEduardoAD/jobbo-api/internal/api/domain/model/message"
	"github.com/CarlosEduardoAD/jobbo-api/internal/api/domain/model/server"
	"gorm.io/gorm"
)

type MessageActions interface {
	Create() error
	Find(id string) (*server.Server, error)
	Update(id string, server *server.Server) error
	Delete(id string) error
	FindWithOrganizationId(organizationId string) ([]*message.Message, error)
	FindWithCampaignId(campaignId string) ([]*message.Message, error)
}

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{
		db: db,
	}
}

func (message *MessageRepository) Create(m *message.Message) error {
	context := message.db.Create(m)

	if context.Error != nil {
		context.Rollback()
		log.Panic(context.Error)
		return context.Error
	}

	context.Commit()
	return nil
}

func (msg *MessageRepository) Find(id string) (*message.Message, error) {
	var messageDTO message.Message
	context := msg.db.First(&messageDTO, "id = ?", id)

	if context.Error != nil {
		context.Rollback()
		return &messageDTO, context.Error
	}

	context.Commit()
	return &messageDTO, nil
}

func (msg *MessageRepository) Update(id string, m *message.Message) error {
	context := msg.db.Model(&message.Message{}).Where("id = ?", id).Updates(m)

	if context.Error != nil {
		context.Rollback()
		return context.Error
	}

	context.Commit()
	return nil
}

func (msg *MessageRepository) Delete(id string) error {
	context := msg.db.Delete(&message.Message{}, "id = ?", id)

	if context.Error != nil {
		context.Rollback()
		return context.Error
	}

	context.Commit()
	return nil
}

func (msg *MessageRepository) FindWithOrganizationId(organizationId string) ([]*message.Message, error) {
	var messages []*message.Message
	context := msg.db.Find(&messages, "organization_id = ?", organizationId)

	if context.Error != nil {
		context.Rollback()
		return messages, context.Error
	}

	context.Commit()
	return messages, nil
}

func (msg *MessageRepository) FindWithCampaignId(campaignId string) ([]*message.Message, error) {
	var messages []*message.Message
	context := msg.db.Find(&messages, "campaign_id = ?", campaignId)

	if context.Error != nil {
		context.Rollback()
		return messages, context.Error
	}

	context.Commit()
	return messages, nil
}
