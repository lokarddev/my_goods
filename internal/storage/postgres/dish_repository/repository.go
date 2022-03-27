package dish_repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"my_goods/internal/entity"
	"my_goods/internal/storage/postgres"
	"my_goods/pkg/logger"
)

const (
	dishTable = "dishes"
)

type DishRepository struct {
	db  postgres.PgxPoolInterface
	ctx context.Context
}

func NewDishRepository(db postgres.PgxPoolInterface) *DishRepository {
	return &DishRepository{db: db, ctx: context.Background()}
}

func (r *DishRepository) GetDish(id int) (*entity.Dish, error) {
	var dish entity.Dish
	query := fmt.Sprintf("SELECT id, title, description FROM %s WHERE id=$1", dishTable)
	rows, err := r.db.Query(r.ctx, query, id)
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&dish); err != nil {
			logger.Error(err)
		}
	}
	if err != nil {
		logger.Error(err)
	}
	return &dish, err
}

func (r *DishRepository) GetAllDishes() (*[]entity.Dish, error) {
	var dishes []entity.Dish
	query := fmt.Sprintf("SELECT * FROM %s", dishTable)
	rows, err := r.db.Query(r.ctx, query)
	defer rows.Close()
	for rows.Next() {
		var dish entity.Dish
		if err = rows.Scan(&dish); err != nil {
			logger.Error(err)
		}
		dishes = append(dishes, dish)
	}
	if err != nil {
		logger.Error(err)
	}
	return &dishes, err
}

func (r *DishRepository) CreateDish(dish *entity.Dish) (*entity.Dish, error) {
	query := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING *", dishTable)
	err := r.db.BeginTxFunc(r.ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		rows, err := r.db.Query(r.ctx, query, dish.Title, dish.Description)
		defer rows.Close()
		for rows.Next() {
			if err = rows.Scan(dish); err != nil {
				return err
			}
		}
		return err
	})
	if err != nil {
		logger.Error(err)
	}
	return dish, err
}

func (r *DishRepository) UpdateDish(dish *entity.Dish, id int) (*entity.Dish, error) {
	query := fmt.Sprintf("UPDATE %s SET title=$1, description=$2 WHERE id=$3 RETURNING *", dishTable)
	err := r.db.BeginTxFunc(r.ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		rows, err := r.db.Query(r.ctx, query, dish.Title, dish.Description, id)
		defer rows.Close()
		for rows.Next() {
			if err = rows.Scan(dish); err != nil {
				return err
			}
		}
		return err
	})
	if err != nil {
		logger.Error(err)
	}
	return dish, err
}

func (r *DishRepository) DeleteDish(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", dishTable)
	err := r.db.BeginTxFunc(r.ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		_, err := r.db.Exec(r.ctx, query, id)
		return err
	})
	if err != nil {
		logger.Error(err)
	}
	return err
}
