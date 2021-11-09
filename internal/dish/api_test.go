package dish

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"my_goods/internal/entity"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	service     = mockService{}
	testHandler = NewDishHandler(&service)
)

type mockService struct{}

func (s *mockService) GetDish(id int) *entity.Dish {
	dish := entity.Dish{Model: gorm.Model{ID: 1}}
	if id == int(dish.Model.ID) {
		return &dish
	}
	return &entity.Dish{}
}

func (s *mockService) GetAllDishes() *[]entity.Dish {
	var dishes []entity.Dish
	return &dishes
}

func (s *mockService) CreateDish(dish *entity.Dish) *entity.Dish {
	result := entity.Dish{Title: dish.Title, Description: dish.Description, Model: gorm.Model{ID: 12}}
	return &result
}

func (s *mockService) UpdateDish(dish *entity.Dish) *entity.Dish {
	result := entity.Dish{}
	return &result
}

func (s *mockService) DeleteDish(id int) {
	_ = entity.Dish{}
}

func TestHandler_GetAllDishes(t *testing.T) {
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)

	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	testHandler.GetAllDishes(c)
	assert.NotNil(t, response.Code)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestHandler_GetDish_ValidInput(t *testing.T) {
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)

	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	c.Params = []gin.Param{{Key: "id", Value: "1"}}
	testHandler.GetDish(c)
	var result entity.Dish
	err := json.Unmarshal(response.Body.Bytes(), &result)
	assert.NotNil(t, response.Body)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Nil(t, err)
	assert.Equal(t, int(result.ID), 1)

	response = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(response)
	c.Params = []gin.Param{{Key: "id", Value: "2"}}
	testHandler.GetDish(c)
	assert.NotNil(t, response.Body)
	assert.Equal(t, http.StatusNotFound, response.Code)
}

func TestHandler_GetDish_InvalidInput(t *testing.T) {
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)

	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	c.Params = []gin.Param{{Key: "id", Value: "invalid"}}
	testHandler.GetDish(c)
	assert.NotNil(t, response.Code)
	assert.NotNil(t, response.Body)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	testHandler.GetDish(c)
	assert.NotNil(t, response.Code)
	assert.NotNil(t, response.Body)
	assert.Equal(t, http.StatusBadRequest, response.Code)
}

func TestHandler_CreateDish(t *testing.T) {
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	data := `{   
    "description": "test",
    "title":"test",
    "goods": [
        {"ID": 10,
        "Title": "test",
        "Description": "test"
        }]
	}`
	c.Request, _ = http.NewRequest(http.MethodPost, "", strings.NewReader(data))
	testHandler.CreateDish(c)
	assert.NotNil(t, response.Body)
	var result entity.Dish
	err := json.Unmarshal(response.Body.Bytes(), &result)
	assert.NotNil(t, response.Code)
	assert.Nil(t, err)
	assert.Equal(t, "test", result.Description)
	assert.Equal(t, "test", result.Title)
	assert.NotEqual(t, result.ID, 0)
}

func TestHandler_UpdateDish(t *testing.T) {
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodPut, "", nil)

	testHandler.UpdateDish(c)
}

func TestHandler_DeleteDish(t *testing.T) {
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodDelete, "", nil)

	testHandler.DeleteDish(c)
}
