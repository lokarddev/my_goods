package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetGoods(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}

func (h *Handler) GetAllGoods(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}

func (h *Handler) CreateGoods(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}

func (h *Handler) UpdateGoods(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}

func (h *Handler) DeleteGoods(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}
