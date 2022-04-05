package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	BearerSchema   = "Bearer"
	userCtx        = "userId"
	authHeaderName = "Authorization"
)

func AuthenticationMiddleware(c *gin.Context) {
	token, err := getAuthenticationHeader(c)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	manager, err := NewTokenManager()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	userId, err := manager.Parse(token)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.Set(userCtx, userId)
}

func getAuthenticationHeader(c *gin.Context) (string, error) {
	header := c.GetHeader(authHeaderName)
	if header == "" {
		return "", errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != BearerSchema {
		return "", errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New("token is empty")
	}
	return headerParts[1], nil
}
