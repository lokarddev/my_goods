package dish_http

import (
	"github.com/gin-gonic/gin"
	"my_goods/internal/delivery"
	"my_goods/internal/delivery/web/auth"
	"my_goods/internal/entity"
	"my_goods/internal/entity/dto"
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
	api.POST("add-goods-to-dish/", h.AddGoodsToDish)
	api.DELETE("delete-dish/:dish_id", h.DeleteDish)
	api.DELETE("remove-goods-from-dish/", h.RemoveGoodsFromDish)
}

// GetDish godoc
// @Security     ApiKeyAuth
// @Summary      GetDish
// @Tags         dishes
// @Description  getting specific dish dish from user
// @ID           get-dish
// @Accept       json
// @Produce      json
// @Param        dish_id  path      integer  true  "dish id"
// @Success      200      {object}  dto.DishesResponse
// @Failure      400
// @Failure      500
// @Router       /api/get-dish/{dish_id} [get]
func (h *DishHttpHandler) GetDish(c *gin.Context) {
	userId := auth.GetUserId(c)
	dishId, err := strconv.Atoi(c.Param("dish_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	dish, err := h.service.GetDish(int32(dishId), userId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, dish)
}

// GetAllDishes godoc
// @Security     ApiKeyAuth
// @Summary      GetAllDishes
// @Tags         dishes
// @Description  getting all dishes by user
// @ID           get-all-dishes
// @Accept       json
// @Produce      json
// @Success      200  {object}  []dto.DishesResponse
// @Failure      400
// @Failure      500
// @Router       /api/get-dishes/ [get]
func (h *DishHttpHandler) GetAllDishes(c *gin.Context) {
	userId := auth.GetUserId(c)
	dishes, err := h.service.GetAllDishes(userId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, dishes)
}

// CreateDish godoc
// @Security     ApiKeyAuth
// @Summary      CreateDish
// @Tags         dishes
// @Description  creating new dish from input
// @ID           create-dish
// @Accept       json
// @Produce      json
// @Param        input  body      entity.Dish  true  "entry data for dish"
// @Success      200    {object}  entity.Dish
// @Failure      400
// @Failure      500
// @Router       /api/create-dish [post]
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

// UpdateDish godoc
// @Security     ApiKeyAuth
// @Summary      UpdateDish
// @Tags         dishes
// @Description  updating existing dish
// @ID           update-dish
// @Accept       json
// @Produce      json
// @Param        dish_id  path      integer      true  "dish id"
// @Param        input    body      entity.Dish  true  "updating data for dish"
// @Success      200      {object}  dto.DishesResponse
// @Failure      400
// @Failure      500
// @Router       /api/update-dish/{dish_id} [post]
func (h *DishHttpHandler) UpdateDish(c *gin.Context) {
	userId := auth.GetUserId(c)
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
	res, err := h.service.UpdateDish(&dish, int32(id), userId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, res)
}

// DeleteDish godoc
// @Security     ApiKeyAuth
// @Summary      DeleteDish
// @Tags         dishes
// @Description  deleting existing dish
// @ID           delete-dish
// @Accept       json
// @Produce      json
// @Param        dish_id  path  integer  true  "dish id"
// @Success      200
// @Failure      400
// @Failure      500
// @Router       /api/delete-dish/{dish_id} [delete]
func (h *DishHttpHandler) DeleteDish(c *gin.Context) {
	userId := auth.GetUserId(c)
	id, err := strconv.Atoi(c.Param("dish_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err = h.service.DeleteDish(int32(id), userId); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

// AddGoodsToDish godoc
// @Security     ApiKeyAuth
// @Summary      AddGoodsToDish
// @Tags         dishes
// @Description  append existing good to existing dish
// @ID           add-goods-to-dish
// @Accept       json
// @Produce      json
// @Param        input  body  dto.AddToDishRequest  true  "goods that should be added to dish"
// @Success      200
// @Failure      400
// @Failure      500
// @Router       /api/add-goods-to-dish/ [post]
func (h *DishHttpHandler) AddGoodsToDish(c *gin.Context) {
	var goods dto.AddToDishRequest
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

// RemoveGoodsFromDish godoc
// @Security     ApiKeyAuth
// @Summary      RemoveGoodsFromDish
// @Tags         dishes
// @Description  removing goods from existing dish
// @ID           remove-goods-from-dish
// @Accept       json
// @Produce      json
// @Param        input  body  dto.RemoveFromDishRequest  true  "goods that should be removed to dish"
// @Success      200
// @Failure      400
// @Failure      500
// @Router       /api/remove-goods-from-dish/ [post]
func (h *DishHttpHandler) RemoveGoodsFromDish(c *gin.Context) {
	var goods dto.RemoveFromDishRequest
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
