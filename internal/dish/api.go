package dish

import (
	"github.com/gin-gonic/gin"
	"my_goods/internal/entity"
	"my_goods/pkg/logger"
	"net/http"
	"strconv"
)

type Handler struct {
	services ServeDish
}

func NewDishHandler(services ServeDish) *Handler {
	return &Handler{services: services}
}

func (h *Handler) GetDish(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"ERROR": err.Error()})
		return
	}
	dish := h.services.GetDish(id)
	if dish.ID == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, *dish)
}

func (h *Handler) GetAllDishes(c *gin.Context) {
	allGoods := h.services.GetAllDishes()
	c.JSON(http.StatusOK, *allGoods)
}

func (h *Handler) CreateDish(c *gin.Context) {
	var dish entity.Dish
	err := c.BindJSON(&dish)
	dishes := h.services.CreateDish(&dish)
	if err != nil {
		logger.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"ERROR": err.Error()})
		return
	}
	c.JSON(http.StatusOK, *dishes)
}

func (h *Handler) UpdateDish(c *gin.Context) {
	dish := entity.Dish{}
	err := c.BindJSON(&dish)
	dishes := h.services.UpdateDish(&dish)
	if err != nil {
		logger.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"ERROR": err.Error()})
		return
	}
	c.JSON(http.StatusOK, *dishes)
}

func (h *Handler) DeleteDish(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	h.services.DeleteDish(id)
	if err != nil {
		logger.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"ERROR": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
