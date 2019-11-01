package routers

import (
	ginControllers "github.com/evgesoch/gofwc/backend/gin/controllers"
	"github.com/gin-gonic/gin"
)

var r = gin.Default()

// Main Router for gin application
func SetupRouter() *gin.Engine {
	// Api
	r.GET("/posts", ginControllers.GetAllPosts())
	r.GET("/posts/:postID", ginControllers.GetPost())
	r.POST("/posts", ginControllers.CreatePost())
	r.PUT("/posts/:postID", ginControllers.UpdatePost())
	r.DELETE("/posts/:postID", ginControllers.DeletePost())

	// Frontend
	r.Static("/frontend", "../../frontend")
	r.StaticFile("/speak4env", "../../frontend/index.html")

	return r
}
