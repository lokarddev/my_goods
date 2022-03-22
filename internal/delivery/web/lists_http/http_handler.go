package lists_http

import (
	"github.com/gin-gonic/gin"
	"my_goods/internal/delivery"
)

type ListsHttpHandler struct {
	service delivery.ListServiceInterface
}

func NewListsHttpHandler(service delivery.ListServiceInterface) *ListsHttpHandler {
	return &ListsHttpHandler{service: service}
}

func (h *ListsHttpHandler) RegisterRoutes(api *gin.RouterGroup) {

}
