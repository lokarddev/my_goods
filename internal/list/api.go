package list

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	services *Service
}

func NewListHandler(services *Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		list := api.Group("/list")
		{
			list.GET("/:id", h.getList)
			list.GET("/", h.getAllLists)
			list.POST("/", h.createList)
			list.PUT("/:id", h.updateList)
			list.DELETE("/:id", h.deleteList)
		}
	}
}

func (h *Handler) getList(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}

func (h *Handler) getAllLists(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}

func (h *Handler) createList(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}

func (h *Handler) updateList(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}

func (h *Handler) deleteList(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}
