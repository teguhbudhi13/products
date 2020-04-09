package main

import (
	"github.com/teguhbudhi13/products/app"
	"github.com/teguhbudhi13/products/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":" + config.Server.Port)
}
