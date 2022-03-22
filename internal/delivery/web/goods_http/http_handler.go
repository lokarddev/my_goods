package goods_http

import (
	"github.com/gin-gonic/gin"
	"my_goods/internal/delivery"
)

type GoodsHttpHandler struct {
	service delivery.GoodsServiceInterface
}

func NewGoodsHttpHandler(service delivery.GoodsServiceInterface) *GoodsHttpHandler {
	return &GoodsHttpHandler{service: service}
}

func (h *GoodsHttpHandler) RegisterRoutes(api *gin.RouterGroup) {

}
