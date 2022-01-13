package repository

import (
	"context"
	"github.com/driftprogramming/pgxpoolmock"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgtype"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDish(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
	id := pgtype.Int4{Int: int32(1), Status: pgtype.Present}
	title := pgtype.Varchar{String: "test_title", Status: pgtype.Present}
	description := pgtype.Varchar{String: "test_description", Status: pgtype.Present}

	columns := []string{"id", "title", "description"}
	pgxRows := pgxpoolmock.NewRows(columns).AddRow(id, title, description).ToPgxRows()
	mockPool.EXPECT().Query(context.Background(), "SELECT id, title, description FROM dishes WHERE id=$1", 1).Return(pgxRows, nil)
	db := NewDishRepository(mockPool)

	result, err := db.GetDish(1)

	assert.Equal(t, nil, err)
	assert.NotNil(t, result)
	assert.Equal(t, title, result.Title)
	assert.Equal(t, description, result.Description)
}

func TestGetAllDishes(t *testing.T) {

}

func TestCreateDish(t *testing.T) {

}

func TestUpdateDish(t *testing.T) {

}

func TestDeleteDish(t *testing.T) {

}
