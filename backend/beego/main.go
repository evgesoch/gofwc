package main

import (
	"os"
	"os/signal"
	"syscall"

	beegoModel "github.com/evgesoch/gofwc/backend/beego/models"
	_ "github.com/evgesoch/gofwc/backend/beego/routers"

	"github.com/astaxie/beego"
)

func main() {
	beegoModel.OpenDB()

	// Run the application
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()

	handleServerTermination()
}

// Close the database after server termination
func handleServerTermination() {
	c := make(chan os.Signal, 2)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		beegoModel.CloseDB()
		os.Exit(0)
	}()
}
