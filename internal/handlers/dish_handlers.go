package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Test(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}
