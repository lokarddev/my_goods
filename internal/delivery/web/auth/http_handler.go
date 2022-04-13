package auth

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

// SignIn godoc
// @Summary      SignIn
// @Tags         auth
// @Description  sign into account
// @ID           sign in process
// @Accept       json
// @Produce      json
// @Param        input  body      dto.LoginRequest  true  "user name + password"
// @Success      200    {object}  dto.Access
// @Failure      400
// @Failure      500
// @Router       /sign-in [post]
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

// SignUp godoc
// @Summary      SignUp
// @Tags         auth
// @Description  create account
// @ID           create-account
// @Accept       json
// @Produce      json
// @Param        input  body      dto.LoginRequest  true  "user name + password"
// @Success      200    {object}  dto.Access
// @Failure      400
// @Failure      500
// @Router       /sign-up [post]
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
// @Summary      RefreshAccessToken
// @Tags         auth
// @Description  get-new-access
// @ID           get-new-access
// @Accept       json
// @Produce      json
// @Param        input  body      dto.RefreshRequest  true  "refresh token"
// @Success      200    {object}  dto.Access
// @Failure      400
// @Router       /refresh-token [post]
func (h *UsersHttpHandler) RefreshAccessToken(c *gin.Context) {
	input := dto.RefreshRequest{}
	err := c.BindJSON(&input)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	newAccess, err := h.service.RefreshAccess(input.Refresh)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, newAccess)
}
