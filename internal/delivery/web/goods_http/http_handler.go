package goods_http

import (
	"github.com/gin-gonic/gin"
	"my_goods/internal/delivery"
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

func (h *GoodsHttpHandler) GetGoods(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("goods_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	goods, err := h.service.GetGoods(id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, goods)
}

func (h *GoodsHttpHandler) GetAllGoods(c *gin.Context) {
	dishes, err := h.service.GetAllGoods()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, dishes)
}

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

func (h *GoodsHttpHandler) UpdateGoods(c *gin.Context) {
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
	res, err := h.service.UpdateGoods(&goods, id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *GoodsHttpHandler) DeleteGoods(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("goods_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err = h.service.DeleteGoods(id); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}
