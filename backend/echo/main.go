package main

import (
	echoRouter "github.com/evgesoch/gofwc/backend/echo/routers"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	echoRouter.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
