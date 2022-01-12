package repository

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"my_goods/internal/entities"
	"regexp"
	"testing"
)

func TestDishSuite(t *testing.T) {
	suite.Run(t, new(testDishSuite))
}

type testDishSuite struct {
	suite.Suite
	db   *gorm.DB
	mock sqlmock.Sqlmock
	repo DishRepo
}

func (ts *testDishSuite) SetupMockDB() {
	var (
		db  *sql.DB
		err error
	)
	db, ts.mock, err = sqlmock.New()
	ts.db, err = gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	require.NoError(ts.T(), err)
	ts.repo = NewDishRepository(ts.db)
}

func (ts *testDishSuite) TestGetDish() {
	ts.SetupMockDB()
	ts.T().Run("Valid", func(t *testing.T) {
		ts.mock.ExpectQuery(regexp.QuoteMeta(
			"select * from "))

		result := ts.repo.GetDish(1)
		assert.IsType(ts.T(), entities.Dish{}, result)
	})
}

func (ts *testDishSuite) TestGetAllDishes() {

}

func (ts *testDishSuite) TestCreateDish() {

}

func (ts *testDishSuite) TestUpdateDish() {

}

func (ts *testDishSuite) TestDeleteDish() {

}
