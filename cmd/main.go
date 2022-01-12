package main

import (
	"log"
	"my_goods/internal/delivery/web"
	"my_goods/internal/repository"
	"my_goods/internal/service"
	"my_goods/pkg/db"
	"my_goods/pkg/env"
)

func main() {
	err := env.InitEnvVariables()
	database, err := db.NewDatabasePostgres()

	repo := repository.NewRepository(database)
	services := service.NewService(repo)
	handler := web.NewAPIHandler(services)

	server := handler.InitHTTPHandler()
	if err = server.Run(); err != nil {
		log.Fatalf("error occurred while running http server: %s\n", err.Error())
	}
}
