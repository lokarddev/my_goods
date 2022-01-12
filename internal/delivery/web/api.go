package web

import (
	"github.com/gin-gonic/gin"
	"my_goods/internal/service"
)

type BaseHTTPHandler struct {
	services *service.Service
}

func (h *BaseHTTPHandler) InitHTTPHandler() *gin.Engine {
	handler := gin.New()
	handler.Use(gin.Logger(), gin.Recovery())
	h.initAPI(handler)
	return handler
}

func (h *BaseHTTPHandler) initAPI(router *gin.Engine) {
	controller := NewController(h.services)
	api := router.Group("/api")
	{
		controller.RegisterRoutes(api)
	}
}

func NewAPIHandler(services *service.Service) *BaseHTTPHandler {
	return &BaseHTTPHandler{services: services}
}
