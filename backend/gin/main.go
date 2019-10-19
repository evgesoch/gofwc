package main

import (
	ginRouter "github.com/evgesoch/gofwc/backend/gin/routers"
)

func main() {
	r := ginRouter.SetupRouter()

	r.Run(":8080")
}
