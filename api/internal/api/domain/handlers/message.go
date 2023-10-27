package services

import (
	"github.com/CarlosEduardoAD/jobbo-api/internal/api/domain/model/message"
	message_repo "github.com/CarlosEduardoAD/jobbo-api/internal/api/infra/repo/persistence/message"
)

type MessageHandler struct {
	message_repo *message_repo.MessageRepository
}

func NewMessageHandler(repo *message_repo.MessageRepository) *MessageHandler {
	return &MessageHandler{
		message_repo: repo,
	}
}

func (mh *MessageHandler) CreateMessage(message *message.Message) error {
	err := mh.message_repo.Create(message)

	if err != nil {
		return err
	}

	return nil
}

func (mh *MessageHandler) FindMessage(id string) (*message.Message, error) {
	result, err := mh.message_repo.Find(id)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (mh *MessageHandler) UpdateMessage(id string, message *message.Message) error {
	err := mh.message_repo.Update(id, message)

	if err != nil {
		return err
	}

	return nil
}

func (mh *MessageHandler) FindMessageWithOrganizationId(organizationId string) ([]*message.Message, error) {
	result, err := mh.message_repo.FindWithOrganizationId(organizationId)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (mh *MessageHandler) FindMessageWithCampaignId(campaignId string) ([]*message.Message, error) {
	result, err := mh.message_repo.FindWithCampaignId(campaignId)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (mh *MessageHandler) DeleteMessage(id string) error {
	err := mh.message_repo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
