package dish_http

import (
	"github.com/gin-gonic/gin"
	"my_goods/internal/delivery"
)

type DishHttpHandler struct {
	service delivery.DishServiceInterface
}

func NewDishHttpHandler(service delivery.DishServiceInterface) *DishHttpHandler {
	return &DishHttpHandler{service: service}
}

func (h *DishHttpHandler) RegisterRoutes(api *gin.RouterGroup) {

}
