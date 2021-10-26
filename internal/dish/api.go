package dish

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DishHandler struct {
	services *DishService
}

func NewDishHandler(services *DishService) *DishHandler {
	return &DishHandler{services: services}
}

func (h *DishHandler) RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		dish := api.Group("/dish")
		{
			dish.GET("/:id", h.GetDish)
			dish.GET("/", h.GetAllDishes)
			dish.POST("/", h.CreateDish)
			dish.PUT("/:id", h.UpdateDish)
			dish.DELETE("/:id", h.DeleteDish)
		}
	}
}

func (h *DishHandler) GetDish(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}

func (h *DishHandler) GetAllDishes(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}

func (h *DishHandler) CreateDish(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}

func (h *DishHandler) UpdateDish(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}

func (h *DishHandler) DeleteDish(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Hello": "world"})
}
