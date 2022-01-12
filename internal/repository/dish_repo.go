package repository

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"my_goods/internal/entities"
	"my_goods/pkg/logger"
)

type DishRepo interface {
	GetAllDishes() *[]entities.Dish
	GetDish(id int) (*entities.Dish, error)
	CreateDish(dish *entities.Dish) *entities.Dish
	UpdateDish(dish *entities.Dish) *entities.Dish
	DeleteDish(id int)
}

type DishRepository struct {
	db  PgxPoolInterface
	ctx context.Context
}

func NewDishRepository(db PgxPoolInterface) *DishRepository {
	return &DishRepository{db: db, ctx: context.Background()}
}

func (r *DishRepository) GetDish(id int) (*entities.Dish, error) {
	var dish entities.Dish
	query := fmt.Sprintf("SELECT id, title, description FROM %s WHERE id=$1", "dishes")
	err := pgxscan.Get(r.ctx, r.db, &dish, query, id)
	if err != nil {
		logger.Error(err)
	}
	return &dish, err
}

func (r *DishRepository) GetAllDishes() *[]entities.Dish {
	var dish []entities.Dish

	return &dish
}

func (r *DishRepository) CreateDish(dish *entities.Dish) *entities.Dish {

	return dish
}

func (r *DishRepository) UpdateDish(dish *entities.Dish) *entities.Dish {

	return dish
}

func (r *DishRepository) DeleteDish(id int) {

}
