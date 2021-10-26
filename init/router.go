package init

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"my_goods/internal/dish"
	"my_goods/internal/goods"
	"my_goods/internal/list"
)

func Router(db *gorm.DB, router *gin.Engine) *gin.Engine {
	dishRepo := dish.NewDishRepo(db)
	goodsRepo := goods.NewGoodsRepo(db)
	listRepo := list.NewListRepo(db)

	dishService := dish.NewDishService(dishRepo)
	goodsService := goods.NewGoodsService(goodsRepo)
	listService := list.NewListService(listRepo)

	dishHandler := dish.NewDishHandler(dishService)
	goodsHandler := goods.NewGoodsHandler(goodsService)
	listHandler := list.NewListHandler(listService)

	dishHandler.RegisterRoutes(router)
	goodsHandler.RegisterRoutes(router)
	listHandler.RegisterRoutes(router)
	return router
}
