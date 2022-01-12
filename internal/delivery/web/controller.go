package web

import (
	"github.com/gin-gonic/gin"
	"my_goods/internal/service"
	"net/http"
)

const (
	someCommand = "test command"
)

type Controller struct {
	services *service.Service
}

func (c *Controller) RegisterRoutes(api *gin.RouterGroup) {
	api.GET("/health-check", c.healthCheck)
}

func NewController(services *service.Service) *Controller {
	return &Controller{services: services}
}

func (c *Controller) healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}

type DispatcherInterface interface {
}

type Dispatcher struct {
}

func (d *Dispatcher) MatchCommands() {

}
