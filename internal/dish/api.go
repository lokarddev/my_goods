package dish

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	services *Service
}

func NewDishHandler(services *Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		dish := api.Group("/dish")
		{
			dish.GET("/:id", h.getDish)
			dish.GET("/", h.getAllDishes)
			dish.POST("/", h.createDish)
			dish.PUT("/:id", h.updateDish)
			dish.DELETE("/:id", h.deleteDish)
		}
	}
}

func (h *Handler) getDish(c *gin.Context) {

	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}

func (h *Handler) getAllDishes(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}

func (h *Handler) createDish(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}

func (h *Handler) updateDish(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}

func (h *Handler) deleteDish(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}
