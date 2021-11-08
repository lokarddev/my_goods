package dish

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"my_goods/internal/entity"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	service     = mockService{}
	testHandler = NewDishHandler(&service)
)

type mockService struct{}

func (s *mockService) GetDish(id int) *entity.Dish {
	dish := entity.Dish{}
	return &dish
}

func (s *mockService) GetAllDishes() *[]entity.Dish {
	var dishes []entity.Dish
	return &dishes
}

func (s *mockService) CreateDish(dish *entity.Dish) *entity.Dish {
	result := entity.Dish{}
	return &result
}

func (s *mockService) UpdateDish(dish *entity.Dish) *entity.Dish {
	result := entity.Dish{}
	return &result
}

func (s *mockService) DeleteDish(id int) {
	_ = entity.Dish{}
}

func TestAllDishesDish(t *testing.T) {
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	testHandler.GetAllDishes(c)
	assert.Equal(t, http.StatusOK, response.Code)
}
