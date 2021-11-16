package goods

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

func NewGoodsHandler(services *Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) GetGoods(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	goods := h.services.getGoods(id)
	if err != nil {
		logger.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"ERROR": err.Error()})
	}
	c.JSON(http.StatusOK, *goods)
}

func (h *Handler) GetAllGoods(c *gin.Context) {
	allGoods := h.services.getAllGoods()
	c.JSON(http.StatusOK, *allGoods)
}

func (h *Handler) CreateGoods(c *gin.Context) {
	good := entity.Goods{}
	err := c.Bind(&good)
	goods := h.services.createGoods(&good)
	if err != nil {
		logger.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"ERROR": err.Error()})
	}
	c.JSON(http.StatusOK, *goods)
}

func (h *Handler) UpdateGoods(c *gin.Context) {
	good := entity.Goods{}
	err := c.Bind(&good)
	goods := h.services.updateGoods(&good)
	if err != nil {
		logger.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"ERROR": err.Error()})
	}
	c.JSON(http.StatusOK, *goods)
}

func (h *Handler) DeleteGoods(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	h.services.deleteGoods(id)
	if err != nil {
		logger.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"ERROR": err.Error()})
	}
	c.Status(http.StatusOK)
}
