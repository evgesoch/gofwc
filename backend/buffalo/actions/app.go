package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"

	contenttype "github.com/gobuffalo/mw-contenttype"
	"github.com/gobuffalo/x/sessions"
	"github.com/rs/cors"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
//
// Routing, middleware, groups, etc... are declared TOP -> DOWN.
// This means if you add a middleware to `app` *after* declaring a
// group, that group will NOT have that new middleware. The same
// is true of resource declarations as well.
//
// It also means that routes are checked in the order they are declared.
// `ServeFiles` is a CATCH-ALL route, so it should always be
// placed last in the route declarations, as it will prevent routes
// declared after it to never be called.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:          ENV,
			SessionStore: sessions.Null{},
			PreWares: []buffalo.PreWare{
				cors.Default().Handler,
			},
			SessionName: "_buffalo_session",
		})

		// Set the request content type to JSON
		app.Use(contenttype.Set("application/json"))

		// Api
		app.GET("/posts", GetAllPosts)
		app.GET("/posts/{postID}", GetPost)
		app.POST("/posts", CreatePost)
		app.PUT("/posts/{postID}", UpdatePost)
		app.DELETE("/posts/{postID}", DeletePost)

		// Frontend
		app.ServeFiles("/speak4env/frontend", http.Dir("frontend"))
		app.GET("/speak4env", GetIndexPage)
	}

	return app
}
