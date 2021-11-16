package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"my_goods/pkg/db"
	"my_goods/pkg/environ"
)

func main() {
	environ.Env()
	database, err := db.DB(db.NewDatabaseConf())
	handler := gin.New()
	handler.Use(gin.Logger())
	server := Router(database, handler)
	err = server.Run(fmt.Sprintf(":%s", environ.Port))
	if err != nil {
		logrus.Fatalf("Error while running server")
	}
}
