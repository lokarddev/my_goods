package dish

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"my_goods/pkg/db"
	"my_goods/pkg/environ"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	dish = "/api/dish"
)

var (
	testDb      *gorm.DB
	testHandler *gin.Engine
	TS          *httptest.Server
)

func init() {
	environ.Env()
	testDb, _ = db.DB(db.NewDatabaseConf())
	testHandler = gin.New()

	dishRepo := NewDishRepo(testDb)
	dishService := NewDishService(*dishRepo)
	TestDishHandler := NewDishHandler(dishService)
	TestDishHandler.RegisterRoutes(testHandler)
}

func TestGetDish(t *testing.T) {
	TS = httptest.NewServer(testHandler)
	defer TS.Close()

	resp, err := http.Get(fmt.Sprintf("%s%s", TS.URL, dish))
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
