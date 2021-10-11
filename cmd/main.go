package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"my_goods/configs"
	"my_goods/internal/handlers"
	"my_goods/internal/repos"
	"my_goods/internal/services"
)

func main() {
	configs.InitEnv()
	db := configs.InitDB(configs.NewDatabaseConf())

	repo := repos.NewRepository(db)
	service := services.NewService(repo)
	handler := handlers.NewHandler(service)

	server := handler.InitRoutes()
	err := server.Run(fmt.Sprintf("%s:%s", configs.Host, configs.Port))
	if err != nil {
		logrus.Fatalf("Error while running server")
	}
}
