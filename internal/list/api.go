package list

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ListHandler struct {
	services *ListService
}

func NewListHandler(services *ListService) *ListHandler {
	return &ListHandler{services: services}
}

func (h *ListHandler) RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		list := api.Group("/list")
		{
			list.GET("/:id", h.GetList)
			list.GET("/", h.GetAllLists)
			list.POST("/", h.CreateList)
			list.PUT("/:id", h.UpdateList)
			list.DELETE("/:id", h.DeleteList)
		}
	}
}

func (h *ListHandler) GetList(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}

func (h *ListHandler) GetAllLists(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}

func (h *ListHandler) CreateList(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}

func (h *ListHandler) UpdateList(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}

func (h *ListHandler) DeleteList(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}
