package main

import (
	"fmt"

	"github.com/teguhbudhi13/products/app"
	"github.com/teguhbudhi13/products/config"
)

func main() {
	fmt.Println("Your API Now Is Published")
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":" + config.Server.Port)
}
