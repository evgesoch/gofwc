package main

import (
	irisModel "github.com/evgesoch/gofwc/backend/iris/models"
	irisRouter "github.com/evgesoch/gofwc/backend/iris/routers"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	irisModel.OpenDB()

	app := iris.New()
	app.Logger().SetLevel("debug")
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	irisRouter.SetupRoutes(app)

	err := app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
	if err != nil {
		irisModel.CloseDB()
	}
}
