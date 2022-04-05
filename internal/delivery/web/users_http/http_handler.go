package users_http

import (
	"github.com/gin-gonic/gin"
	"my_goods/internal/delivery"
)

type UsersHttpHandler struct {
	service delivery.DishServiceInterface
}

func NewUsersHttpHandler(service delivery.DishServiceInterface) *UsersHttpHandler {
	return &UsersHttpHandler{service: service}
}

func (h *UsersHttpHandler) RegisterRoutes(root *gin.RouterGroup) {
	root.POST("sign-in/", h.SignIn)
	root.POST("sign-up/", h.SignUp)
	root.POST("refresh-token/", h.RefreshAccessToken)
}

func (h *UsersHttpHandler) SignIn(c *gin.Context) {

}

func (h *UsersHttpHandler) SignUp(c *gin.Context) {

}

// RefreshAccessToken generate new token pair and new session for user
func (h *UsersHttpHandler) RefreshAccessToken(c *gin.Context) {

}
