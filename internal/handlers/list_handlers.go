package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetList(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}

func (h *Handler) GetAllLists(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}

func (h *Handler) CreateList(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}

func (h *Handler) UpdateList(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}

func (h *Handler) DeleteList(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}
