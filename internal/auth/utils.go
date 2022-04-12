package auth

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetUserId(c *gin.Context) int32 {
	userId, ok := c.Get("user_id")
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return 0
	}
	return userId.(int32)
}
