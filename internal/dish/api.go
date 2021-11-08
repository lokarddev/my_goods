package dish

import (
	"github.com/gin-gonic/gin"
	"my_goods/internal/entity"
	"my_goods/pkg/logger"
	"net/http"
	"strconv"
)

type HandleDish interface {
	GetAllDishes(c *gin.Context)
}

type Handler struct {
	services ServeDish
}

func NewDishHandler(services ServeDish) *Handler {
	return &Handler{services: services}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		dish := api.Group("/dish")
		{
			dish.GET("/:id", h.GetDish)
			dish.GET("/", h.GetAllDishes)
			dish.POST("/", h.CreateDish)
			dish.PUT("/", h.UpdateDish)
			dish.DELETE("/:id", h.DeleteDish)
		}
	}
}

func (h *Handler) GetDish(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	dish := h.services.GetDish(id)
	if err != nil {
		logger.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"ERROR": err.Error()})
	}
	c.JSON(http.StatusOK, *dish)
}

func (h *Handler) GetAllDishes(c *gin.Context) {
	allGoods := h.services.GetAllDishes()
	c.JSON(http.StatusOK, *allGoods)
}

func (h *Handler) CreateDish(c *gin.Context) {
	dish := entity.Dish{}
	err := c.Bind(&dish)
	dishes := h.services.CreateDish(&dish)
	if err != nil {
		logger.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"ERROR": err.Error()})
	}
	c.JSON(http.StatusOK, *dishes)
}

func (h *Handler) UpdateDish(c *gin.Context) {
	dish := entity.Dish{}
	err := c.Bind(&dish)
	dishes := h.services.UpdateDish(&dish)
	if err != nil {
		logger.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"ERROR": err.Error()})
	}
	c.JSON(http.StatusOK, *dishes)
}

func (h *Handler) DeleteDish(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	h.services.DeleteDish(id)
	if err != nil {
		logger.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"ERROR": err.Error()})
	}
	c.Status(http.StatusOK)
}
