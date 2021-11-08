package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"my_goods/internal/dish"
	"my_goods/internal/goods"
	"my_goods/internal/list"
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

func Router(db *gorm.DB, router *gin.Engine) *gin.Engine {
	dishRepo := dish.NewDishRepo(db)
	goodsRepo := goods.NewGoodsRepo(db)
	listRepo := list.NewListRepo(db)

	dishService := dish.NewDishService(*dishRepo)
	goodsService := goods.NewGoodsService(*goodsRepo)
	listService := list.NewListService(*listRepo)

	dishHandler := dish.NewDishHandler(dishService)
	goodsHandler := goods.NewGoodsHandler(goodsService)
	listHandler := list.NewListHandler(listService)

	dishHandler.RegisterRoutes(router)
	goodsHandler.RegisterRoutes(router)
	listHandler.RegisterRoutes(router)
	return router
}
