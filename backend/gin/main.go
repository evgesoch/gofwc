package main

import (
	ginModel "github.com/evgesoch/gofwc/backend/gin/models"
	ginRouter "github.com/evgesoch/gofwc/backend/gin/routers"
)

func main() {
	ginModel.OpenDB()

	r := ginRouter.SetupRouter()

	err := r.Run(":8080")
	if err != nil {
		ginModel.CloseDB()
		return
	}
}
