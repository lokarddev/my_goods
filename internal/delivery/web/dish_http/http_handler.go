package dish_http

import (
	"github.com/gin-gonic/gin"
	"my_goods/internal/delivery"
	"my_goods/internal/entity"
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
	api.GET("get_dish/:dish_id", h.GetDish)
	api.GET("get_dishes/", h.GetAllDishes)
	api.POST("create_dish/", h.CreateDish)
	api.POST("update_dish/:dish_id", h.UpdateDish)
	api.POST("add_goods_to_dish/", h.AddGoodsToDish)
	api.DELETE("delete_dish/:dish_id", h.DeleteDish)
	api.DELETE("remove_goods_from_dish/", h.RemoveGoodsFromDish)
}

func (h *DishHttpHandler) GetDish(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("dish_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	dish, err := h.service.GetDish(int32(id))
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
	dish := entity.Dish{}
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
	dish := entity.Dish{}
	if err = c.BindJSON(&dish); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	res, err := h.service.UpdateDish(&dish, int32(id))
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
	if err = h.service.DeleteDish(int32(id)); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func (h *DishHttpHandler) AddGoodsToDish(c *gin.Context) {
	var goods entity.AddToDishRequest
	if err := c.BindJSON(&goods); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err := h.service.AddGoods(goods.DishId, goods.Ids)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func (h *DishHttpHandler) RemoveGoodsFromDish(c *gin.Context) {
	var goods entity.RemoveFromDishRequest
	if err := c.BindJSON(&goods); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err := h.service.RemoveGoodsFromDish(goods.DishId, goods.Ids)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}
