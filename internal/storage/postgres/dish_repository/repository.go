package dish_repository

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"my_goods/internal/entities"
	"my_goods/internal/storage/postgres"
	"my_goods/pkg/logger"
)

type DishRepository struct {
	db  postgres.PgxPoolInterface
	ctx context.Context
}

func NewDishRepository(db postgres.PgxPoolInterface) *DishRepository {
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

func (r *DishRepository) GetAllDishes() (*[]entities.Dish, error) {
	var dish []entities.Dish
	return &dish, nil
}

func (r *DishRepository) CreateDish(dish *entities.Dish) (*entities.Dish, error) {
	return dish, nil
}

func (r *DishRepository) UpdateDish(dish *entities.Dish, id int) (*entities.Dish, error) {
	return dish, nil
}

func (r *DishRepository) DeleteDish(id int) error {
	return nil
}
