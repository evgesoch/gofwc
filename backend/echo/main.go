package main

import (
	echoModel "github.com/evgesoch/gofwc/backend/echo/models"
	echoRouter "github.com/evgesoch/gofwc/backend/echo/routers"

	"github.com/labstack/echo"
)

func main() {
	echoModel.OpenDB()

	e := echo.New()

	echoRouter.SetupRoutes(e)

	err := e.Start(":8080")
	if err != nil {
		echoModel.CloseDB()
	}

	e.Logger.Fatal(err)
}
