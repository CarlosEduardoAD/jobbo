package routes

import (
	"net/http"

	email_service "github.com/CarlosEduardoAD/jobbo-api/internal/api/app/services"
	"github.com/CarlosEduardoAD/jobbo-api/internal/api/domain/email"
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

	emailToBeDelivered := email.NewEmail(e.From, e.To, e.Subject, e.Body)

	err, _ = email_service.DeliverEmail(dialer, emailToBeDelivered)

	if err != nil {
		c.Error(err)
		return err
	}

	return c.JSON(http.StatusOK, "email sent successfully!")
}