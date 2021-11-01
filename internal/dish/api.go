package dish

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

func NewDishHandler(services *Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		dish := api.Group("/dish")
		{
			dish.GET("/:id", h.getDish)
			dish.GET("/", h.getAllDishes)
			dish.POST("/", h.createDish)
			dish.PUT("/", h.updateDish)
			dish.DELETE("/:id", h.deleteDish)
		}
	}
}

func (h *Handler) getDish(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	dish := h.services.getDish(id)
	if err != nil {
		logger.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"ERROR": err.Error()})
	}
	c.JSON(http.StatusOK, *dish)
}

func (h *Handler) getAllDishes(c *gin.Context) {
	allGoods := h.services.getAllDishes()
	c.JSON(http.StatusOK, *allGoods)
}

func (h *Handler) createDish(c *gin.Context) {
	dish := entity.Dish{}
	err := c.Bind(&dish)
	dishes := h.services.createDish(&dish)
	if err != nil {
		logger.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"ERROR": err.Error()})
	}
	c.JSON(http.StatusOK, *dishes)
}

func (h *Handler) updateDish(c *gin.Context) {
	dish := entity.Dish{}
	err := c.Bind(&dish)
	dishes := h.services.updateDish(&dish)
	if err != nil {
		logger.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"ERROR": err.Error()})
	}
	c.JSON(http.StatusOK, *dishes)
}

func (h *Handler) deleteDish(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	h.services.deleteDish(id)
	if err != nil {
		logger.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"ERROR": err.Error()})
	}
	c.Status(http.StatusOK)
}
