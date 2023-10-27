package unit

import (
	"testing"

	"github.com/CarlosEduardoAD/jobbo-api/internal/api/domain/model/server"
	repository "github.com/CarlosEduardoAD/jobbo-api/internal/api/infra/repo/persistence/server"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestServerRepository_Create(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}
	err = db.AutoMigrate(&server.Server{})
	if err != nil {
		t.Fatalf("failed to migrate database: %v", err)
	}

	repo := repository.NewServerRepository(db)

	s := &server.Server{
		Name: "teste",
	}

	err = repo.Create(s)
	assert.NoError(t, err)

	err = repo.Create(s)
	assert.Error(t, err)
}

func TestServerRepository_Find(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}
	err = db.AutoMigrate(&server.Server{})
	if err != nil {
		t.Fatalf("failed to migrate database: %v", err)
	}

	repo := repository.NewServerRepository(db)

	s := &server.Server{
		Name: "test",
	}
	err = repo.Create(s)
	assert.NoError(t, err)

	found, err := repo.Find(s.ID)
	assert.NoError(t, err)
	assert.Equal(t, s.ID, found.ID)

	_, err = repo.Find("invalid-id")
	assert.Error(t, err)
}

func TestServerRepository_Update(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}
	err = db.AutoMigrate(&server.Server{})
	if err != nil {
		t.Fatalf("failed to migrate database: %v", err)
	}

	repo := repository.NewServerRepository(db)

	s := &server.Server{
		Name: "test",
	}
	err = repo.Create(s)
	assert.NoError(t, err)

	s.Name = "updated"
	err = repo.Update(s.ID, s)
	assert.NoError(t, err)

	found, err := repo.Find(s.ID)
	assert.NoError(t, err)
	assert.Equal(t, "updated", found.Name)

	err = repo.Update("invalid-id", s)
	assert.Error(t, err)
}

func TestServerRepository_Delete(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}
	err = db.AutoMigrate(&server.Server{})
	if err != nil {
		t.Fatalf("failed to migrate database: %v", err)
	}

	repo := repository.NewServerRepository(db)

	s := &server.Server{
		Name: "test",
	}
	err = repo.Create(s)
	assert.NoError(t, err)

	err = repo.Delete(s.ID)
	assert.NoError(t, err)

	_, err = repo.Find(s.ID)
	assert.Error(t, err)

	err = repo.Delete("invalid-id")
	assert.Error(t, err)
}
