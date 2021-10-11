package handlers

import (
	"github.com/gin-gonic/gin"
	"my_goods/internal/services"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	api := router.Group("/api")
	{
		dish := api.Group("/dish")
		{
			dish.GET("/:id", h.GetDish)
			dish.GET("/", h.GetAllDishes)
			dish.POST("/", h.CreateDish)
			dish.PUT("/:id", h.UpdateDish)
			dish.DELETE("/:id", h.DeleteDish)
		}
		goods := api.Group("/goods")
		{
			goods.GET("/:id", h.GetGoods)
			goods.GET("/", h.GetAllGoods)
			goods.POST("/", h.CreateGoods)
			goods.PUT("/:id", h.UpdateGoods)
			goods.DELETE("/:id", h.DeleteList)
		}
		list := api.Group("/list")
		{
			list.GET("/:id", h.GetList)
			list.GET("/", h.GetAllLists)
			list.POST("/", h.CreateList)
			list.PUT("/:id", h.UpdateList)
			list.DELETE("/:id", h.DeleteList)
		}
	}
	return router
}
