package list

import (
	"github.com/gin-gonic/gin"
	"my_goods/internal/entity"
	"my_goods/pkg/logger"
	"net/http"
	"strconv"
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
	id, err := strconv.Atoi(c.Param("id"))
	list := h.services.getList(id)
	if err != nil {
		logger.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"ERROR": err.Error()})
	}
	c.JSON(http.StatusOK, *list)
}

func (h *Handler) getAllLists(c *gin.Context) {
	allList := h.services.getAllLists()
	c.JSON(http.StatusOK, *allList)
}

func (h *Handler) createList(c *gin.Context) {
	list := entity.List{}
	err := c.Bind(&list)
	lists := h.services.createList(&list)
	if err != nil {
		logger.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"ERROR": err.Error()})
	}
	c.JSON(http.StatusOK, *lists)
}

func (h *Handler) updateList(c *gin.Context) {
	list := entity.List{}
	err := c.Bind(&list)
	lists := h.services.updateList(&list)
	if err != nil {
		logger.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"ERROR": err.Error()})
	}
	c.JSON(http.StatusOK, *lists)
}

func (h *Handler) deleteList(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	h.services.deleteList(id)
	if err != nil {
		logger.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"ERROR": err.Error()})
	}
	c.Status(http.StatusOK)
}
