package web

import (
	"github.com/gin-gonic/gin"
	"my_goods/internal/service"
)

type BaseHandler struct {
	services *service.Service
}

func (h *BaseHandler) InitHTTPHandler() *gin.Engine {
	handler := gin.New()
	handler.Use(gin.Logger(), gin.Recovery())
	h.initAPI(handler)
	return handler
}

func (h *BaseHandler) initAPI(router *gin.Engine) {
	controller := NewController(h.services)
	api := router.Group("/api")
	{
		controller.RegisterRoutes(api)
	}
}

func NewAPIHandler(services *service.Service) *BaseHandler {
	return &BaseHandler{services: services}
}
