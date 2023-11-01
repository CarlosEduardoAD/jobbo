package routes

import (
	"fmt"
	"log"

	message_handler "github.com/CarlosEduardoAD/jobbo-api/internal/api/domain/handlers"
	server_handler "github.com/CarlosEduardoAD/jobbo-api/internal/api/domain/handlers"
	"github.com/CarlosEduardoAD/jobbo-api/internal/api/domain/model/message"
	message_repo "github.com/CarlosEduardoAD/jobbo-api/internal/api/infra/repo/persistence/message"
	server_repo "github.com/CarlosEduardoAD/jobbo-api/internal/api/infra/repo/persistence/server"
	"github.com/CarlosEduardoAD/jobbo-api/internal/api/utils"
	"github.com/google/uuid"

	"github.com/labstack/echo/v4"
)

func MessageRoutes(echo *echo.Echo) {
	serverGroup := echo.Group("/message")
	serverGroup.POST("/create", CreateMessage)
	serverGroup.GET("/:id", findMessage)
	serverGroup.PUT("/:id", updateMessage)
	serverGroup.DELETE("/:id", deleteMessage)
	serverGroup.GET("/org/:orgId", findMessageWithOrganizationId)
	serverGroup.GET("/campaign/:campaignId", findMessageWithCampaignId)
}

func CreateMessage(c echo.Context) error {
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

	serverPayload := new(message.Message)
	if err = c.Bind(serverPayload); err != nil {
		return c.JSON(400, "invalid payload")
	}

	err = serverPayload.Validate()

	log.Default().Println(serverPayload)

	if err != nil {
		return c.JSON(500, fmt.Sprintf("validtion amogus failed: %s", err))
	}

	server_rep := server_repo.NewServerRepository(db)
	server_h := server_handler.NewServerHandler(server_rep)

	messageServer, err := server_h.FindServer(serverPayload.ServerID)

	log.Println("ACHEI O SERVIDOR: ", messageServer)

	if err != nil {
		log.Print(err)
		return c.JSON(500, err)
	}

	log.Println(messageServer)

	messageToBeSaved, err := message.NewMessage(uuid.New().String(), serverPayload.UserID, serverPayload.OrganizationId, serverPayload.CampaignId, serverPayload.From, serverPayload.To, serverPayload.Subject, serverPayload.Body, serverPayload.SentAt, messageServer.ID, *messageServer)

	if err != nil {
		log.Print(err)
		return c.JSON(500, err)
	}

	message_repo := message_repo.NewMessageRepository(db)
	message_handler := message_handler.NewMessageHandler(message_repo)

	err = message_handler.CreateMessage(messageToBeSaved)

	if err != nil {
		return c.JSON(500, err)
	}

	return nil
}

func findMessage(c echo.Context) error {
	var err error
	messageId := c.Param("id")

	db, err := utils.ConnectDB()

	if err != nil {
		return c.JSON(500, err)
	}

	dbClose, err := db.DB()

	if err != nil {
		return c.JSON(500, "an error ocurred during db setup")
	}

	defer dbClose.Close()

	message_repo := message_repo.NewMessageRepository(db)
	message_handler := message_handler.NewMessageHandler(message_repo)

	result, err := message_handler.FindMessage(messageId)

	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, result)
}

func updateMessage(c echo.Context) error {
	var err error

	messageId := c.Param("id")

	db, err := utils.ConnectDB()

	if err != nil {
		return c.JSON(500, err)
	}

	dbClose, err := db.DB()

	if err != nil {
		return c.JSON(500, "an error ocurred during db setup")
	}

	defer dbClose.Close()

	serverPayload := new(message.Message)
	if err = c.Bind(serverPayload); err != nil {
		return c.JSON(400, "invalid payload")
	}

	err = serverPayload.Validate()

	if err != nil {
		return c.JSON(500, fmt.Sprintf("validtion failed: %s", err))
	}

	messageToBeSaved, err := message.NewMessage(messageId, serverPayload.UserID, serverPayload.OrganizationId, serverPayload.CampaignId, serverPayload.From, serverPayload.To, serverPayload.Subject, serverPayload.Body, serverPayload.SentAt, serverPayload.ServerID, serverPayload.Server)

	if err != nil {
		return c.JSON(500, err)
	}

	message_repo := message_repo.NewMessageRepository(db)
	message_handler := message_handler.NewMessageHandler(message_repo)

	err = message_handler.UpdateMessage(messageId, messageToBeSaved)

	if err != nil {
		return c.JSON(500, err)
	}

	return nil
}

func deleteMessage(c echo.Context) error {
	var err error
	messageId := c.Param("id")

	db, err := utils.ConnectDB()

	if err != nil {
		return c.JSON(500, err)
	}

	dbClose, err := db.DB()

	if err != nil {
		return c.JSON(500, "an error ocurred during db setup")
	}

	defer dbClose.Close()

	message_repo := message_repo.NewMessageRepository(db)
	message_handler := message_handler.NewMessageHandler(message_repo)

	err = message_handler.DeleteMessage(messageId)

	if err != nil {
		return c.JSON(500, err)
	}

	return nil
}

func findMessageWithOrganizationId(c echo.Context) error {
	var err error
	organizationId := c.Param("orgId")

	db, err := utils.ConnectDB()

	if err != nil {
		return c.JSON(500, err)
	}

	dbClose, err := db.DB()

	if err != nil {
		return c.JSON(500, "an error ocurred during db setup")
	}

	defer dbClose.Close()

	message_repo := message_repo.NewMessageRepository(db)
	message_handler := message_handler.NewMessageHandler(message_repo)

	result, err := message_handler.FindMessageWithOrganizationId(organizationId)

	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, result)
}

func findMessageWithCampaignId(c echo.Context) error {
	var err error
	campaignId := c.Param("campaignId")

	db, err := utils.ConnectDB()

	if err != nil {
		return c.JSON(500, err)
	}

	dbClose, err := db.DB()

	if err != nil {
		return c.JSON(500, "an error ocurred during db setup")
	}

	defer dbClose.Close()

	message_repo := message_repo.NewMessageRepository(db)
	message_handler := message_handler.NewMessageHandler(message_repo)

	result, err := message_handler.FindMessageWithCampaignId(campaignId)

	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, result)
}
