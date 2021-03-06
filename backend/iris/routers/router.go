package routers

import (
	irisControllers "github.com/evgesoch/gofwc/backend/iris/controllers"
	"github.com/kataras/iris"
)

func SetupRoutes(app *iris.Application) {
	// Api
	app.Get("/posts", irisControllers.GetAllPosts())
	app.Get("/posts/{postID}", irisControllers.GetPost())
	app.Post("/posts", irisControllers.CreatePost())
	app.Put("/posts/{postID}", irisControllers.UpdatePost())
	app.Delete("/posts/{postID}", irisControllers.DeletePost())

	// Frontend
	app.HandleDir("/frontend", "../../frontend")
	app.Get("/speak4env", irisControllers.GetIndexPage())
}
