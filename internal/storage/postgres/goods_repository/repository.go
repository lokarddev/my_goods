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
	db  postgres.PgxPoolInterface
	ctx context.Context
}

func NewGoodsRepository(db postgres.PgxPoolInterface) *GoodsRepository {
	return &GoodsRepository{db: db, ctx: context.Background()}
}

const (
	goodsTable = "goods"
)

func (r *GoodsRepository) GetGoods(id int) (*entity.Goods, error) {
	var good entity.Goods
	query := fmt.Sprintf("SELECT id, title, description FROM %s WHERE id=$1", goodsTable)
	rows, err := r.db.Query(r.ctx, query, id)
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&good); err != nil {
			logger.Error(err)
		}
	}
	if err != nil {
		logger.Error(err)
	}
	return &good, err
}

func (r *GoodsRepository) GetAllGoods() (*[]entity.Goods, error) {
	var goods []entity.Goods
	query := fmt.Sprintf("SELECT * FROM %s", goodsTable)
	rows, err := r.db.Query(r.ctx, query)
	defer rows.Close()
	for rows.Next() {
		var good entity.Goods
		if err = rows.Scan(&good); err != nil {
			logger.Error(err)
		}
		goods = append(goods, good)
	}
	if err != nil {
		logger.Error(err)
	}
	return &goods, err
}

func (r *GoodsRepository) CreateGoods(good *entity.Goods) (*entity.Goods, error) {
	query := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING *", goodsTable)
	err := r.db.BeginTxFunc(r.ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		rows, err := r.db.Query(r.ctx, query, good.Title, good.Description)
		defer rows.Close()
		for rows.Next() {
			if err = rows.Scan(good); err != nil {
				return err
			}
		}
		return err
	})
	if err != nil {
		logger.Error(err)
	}
	return good, err
}

func (r *GoodsRepository) UpdateGoods(good *entity.Goods, id int) (*entity.Goods, error) {
	query := fmt.Sprintf("UPDATE %s SET title=$1, description=$2 WHERE id=$3 RETURNING *", goodsTable)
	err := r.db.BeginTxFunc(r.ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		rows, err := r.db.Query(r.ctx, query, good.Title, good.Description, id)
		defer rows.Close()
		for rows.Next() {
			if err = rows.Scan(good); err != nil {
				return err
			}
		}
		return err
	})
	if err != nil {
		logger.Error(err)
	}
	return good, err
}

func (r *GoodsRepository) DeleteGoods(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", goodsTable)
	err := r.db.BeginTxFunc(r.ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		_, err := r.db.Exec(r.ctx, query, id)
		return err
	})
	if err != nil {
		logger.Error(err)
	}
	return err
}
