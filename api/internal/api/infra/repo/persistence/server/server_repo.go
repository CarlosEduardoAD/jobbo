package server

import (
	"github.com/CarlosEduardoAD/jobbo-api/internal/api/domain/model/server"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type ServerActions interface {
	Create() error
	Find(id string) (*server.Server, error)
	Update(id string, server *server.Server) error
	Delete(id string) error
}

type ServerRepository struct {
	db *gorm.DB
}

func NewServerRepository(db *gorm.DB) *ServerRepository {
	return &ServerRepository{
		db: db,
	}
}

func (serve *ServerRepository) Create(s *server.Server) error {
	context := serve.db.Create(s)

	if context.Error != nil {
		context.Rollback()
		log.Error(context.Error)
		return context.Error
	}

	context.Commit()
	return nil
}

func (serve *ServerRepository) Find(id string) (*server.Server, error) {
	var serverDTO server.Server
	context := serve.db.First(&serverDTO, "id = ?", id)

	if context.Error != nil {
		context.Rollback()
		log.Error(context.Error)
		return &serverDTO, context.Error
	}

	context.Commit()
	return &serverDTO, nil
}

func (serve *ServerRepository) Update(id string, s *server.Server) error {
	context := serve.db.Model(&server.Server{}).Where("id = ?", id).Updates(s)

	if context.Error != nil {
		context.Rollback()
		log.Error(context.Error)
		return context.Error
	}

	context.Commit()
	return nil
}

func (serve *ServerRepository) Delete(id string) error {
	context := serve.db.Model(&server.Server{}).Where("id = ?", id).Delete(&server.Server{})

	if context.Error != nil {
		context.Rollback()
		log.Error(context.Error)
		return context.Error
	}

	context.Commit()
	return nil
}
