package users_http

import (
	"github.com/gin-gonic/gin"
	"my_goods/internal/delivery"
	"my_goods/internal/entity/dto"
	"net/http"
)

type UsersHttpHandler struct {
	service delivery.UserServiceInterface
}

func NewUsersHttpHandler(service delivery.UserServiceInterface) *UsersHttpHandler {
	return &UsersHttpHandler{service: service}
}

func (h *UsersHttpHandler) RegisterRoutes(root *gin.RouterGroup) {
	root.POST("sign-in/", h.SignIn)
	root.POST("sign-up/", h.SignUp)
	root.POST("refresh-token/", h.RefreshAccessToken)
}

func (h *UsersHttpHandler) SignIn(c *gin.Context) {
	input := dto.LoginRequest{}
	err := c.BindJSON(&input)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	access, err := h.service.SignIn(input)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, access)
}

func (h *UsersHttpHandler) SignUp(c *gin.Context) {
	input := dto.LoginRequest{}
	err := c.BindJSON(&input)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	access, err := h.service.SignUp(input)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, access)
}

// RefreshAccessToken generate new token pair and new session for user
func (h *UsersHttpHandler) RefreshAccessToken(c *gin.Context) {
	input := map[string]string{}
	err := c.BindJSON(&input)
	token, ok := input["refresh"]
	if err != nil || !ok {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	newAccess, err := h.service.RefreshAccess(token)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, newAccess)
}
