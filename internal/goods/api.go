package goods

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	services *Service
}

func NewGoodsHandler(services *Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		goods := api.Group("/goods")
		{
			goods.GET("/:id", h.getGoods)
			goods.GET("/", h.getAllGoods)
			goods.POST("/", h.createGoods)
			goods.PUT("/:id", h.updateGoods)
			goods.DELETE("/:id", h.deleteGoods)
		}
	}
}

func (h *Handler) getGoods(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}

func (h *Handler) getAllGoods(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}

func (h *Handler) createGoods(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}

func (h *Handler) updateGoods(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}

func (h *Handler) deleteGoods(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}
