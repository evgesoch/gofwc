package main

import (
	"github.com/evgesoch/gofwc/backend/beego/models"

	_ "github.com/evgesoch/gofwc/backend/beego/routers"
)

func main() {

	// Create the main database
	models.CreateDB()

	/*if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()*/
}
