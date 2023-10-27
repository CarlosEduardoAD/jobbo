package routes

import (
	"fmt"
	"log"

	server_handler "github.com/CarlosEduardoAD/jobbo-api/internal/api/domain/handlers"

	"github.com/CarlosEduardoAD/jobbo-api/internal/api/domain/model/server"
	server_repo "github.com/CarlosEduardoAD/jobbo-api/internal/api/infra/repo/persistence/server"
	"github.com/CarlosEduardoAD/jobbo-api/internal/api/utils"
	"github.com/labstack/echo/v4"
)

func ServerRoutes(echo *echo.Echo) *echo.Group {
	serverGroup := echo.Group("/server")
	serverGroup.POST("/", createServer)
	serverGroup.GET("/:id", findServer)
	serverGroup.PUT("/:id", updateServer)
	serverGroup.DELETE("/:id", deleteServer)

	return serverGroup
}

func createServer(c echo.Context) error {
	var err error

	db, err := utils.ConnectDB()

	if err != nil {
		return c.JSON(500, err)
	}

	dbClose, err := db.DB()

	if err != nil {
		return c.JSON(500, "an error ocurred during db setup")
	}

	defer dbClose.Close()

	if err != nil {
		return c.JSON(500, err)
	}

	serverPayload := new(server.Server)
	if err = c.Bind(serverPayload); err != nil {
		return c.JSON(400, "invalid payload")
	}

	err = serverPayload.Validate()

	if err != nil {
		return c.JSON(500, fmt.Sprintf("validtion failed: %s", err))
	}

	serverToBeSaved := server.NewServer(serverPayload.Name, serverPayload.Address, serverPayload.Port, serverPayload.User, serverPayload.Password, serverPayload.Cryptography)
	server_repo := server_repo.NewServerRepository(db)
	server_handler := server_handler.NewServerHandler(server_repo)

	err = server_handler.CreateServer(serverToBeSaved)

	if err != nil {
		return c.JSON(500, err)
	}

	return nil
}

func findServer(c echo.Context) error {
	var err error
	serverId := c.Param("id")

	db, err := utils.ConnectDB()

	if err != nil {
		return c.JSON(500, err)
	}

	dbClose, err := db.DB()

	if err != nil {
		return c.JSON(500, "an error ocurring during db setup")
	}

	defer dbClose.Close()

	server_repo := server_repo.NewServerRepository(db)
	server_handler := server_handler.NewServerHandler(server_repo)
	result, err := server_handler.FindServer(serverId)

	if err != nil {
		log.Print("Meu erro: ", err)
		c.Error(err)
		return err
	}

	return c.JSON(201, result)
}

func updateServer(c echo.Context) error {

	var err error
	serverId := c.Param("id")

	db, err := utils.ConnectDB()

	if err != nil {
		c.Error(err)
		return err
	}

	dbClose, err := db.DB()

	if err != nil {
		c.Error(err)
		return err
	}

	defer dbClose.Close()

	serverPayload := new(server.Server)
	if err = c.Bind(serverPayload); err != nil {
		c.Error(err)
		return err
	}

	err = serverPayload.Validate()

	if err != nil {
		c.Error(err)
		return err
	}

	server_repo := server_repo.NewServerRepository(db)

	server_handler := server_handler.NewServerHandler(server_repo)
	err = server_handler.UpdateServer(serverId, serverPayload)

	if err != nil {
		c.Error(err)
		return err
	}

	return nil
}

func deleteServer(c echo.Context) error {
	var err error
	serverId := c.Param("id")

	db, err := utils.ConnectDB()

	if err != nil {
		c.Error(err)
		return err
	}

	dbClose, err := db.DB()

	if err != nil {
		c.Error(err)
		return err
	}

	defer dbClose.Close()

	server_repo := server_repo.NewServerRepository(db)
	server_handler := server_handler.NewServerHandler(server_repo)
	err = server_handler.DeleteServer(serverId)

	if err != nil {
		c.Error(err)
		return err
	}

	return nil
}
