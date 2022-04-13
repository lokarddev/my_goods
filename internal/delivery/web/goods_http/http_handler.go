package goods_http

import (
	"github.com/gin-gonic/gin"
	"my_goods/internal/delivery"
	"my_goods/internal/delivery/web/auth"
	"my_goods/internal/entity"
	"net/http"
	"strconv"
)

type GoodsHttpHandler struct {
	service delivery.GoodsServiceInterface
}

func NewGoodsHttpHandler(service delivery.GoodsServiceInterface) *GoodsHttpHandler {
	return &GoodsHttpHandler{service: service}
}

func (h *GoodsHttpHandler) RegisterRoutes(api *gin.RouterGroup) {
	api.GET("get-goods/:goods_id", h.GetGoods)
	api.GET("get-goods/", h.GetAllGoods)
	api.POST("create-goods/", h.CreateGoods)
	api.POST("update-goods/:goods_id", h.UpdateGoods)
	api.DELETE("delete-goods/:goods_id", h.DeleteGoods)
}

// GetGoods godoc
// @Security     ApiKeyAuth
// @Summary      GetGoods
// @Tags         goods
// @Description  getting specific goods from user
// @ID           get-goods
// @Accept       json
// @Produce      json
// @Param        goods_id  path      integer  true  "goods id"
// @Success      200       {object}  entity.Goods
// @Failure      400
// @Failure      500
// @Router       /api/get-goods/{goods_id} [get]
func (h *GoodsHttpHandler) GetGoods(c *gin.Context) {
	userId := auth.GetUserId(c)
	id, err := strconv.Atoi(c.Param("goods_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	goods, err := h.service.GetGoods(int32(id), userId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, goods)
}

// GetAllGoods godoc
// @Security     ApiKeyAuth
// @Summary      GetAllGoods
// @Tags         goods
// @Description  getting all goods by user
// @ID           get-all-goods
// @Accept       json
// @Produce      json
// @Success      200  {object}  []entity.Goods
// @Failure      400
// @Failure      500
// @Router       /api/get-goods/ [get]
func (h *GoodsHttpHandler) GetAllGoods(c *gin.Context) {
	userId := auth.GetUserId(c)
	dishes, err := h.service.GetAllGoods(userId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, dishes)
}

// CreateGoods godoc
// @Security     ApiKeyAuth
// @Summary      CreateGoods
// @Tags         goods
// @Description  creating new goods from input
// @ID           create-goods
// @Accept       json
// @Produce      json
// @Param        input  body      entity.Goods  true  "entry data for goods"
// @Success      200       {object}  entity.Goods
// @Failure      400
// @Failure      500
// @Router       /api/create-goods [post]
func (h *GoodsHttpHandler) CreateGoods(c *gin.Context) {
	goods := entity.Goods{}
	if err := c.BindJSON(&goods); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	res, err := h.service.CreateGoods(&goods)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, res)
}

// UpdateGoods godoc
// @Security     ApiKeyAuth
// @Summary      UpdateGoods
// @Tags         goods
// @Description  updating existing goods
// @ID           update-goods
// @Accept       json
// @Produce      json
// @Param        goods_id  path      integer       true  "goods id"
// @Param        input     body      entity.Goods  true  "updating data for goods"
// @Success      200    {object}  entity.Goods
// @Failure      400
// @Failure      500
// @Router       /api/update-goods/{goods_id} [post]
func (h *GoodsHttpHandler) UpdateGoods(c *gin.Context) {
	userId := auth.GetUserId(c)
	id, err := strconv.Atoi(c.Param("goods_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	goods := entity.Goods{}
	if err = c.BindJSON(&goods); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	res, err := h.service.UpdateGoods(&goods, int32(id), userId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, res)
}

// DeleteGoods godoc
// @Security     ApiKeyAuth
// @Summary      DeleteGoods
// @Tags         goods
// @Description  deleting existing goods
// @ID           delete-goods
// @Accept       json
// @Produce      json
// @Param        goods_id  path  integer  true  "goods id"
// @Success      200
// @Failure      400
// @Failure      500
// @Router       /api/delete-goods/{goods_id} [delete]
func (h *GoodsHttpHandler) DeleteGoods(c *gin.Context) {
	userId := auth.GetUserId(c)
	id, err := strconv.Atoi(c.Param("goods_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err = h.service.DeleteGoods(int32(id), userId); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}
