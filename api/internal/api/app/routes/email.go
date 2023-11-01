package routes

import (
	"net/http"
	"strconv"

	email_handler "github.com/CarlosEduardoAD/jobbo-api/internal/api/domain/handlers"
	server_handler "github.com/CarlosEduardoAD/jobbo-api/internal/api/domain/handlers"
	"github.com/CarlosEduardoAD/jobbo-api/internal/api/domain/model/email"
	server_repo "github.com/CarlosEduardoAD/jobbo-api/internal/api/infra/repo/persistence/server"
	email_repo "github.com/CarlosEduardoAD/jobbo-api/internal/api/infra/repo/smtp"
	"github.com/CarlosEduardoAD/jobbo-api/internal/api/utils"
	"github.com/labstack/echo/v4"
)

// EmailRoutes creates a group route for the route email.
//
// router: The router instance to attach the group route to.
//
// Returns the created group route.
func EmailRoutes(router *echo.Echo) *echo.Group {
	// I want to create a group route for the route email
	emailGroup := router.Group("/email")
	emailGroup.POST("/", sendEmail)

	return emailGroup
}

// The function takes a `c` of type `echo.Context` as a parameter.
// It returns an `error` if any error occurs during the process.
func sendEmail(c echo.Context) error {
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
	server, err := server_handler.FindServer(serverId)

	if err != nil {
		c.Error(err)
		return err
	}

	port, err := strconv.Atoi(server.Port)

	if err != nil {
		c.Error(err)
		return err
	}

	dialer := utils.ConnectSMTP(server.Address, port, server.User, server.Password)

	e := new(email.Email) // body parsing
	if err = c.Bind(e); err != nil {
		c.Error(err)
		return err
	}

	err = e.Validate()

	if err != nil {
		c.Error(err)
		return err
	}

	emailToBeDelivered := utils.ConvertToMailMessage(email.NewEmail(e.From, e.To, e.Subject, e.Body, e.ServerID))
	emailRepo := email_repo.NewEmailService(dialer)
	email_handler := email_handler.NewEmailHandler(emailRepo)

	err, _ = email_handler.DeliverEmail(emailToBeDelivered)

	if err != nil {
		c.Error(err)
		return err
	}

	return c.JSON(http.StatusOK, "email sent successfully!")
}
