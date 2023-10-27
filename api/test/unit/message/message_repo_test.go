package message

import (
	"testing"
	"time"

	"github.com/CarlosEduardoAD/jobbo-api/internal/api/domain/model/message"
	"github.com/CarlosEduardoAD/jobbo-api/internal/api/domain/model/server"
	msg_repo "github.com/CarlosEduardoAD/jobbo-api/internal/api/infra/repo/persistence/message"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestMessageRepository_Create(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	repo := msg_repo.NewMessageRepository(db)

	m := &message.Message{
		ID:             "91c741f9-f08d-4dc8-9143-12ab6cf9dfae",
		UserID:         "user1",
		OrganizationId: "org1",
		CampaignId:     "camp1",
		From:           "from@example.com",
		To:             "to@example.com",
		Subject:        "Test message",
		Body:           "This is a test message",
		SentAt:         time.Now(),
		ServerID:       "server1",
		Server: server.Server{
			ID:           "91c741f9-f08d-4dc8-9143-12ab6cf9dfae",
			Name:         "singelo servidor",
			Address:      "servidor.smtp",
			Port:         "379",
			User:         "emerson",
			Password:     "amogusbagos",
			Cryptography: "TLS",
		},
	}

	err = repo.Create(m)
	assert.NoError(t, err)

	// Check that the message was created
	var count int64
	db.Model(&message.Message{}).Count(&count)
	assert.Equal(t, int64(1), count)
}

func TestMessageRepository_Find(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	repo := msg_repo.NewMessageRepository(db)

	m := &message.Message{
		ID:             "1",
		UserID:         "user1",
		OrganizationId: "org1",
		CampaignId:     "camp1",
		From:           "from@example.com",
		To:             "to@example.com",
		Subject:        "Test message",
		Body:           "This is a test message",
		SentAt:         time.Now(),
		ServerID:       "server1",
		Server: server.Server{
			ID:   "server1",
			Name: "Test server",
		},
	}

	err = repo.Create(m)
	assert.NoError(t, err)

	// Find the message
	found, err := repo.Find("1")
	assert.NoError(t, err)
	assert.Equal(t, m, found)
}

func TestMessageRepository_Update(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	repo := msg_repo.NewMessageRepository(db)

	m := &message.Message{
		ID:             "1",
		UserID:         "user1",
		OrganizationId: "org1",
		CampaignId:     "camp1",
		From:           "from@example.com",
		To:             "to@example.com",
		Subject:        "Test message",
		Body:           "This is a test message",
		SentAt:         time.Now(),
		ServerID:       "server1",
		Server: server.Server{
			ID:   "server1",
			Name: "Test server",
		},
	}

	err = repo.Create(m)
	assert.NoError(t, err)

	// Update the message
	m.Subject = "Updated subject"
	err = repo.Update("1", m)
	assert.NoError(t, err)

	// Find the message and check that it was updated
	found, err := repo.Find("1")
	assert.NoError(t, err)
	assert.Equal(t, m, found)
}

func TestMessageRepository_Delete(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	repo := msg_repo.NewMessageRepository(db)

	m := &message.Message{
		ID:             "1",
		UserID:         "user1",
		OrganizationId: "org1",
		CampaignId:     "camp1",
		From:           "from@example.com",
		To:             "to@example.com",
		Subject:        "Test message",
		Body:           "This is a test message",
		SentAt:         time.Now(),
		ServerID:       "server1",
		Server: server.Server{
			ID:   "server1",
			Name: "Test server",
		},
	}

	err = repo.Create(m)
	assert.NoError(t, err)

	// Delete the message
	err = repo.Delete("1")
	assert.NoError(t, err)

	// Check that the message was deleted
	var count int64
	db.Model(&message.Message{}).Count(&count)
	assert.Equal(t, int64(0), count)
}

func TestMessageRepository_FindWithOrganizationId(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	repo := msg_repo.NewMessageRepository(db)

	m1 := &message.Message{
		ID:             "1",
		UserID:         "user1",
		OrganizationId: "org1",
		CampaignId:     "camp1",
		From:           "from@example.com",
		To:             "to@example.com",
		Subject:        "Test message 1",
		Body:           "This is a test message 1",
		SentAt:         time.Now(),
		ServerID:       "server1",
		Server: server.Server{
			ID:   "server1",
			Name: "Test server",
		},
	}

	m2 := &message.Message{
		ID:             "2",
		UserID:         "user1",
		OrganizationId: "org2",
		CampaignId:     "camp1",
		From:           "from@example.com",
		To:             "to@example.com",
		Subject:        "Test message 2",
		Body:           "This is a test message 2",
		SentAt:         time.Now(),
		Server: server.Server{
			ID:   "server1",
			Name: "Test server",
		},
	}

	err = repo.Create(m1)
	assert.NoError(t, err)

	err = repo.Create(m2)
	assert.NoError(t, err)

	// Find messages with organization ID "org1"
	found, err := repo.FindWithOrganizationId("org1")
	assert.NoError(t, err)
	assert.Equal(t, []*message.Message{m1}, found)
}

func TestMessageRepository_FindWithCampaignId(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	repo := msg_repo.NewMessageRepository(db)

	m1 := &message.Message{
		ID:             "1",
		UserID:         "user1",
		OrganizationId: "org1",
		CampaignId:     "camp1",
		From:           "from@example.com",
		To:             "to@example.com",
		Subject:        "Test message 1",
		Body:           "This is a test message 1",
		SentAt:         time.Now(),
		ServerID:       "server1",
		Server: server.Server{
			ID:   "server1",
			Name: "Test server",
		},
	}

	m2 := &message.Message{
		ID:             "2",
		UserID:         "user1",
		OrganizationId: "org2",
		CampaignId:     "camp2",
		From:           "from@example.com",
		To:             "to@example.com",
		Subject:        "Test message 2",
		Body:           "This is a test message 2",
		SentAt:         time.Now(),
		ServerID:       "server2",
		Server: server.Server{
			ID:   "server1",
			Name: "Test server",
		},
	}

	err = repo.Create(m1)
	assert.NoError(t, err)

	err = repo.Create(m2)
	assert.NoError(t, err)

	// Find messages with campaign ID "camp1"
	found, err := repo.FindWithCampaignId("camp1")
	assert.NoError(t, err)
	assert.Equal(t, []*message.Message{m1}, found)
}
