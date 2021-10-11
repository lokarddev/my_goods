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
			dish.GET("/", h.Test)
		}
	}
	return router
}
