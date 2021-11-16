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

func (h *Handler) GetList(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	list := h.services.getList(id)
	if err != nil {
		logger.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"ERROR": err.Error()})
	}
	c.JSON(http.StatusOK, *list)
}

func (h *Handler) GetAllLists(c *gin.Context) {
	allList := h.services.getAllLists()
	c.JSON(http.StatusOK, *allList)
}

func (h *Handler) CreateList(c *gin.Context) {
	list := entity.List{}
	err := c.Bind(&list)
	lists := h.services.createList(&list)
	if err != nil {
		logger.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"ERROR": err.Error()})
	}
	c.JSON(http.StatusOK, *lists)
}

func (h *Handler) UpdateList(c *gin.Context) {
	list := entity.List{}
	err := c.Bind(&list)
	lists := h.services.updateList(&list)
	if err != nil {
		logger.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"ERROR": err.Error()})
	}
	c.JSON(http.StatusOK, *lists)
}

func (h *Handler) DeleteList(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	h.services.deleteList(id)
	if err != nil {
		logger.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"ERROR": err.Error()})
	}
	c.Status(http.StatusOK)
}
