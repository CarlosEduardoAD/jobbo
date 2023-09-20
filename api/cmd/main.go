package main

import (
	"github.com/CarlosEduardoAD/jobbo-api/internal/api/framework/routes"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	routes.EmailRoutes(e)
	e.Logger.Fatal(e.Start(":9292"))
}
