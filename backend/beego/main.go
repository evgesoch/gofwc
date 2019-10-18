package main

import (
	"github.com/astaxie/beego"

	_ "github.com/evgesoch/gofwc/backend/beego/routers"
)

func main() {
	// Run the application
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
