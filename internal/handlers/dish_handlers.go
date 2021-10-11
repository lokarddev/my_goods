package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetDish(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}

func (h *Handler) GetAllDishes(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}

func (h *Handler) CreateDish(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}

func (h *Handler) UpdateDish(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}

func (h *Handler) DeleteDish(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}
