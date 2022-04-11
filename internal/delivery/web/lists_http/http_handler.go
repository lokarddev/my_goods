package lists_http

import (
	"github.com/gin-gonic/gin"
	"my_goods/internal/delivery"
	"my_goods/internal/entity"
	"my_goods/internal/entity/dto"
	"net/http"
	"strconv"
)

type ListsHttpHandler struct {
	service delivery.ListServiceInterface
}

func NewListsHttpHandler(service delivery.ListServiceInterface) *ListsHttpHandler {
	return &ListsHttpHandler{service: service}
}

func (h *ListsHttpHandler) RegisterRoutes(api *gin.RouterGroup) {
	api.GET("get-lists/:lists_id", h.GetLists)
	api.GET("get-lists/", h.GetAllLists)
	api.POST("create-lists/", h.CreateLists)
	api.POST("update-lists/:lists_id", h.UpdateLists)
	api.DELETE("delete-lists/:lists_id", h.DeleteLists)
	api.POST("add-goods-to-list/", h.AddGoodsToList)
	api.POST("add-dish-to-list/", h.AddDishToList)
	api.GET("get-shopping-list/:lists_id", h.getShoppingList)
}

func (h *ListsHttpHandler) GetLists(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("lists_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	lists, err := h.service.GetList(int32(id))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, lists)
}

func (h *ListsHttpHandler) GetAllLists(c *gin.Context) {
	lists, err := h.service.GetAllLists()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, lists)
}

func (h *ListsHttpHandler) CreateLists(c *gin.Context) {
	lists := entity.List{}
	if err := c.BindJSON(&lists); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	res, err := h.service.CreateList(&lists)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *ListsHttpHandler) UpdateLists(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("lists_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	lists := entity.List{}
	if err = c.BindJSON(&lists); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	res, err := h.service.UpdateList(&lists, int32(id))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *ListsHttpHandler) DeleteLists(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("lists_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err = h.service.DeleteList(int32(id)); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func (h *ListsHttpHandler) AddGoodsToList(c *gin.Context) {
	var request dto.AddGoodsListRequest
	if err := c.BindJSON(&request); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err := h.service.AddGoodsToList(request.ListId, request.Ids); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func (h *ListsHttpHandler) AddDishToList(c *gin.Context) {
	var request dto.AddDishListRequest
	if err := c.BindJSON(&request); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err := h.service.AddDishToLIst(request.ListId, request.Ids); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func (h *ListsHttpHandler) getShoppingList(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("lists_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	list, err := h.service.GetShopping(int32(id))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, list)
}
