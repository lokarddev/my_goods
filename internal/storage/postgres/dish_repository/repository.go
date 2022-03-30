package dish_repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"my_goods/internal/entity"
	"my_goods/internal/storage/postgres"
	"my_goods/pkg/logger"
	"sync"
)

type DishRepository struct {
	db  postgres.PgxPoolInterface
	ctx context.Context
}

func NewDishRepository(db postgres.PgxPoolInterface) *DishRepository {
	return &DishRepository{db: db, ctx: context.Background()}
}

func (r *DishRepository) GetDish(id int32) (*entity.DishesResponse, error) {
	var err error
	dish := &entity.Dish{}
	goods := &[]entity.GoodsWithAmount{}
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func(wg *sync.WaitGroup, id int32) {
		dish, err = r.getDishInfo(id)
		wg.Done()
	}(wg, id)

	go func(wg *sync.WaitGroup, id int32) {
		goods, err = r.getGoodsInfo(id)
		wg.Done()
	}(wg, id)

	wg.Wait()
	return &entity.DishesResponse{Dish: *dish, Goods: *goods}, err
}

func (r *DishRepository) getGoodsInfo(id int32) (*[]entity.GoodsWithAmount, error) {
	var goods []entity.GoodsWithAmount
	query := fmt.Sprintf("SELECT goods.Id, goods.title, goods.description, dish_goods.amount FROM %s "+
		"FULL JOIN %s ON dish_goods.goods_id = goods.id WHERE dish_id=$1", postgres.DishGoodsTable, postgres.GoodsTable)
	rows, err := r.db.Query(r.ctx, query, id)
	defer rows.Close()
	for rows.Next() {
		goodInfo := entity.GoodsWithAmount{}
		if err = rows.Scan(&goodInfo.Id, &goodInfo.Title, &goodInfo.Description, &goodInfo.Amount); err != nil {
			logger.Error(err)
		}
		goods = append(goods, goodInfo)
	}
	return &goods, err
}

func (r *DishRepository) getDishInfo(id int32) (*entity.Dish, error) {
	var pgxDish entity.PgxDish
	query := fmt.Sprintf("SELECT id, title, description FROM %s WHERE id=$1", postgres.DishTable)
	rows, err := r.db.Query(r.ctx, query, id)
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&pgxDish.Id, &pgxDish.Title, &pgxDish.Description); err != nil {
			logger.Error(err)
		}
	}
	if err != nil {
		logger.Error(err)
	}
	return pgxDish.ToClean(), err
}

func (r *DishRepository) GetAllDishes() (*[]entity.DishesResponse, error) {
	var dishes []entity.Dish
	var response []entity.DishesResponse
	query := fmt.Sprintf("SELECT id, title, description FROM %s", postgres.DishTable)
	rows, err := r.db.Query(r.ctx, query)
	defer rows.Close()
	for rows.Next() {
		var dish entity.PgxDish
		if err = rows.Scan(&dish.Id, &dish.Title, &dish.Description); err != nil {
			logger.Error(err)
		}
		dishes = append(dishes, *dish.ToClean())
	}
	if err != nil {
		logger.Error(err)
	}
	for _, dish := range dishes {
		goods, err := r.getGoodsInfo(dish.Id)
		if err != nil {
			logger.Error(err)
			return &response, err
		}
		response = append(response, entity.DishesResponse{
			Dish:  dish,
			Goods: *goods,
		})
	}
	return &response, err
}

func (r *DishRepository) CreateDish(dish *entity.Dish) (*entity.Dish, error) {
	var pgxDish entity.PgxDish
	query := fmt.Sprintf("INSERT INTO %s (created_at, updated_at, title, description) VALUES (now(), now(), $1, $2) RETURNING id, title, description", postgres.DishTable)
	err := r.db.BeginTxFunc(r.ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		rows, err := r.db.Query(r.ctx, query, dish.Title, dish.Description)
		defer rows.Close()
		for rows.Next() {
			if err = rows.Scan(&pgxDish.Id, &pgxDish.Title, &pgxDish.Description); err != nil {
				return err
			}
		}
		return err
	})
	if err != nil {
		logger.Error(err)
	}
	return pgxDish.ToClean(), err
}

func (r *DishRepository) UpdateDish(dish *entity.Dish, id int32) (*entity.DishesResponse, error) {
	var pgxDish entity.PgxDish
	query := fmt.Sprintf("UPDATE %s SET updated_at=now(), title=$1, description=$2 WHERE id=$3 RETURNING id, title, description", postgres.DishTable)
	err := r.db.BeginTxFunc(r.ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		rows, err := r.db.Query(r.ctx, query, dish.Title, dish.Description, id)
		defer rows.Close()
		for rows.Next() {
			if err = rows.Scan(&pgxDish.Id, &pgxDish.Title, &pgxDish.Description); err != nil {
				return err
			}
		}
		return err
	})
	if err != nil {
		logger.Error(err)
	}
	goods, err := r.getGoodsInfo(id)
	return &entity.DishesResponse{
		Dish:  *pgxDish.ToClean(),
		Goods: *goods,
	}, err
}

func (r *DishRepository) DeleteDish(id int32) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", postgres.DishTable)
	err := r.db.BeginTxFunc(r.ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		_, err := r.db.Exec(r.ctx, query, id)
		return err
	})
	if err != nil {
		logger.Error(err)
	}
	return err
}

func (r *DishRepository) AddGoodsToDish(dishId int32, goods map[int32]int32) error {
	b := &pgx.Batch{}
	for k, v := range goods {
		query := fmt.Sprintf("INSERT INTO %s (dish_id, goods_id, amount) VALUES ($1, $2, $3)", postgres.DishGoodsTable)
		b.Queue(query, dishId, k, v)
	}
	err := r.db.BeginTxFunc(r.ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		res := tx.SendBatch(r.ctx, b)
		err := res.Close()
		if err != nil {
			logger.Error(err)
		}
		return err
	})
	return err
}

func (r *DishRepository) RemoveGoodsFromDish(dishId int32, goodsIds []int32) error {
	b := &pgx.Batch{}
	for _, v := range goodsIds {
		query := fmt.Sprintf("DELETE FROM %s WHERE dish_id=$1 AND goods_id=$2", postgres.DishGoodsTable)
		b.Queue(query, dishId, v)
	}
	err := r.db.BeginTxFunc(r.ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		res := tx.SendBatch(r.ctx, b)
		err := res.Close()
		if err != nil {
			logger.Error(err)
		}
		return err
	})
	return err
}
