package main

import (
	"my_goods/cmd/app"
)

// @title        My goods
// @version      1.0
// @description  Simple shopping list app

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization

func main() {
	application := app.NewApplication()
	application.InitApp()
	application.Run()
}
