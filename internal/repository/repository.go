package repository

import (
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

// PgxPoolInterface minimal interface for mocking and testing
type PgxPoolInterface interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, optionsAndArgs ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, optionsAndArgs ...interface{}) pgx.Row
}

type Repository struct {
	Goods GoodsRepo
	Dish  DishRepo
	List  ListRepo
}

func NewRepository(db PgxPoolInterface) *Repository {
	return &Repository{
		Goods: NewGoodsRepository(db),
		Dish:  NewDishRepository(db),
		List:  NewListRepository(db),
	}
}
