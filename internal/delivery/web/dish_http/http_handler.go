package dish_http

import (
	"github.com/gin-gonic/gin"
	"my_goods/internal/delivery"
	"my_goods/internal/entities"
	"net/http"
	"strconv"
)

type DishHttpHandler struct {
	service delivery.DishServiceInterface
}

func NewDishHttpHandler(service delivery.DishServiceInterface) *DishHttpHandler {
	return &DishHttpHandler{service: service}
}

func (h *DishHttpHandler) RegisterRoutes(api *gin.RouterGroup) {
	api.GET("get-dish/:dish_id", h.GetDish)
	api.GET("get-dishes/", h.GetAllDishes)
	api.POST("create-dish/", h.CreateDish)
	api.POST("update-dish/:dish_id", h.UpdateDish)
	api.DELETE("delete-dish/:dish_id", h.DeleteDish)
}

func (h *DishHttpHandler) GetDish(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("dish_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	dish, err := h.service.GetDish(id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, dish)
}

func (h *DishHttpHandler) GetAllDishes(c *gin.Context) {
	dishes, err := h.service.GetAllDishes()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, dishes)
}

func (h *DishHttpHandler) CreateDish(c *gin.Context) {
	dish := entities.Dish{}
	if err := c.BindJSON(&dish); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	res, err := h.service.CreateDish(&dish)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *DishHttpHandler) UpdateDish(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("dish_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	dish := entities.Dish{}
	if err = c.BindJSON(&dish); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	res, err := h.service.UpdateDish(&dish, id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *DishHttpHandler) DeleteDish(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("dish_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err = h.service.DeleteDish(id); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}
