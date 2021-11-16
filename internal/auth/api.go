package auth

import (
	"github.com/gin-gonic/gin"
	"my_goods/pkg/logger"
	"net/http"
	"strings"
)

type Handler struct {
	services ServeAuth
}

type Auth struct {
	Login string `json:"login"`
	Pass  string `json:"pass"`
}

func NewAuthHandler(services ServeAuth) *Handler {
	return &Handler{services: services}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.SignIn)
		auth.POST("/sign-up", h.SignUp)
	}
	_ = router.Group("/api", h.AuthMiddleware)
}

func (h *Handler) SignIn(c *gin.Context) {
	var input Auth
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"ERROR": "invalid input"})
		return
	}
	token, err := h.services.GenerateToken(input)
	if err != nil {
		logger.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"ERROR": "something goes wrong try again"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *Handler) SignUp(c *gin.Context) {
	var input Auth
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"ERROR": "invalid input"})
		return
	}
	id, err := h.services.CreateUser(input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"ERROR": "something goes wrong try again"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *Handler) AuthMiddleware(c *gin.Context) {
	header := strings.Split(c.GetHeader("Authorization"), " ")
	if len(header) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"ERROR": "invalid authentication"})
		return
	}
	user, err := h.services.ParseToken(header[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"ERROR": "something goes wrong try again"})
		return
	}
	c.Set("user", user)
}
