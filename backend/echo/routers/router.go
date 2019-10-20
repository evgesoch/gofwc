package routers

import (
	echoControllers "github.com/evgesoch/gofwc/backend/echo/controllers"
	"github.com/labstack/echo"
)

// Setup the routes for the echo application
func SetupRoutes(e *echo.Echo) {
	e.GET("posts", echoControllers.GetAllPosts())
	e.GET("posts/:postID", echoControllers.GetPost())
	e.POST("posts", echoControllers.CreatePost())
	e.PUT("posts/:postID", echoControllers.UpdatePost())
	e.DELETE("posts/:postID", echoControllers.DeletePost())
}
