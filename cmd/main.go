package main

import (
	"my_goods/cmd/app"
)

func main() {
	application := app.NewApplication()
	application.InitApp()
	application.Run()
}
