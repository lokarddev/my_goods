package main

import (
	"github.com/sirupsen/logrus"
	"log"
	"my_goods/internal/delivery/web"
	"my_goods/internal/repository"
	"my_goods/internal/service"
	"my_goods/pkg/db"
	"my_goods/pkg/environ"
)

func main() {
	environ.Env()
	database, err := db.DB(db.NewDatabaseConf())
	if err != nil {
		logrus.Fatalf("Error while running server")
	}

	repo := repository.NewRepository(database)
	services := service.NewService(repo)
	handler := web.NewAPIHandler(services)

	server := handler.InitHTTPHandler()
	if err = server.Run(); err != nil {
		log.Fatalf("error occurred while running http server: %s\n", err.Error())
	}
}
