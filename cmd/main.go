package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	initial "my_goods/init"
)

func main() {
	initial.Env()
	db := initial.InitDB(initial.NewDatabaseConf())
	handler := gin.New()
	handler.Use(gin.Logger())
	server := initial.Router(db, handler)
	err := server.Run(fmt.Sprintf("%s:%s", initial.Host, initial.Port))
	if err != nil {
		logrus.Fatalf("Error while running server")
	}
}
