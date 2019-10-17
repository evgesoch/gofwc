package main

import (
	"github.com/astaxie/beego"
	"github.com/evgesoch/gofwc/backend/beego/models"

	_ "github.com/evgesoch/gofwc/backend/beego/routers"
)

func main() {
	// Open the main database
	models.OpenDB()

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
