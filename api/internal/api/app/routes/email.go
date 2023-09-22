package routes

import (
	"net/http"

	email_handler "github.com/CarlosEduardoAD/jobbo-api/internal/api/app/handlers"
	"github.com/CarlosEduardoAD/jobbo-api/internal/api/domain/email"
	email_repo "github.com/CarlosEduardoAD/jobbo-api/internal/api/infra/repo/smtp"
	"github.com/CarlosEduardoAD/jobbo-api/internal/api/utils"
	"github.com/labstack/echo"
)

func EmailRoutes(router *echo.Echo) *echo.Group {
	// I want to create a group route for the route email
	emailGroup := router.Group("/email")
	emailGroup.POST("/", sendEmail)

	return emailGroup
}

func sendEmail(c echo.Context) error {
	var err error

	dialer := utils.ConnectSMTP("smtp.gmail.com", 587, "karl.devcontato@gmail.com", "ehuf hvxx funu frov")

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

	emailToBeDelivered := utils.ConvertToMailMessage(email.NewEmail(e.From, e.To, e.Subject, e.Body))
	emailRepo := email_repo.NewEmailService(dialer)
	email_handler := email_handler.NewEmailHandler(emailRepo)

	err, _ = email_handler.DeliverEmail(emailToBeDelivered)

	if err != nil {
		c.Error(err)
		return err
	}

	return c.JSON(http.StatusOK, "email sent successfully!")
}
