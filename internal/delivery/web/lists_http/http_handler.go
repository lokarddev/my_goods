package lists_http

import (
	"github.com/gin-gonic/gin"
	"my_goods/internal/delivery"
	"my_goods/internal/delivery/web/auth"
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

// GetLists godoc
// @Security     ApiKeyAuth
// @Summary      GetLists
// @Tags         lists
// @Description  getting specific lists from user
// @ID           get-lists
// @Accept       json
// @Produce      json
// @Param        lists_id  path      integer  true  "lists id"
// @Success      200       {object}  dto.ListsResponse
// @Failure      400
// @Failure      500
// @Router       /api/get-lists/{lists_id} [get]
func (h *ListsHttpHandler) GetLists(c *gin.Context) {
	userId := auth.GetUserId(c)
	id, err := strconv.Atoi(c.Param("lists_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	lists, err := h.service.GetList(int32(id), userId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, lists)
}

// GetAllLists godoc
// @Security     ApiKeyAuth
// @Summary      GetAllLists
// @Tags         lists
// @Description  getting all lists by user
// @ID           get-all-lists
// @Accept       json
// @Produce      json
// @Success      200  {object}  []dto.ListsResponse
// @Failure      400
// @Failure      500
// @Router       /api/get-lists/ [get]
func (h *ListsHttpHandler) GetAllLists(c *gin.Context) {
	userId := auth.GetUserId(c)
	lists, err := h.service.GetAllLists(userId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, lists)
}

// CreateLists godoc
// @Security     ApiKeyAuth
// @Summary      CreateLists
// @Tags         lists
// @Description  creating new lists from input
// @ID           create-lists
// @Accept       json
// @Produce      json
// @Param        input  body      entity.List  true  "entry data for lists"
// @Success      200    {object}  entity.List
// @Failure      400
// @Failure      500
// @Router       /api/create-lists [post]
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

// UpdateLists godoc
// @Security     ApiKeyAuth
// @Summary      UpdateLists
// @Tags         lists
// @Description  updating existing lists
// @ID           update-lists
// @Accept       json
// @Produce      json
// @Param        lists_id  path      integer      true  "lists id"
// @Param        input     body      entity.List  true  "updating data for lists"
// @Success      200       {object}  dto.ListsResponse
// @Failure      400
// @Failure      500
// @Router       /api/update-lists/{lists_id} [post]
func (h *ListsHttpHandler) UpdateLists(c *gin.Context) {
	userId := auth.GetUserId(c)
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
	res, err := h.service.UpdateList(&lists, int32(id), userId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, res)
}

// DeleteLists godoc
// @Security     ApiKeyAuth
// @Summary      DeleteLists
// @Tags         lists
// @Description  deleting existing lists
// @ID           delete-lists
// @Accept       json
// @Produce      json
// @Param        lists_id  path  integer  true  "lists id"
// @Success      200
// @Failure      400
// @Failure      500
// @Router       /api/delete-lists/{lists_id} [delete]
func (h *ListsHttpHandler) DeleteLists(c *gin.Context) {
	userId := auth.GetUserId(c)
	id, err := strconv.Atoi(c.Param("lists_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err = h.service.DeleteList(int32(id), userId); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

// AddGoodsToList godoc
// @Security     ApiKeyAuth
// @Summary      AddGoodsToList
// @Tags         lists
// @Description  goods new dish to list
// @ID           add-goods-to-list
// @Accept       json
// @Produce      json
// @Param        input  body  dto.AddGoodsListRequest  true  "goods that should be added to list"
// @Success      200
// @Failure      400
// @Failure      500
// @Router       /api/add-goods-to-list/ [post]
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

// AddDishToList godoc
// @Security     ApiKeyAuth
// @Summary      AddDishToList
// @Tags         lists
// @Description  adding new dish to list
// @ID           add-dish-to-list
// @Accept       json
// @Produce      json
// @Param        input  body  dto.AddDishListRequest  true  "dishes that should be added to list"
// @Success      200
// @Failure      400
// @Failure      500
// @Router       /api/add-dish-to-list/ [post]
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

// getShoppingList godoc
// @Security     ApiKeyAuth
// @Summary      getShoppingList
// @Tags         lists
// @Description  getting full list for shopping
// @ID           get-shopping-list
// @Accept       json
// @Produce      json
// @Param        list_id  path       integer  true  "list id"
// @Success      200      {integer}  integer  1
// @Failure      400
// @Failure      500
// @Router       /api/get-shopping-list/{list_id}/ [post]
func (h *ListsHttpHandler) getShoppingList(c *gin.Context) {
	userId := auth.GetUserId(c)
	id, err := strconv.Atoi(c.Param("lists_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	list, err := h.service.GetShopping(int32(id), userId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, list)
}
