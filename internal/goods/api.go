package goods

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type GoodsHandler struct {
	services *GoodsService
}

func NewGoodsHandler(services *GoodsService) *GoodsHandler {
	return &GoodsHandler{services: services}
}

func (h *GoodsHandler) RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		goods := api.Group("/goods")
		{
			goods.GET("/:id", h.GetGoods)
			goods.GET("/", h.GetAllGoods)
			goods.POST("/", h.CreateGoods)
			goods.PUT("/:id", h.UpdateGoods)
			goods.DELETE("/:id", h.DeleteGoods)
		}
	}
}

func (h *GoodsHandler) GetGoods(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}

func (h *GoodsHandler) GetAllGoods(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}

func (h *GoodsHandler) CreateGoods(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}

func (h *GoodsHandler) UpdateGoods(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}

func (h *GoodsHandler) DeleteGoods(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}
