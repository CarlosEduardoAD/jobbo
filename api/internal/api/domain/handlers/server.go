package services

import (
	"github.com/CarlosEduardoAD/jobbo-api/internal/api/domain/model/server"
	server_repo "github.com/CarlosEduardoAD/jobbo-api/internal/api/infra/repo/persistence/server"
)

type ServerHandler struct {
	server_repo *server_repo.ServerRepository
}

func NewServerHandler(repo *server_repo.ServerRepository) *ServerHandler {
	return &ServerHandler{
		server_repo: repo,
	}
}

func (sh *ServerHandler) CreateServer(server *server.Server) error {
	err := sh.server_repo.Create(server)

	if err != nil {
		return err
	}

	return nil
}

func (sh *ServerHandler) FindServer(id string) (*server.Server, error) {
	result, err := sh.server_repo.Find(id)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (sh *ServerHandler) UpdateServer(id string, server *server.Server) error {
	err := sh.server_repo.Update(id, server)

	if err != nil {
		return err
	}

	return nil
}

func (sh *ServerHandler) DeleteServer(id string) error {
	err := sh.server_repo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
