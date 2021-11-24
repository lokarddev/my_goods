package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"log"
	"my_goods/cmd/server"
	"my_goods/pkg/db"
	"my_goods/pkg/environ"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	environ.Env()
	database, err := db.DB(db.NewDatabaseConf())

	if err != nil {
		logrus.Fatalf("Error while running server")
	}

	srv := new(server.Server)
	go func() {
		if err = srv.Run(environ.Port, server.Router(database)); err != nil {
			log.Fatalf("Error occured while: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err = srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error while server shutting down %s", err.Error())
	}
}
