package goods_repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"my_goods/internal/entity"
	"my_goods/internal/storage/postgres"
	"my_goods/pkg/logger"
)

type GoodsRepository struct {
	postgres.Repository
}

func NewGoodsRepository(db postgres.PgxPoolInterface) *GoodsRepository {
	return &GoodsRepository{Repository: postgres.Repository{
		DB:  db,
		Ctx: context.Background(),
	}}
}

func (r *GoodsRepository) GetGoods(id int32) (*entity.Goods, error) {
	var good entity.PgxGoods
	query := fmt.Sprintf("SELECT id, title, description FROM %s WHERE id=$1", postgres.GoodsTable)
	rows, err := r.DB.Query(r.Ctx, query, id)
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&good.Id, &good.Title, &good.Description); err != nil {
			logger.Error(err)
		}
	}
	if err != nil {
		logger.Error(err)
	}
	return good.ToClean(), err
}

func (r *GoodsRepository) GetAllGoods() (*[]entity.Goods, error) {
	var goods []entity.Goods
	query := fmt.Sprintf("SELECT id, title, description FROM %s", postgres.GoodsTable)
	rows, err := r.DB.Query(r.Ctx, query)
	defer rows.Close()
	for rows.Next() {
		var good entity.PgxGoods
		if err = rows.Scan(&good.Id, &good.Title, &good.Description); err != nil {
			logger.Error(err)
		}
		goods = append(goods, *good.ToClean())
	}
	if err != nil {
		logger.Error(err)
	}
	return &goods, err
}

func (r *GoodsRepository) CreateGoods(good *entity.Goods) (*entity.Goods, error) {
	var pgxGoods entity.PgxGoods
	query := fmt.Sprintf("INSERT INTO %s (created_at, updated_at, title, description) VALUES (now(), now(), $1, $2) RETURNING id, title, description", postgres.GoodsTable)
	err := r.DB.BeginTxFunc(r.Ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		rows, err := r.DB.Query(r.Ctx, query, good.Title, good.Description)
		defer rows.Close()
		for rows.Next() {
			if err = rows.Scan(&pgxGoods.Id, &pgxGoods.Title, &pgxGoods.Description); err != nil {
				return err
			}
		}
		return err
	})
	if err != nil {
		logger.Error(err)
	}
	return pgxGoods.ToClean(), err
}

func (r *GoodsRepository) UpdateGoods(good *entity.Goods, id int32) (*entity.Goods, error) {
	var pgxGoods entity.PgxGoods
	query := fmt.Sprintf("UPDATE %s SET updated_at=now(), title=$1, description=$2 WHERE id=$3 RETURNING id, title, description", postgres.GoodsTable)
	err := r.DB.BeginTxFunc(r.Ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		rows, err := r.DB.Query(r.Ctx, query, good.Title, good.Description, id)
		defer rows.Close()
		for rows.Next() {
			if err = rows.Scan(&pgxGoods.Id, &pgxGoods.Title, &pgxGoods.Description); err != nil {
				return err
			}
		}
		return err
	})
	if err != nil {
		logger.Error(err)
	}
	return pgxGoods.ToClean(), err
}

func (r *GoodsRepository) DeleteGoods(id int32) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", postgres.GoodsTable)
	err := r.DB.BeginTxFunc(r.Ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		_, err := r.DB.Exec(r.Ctx, query, id)
		return err
	})
	if err != nil {
		logger.Error(err)
	}
	return err
}
