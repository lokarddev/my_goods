package goods_repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"my_goods/internal/entities"
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

func (r *GoodsRepository) GetGoods(id int) (*entities.Goods, error) {
	var good entities.Goods
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

func (r *GoodsRepository) GetAllGoods() (*[]entities.Goods, error) {
	var goods []entities.Goods
	query := fmt.Sprintf("SELECT * FROM %s", goodsTable)
	rows, err := r.db.Query(r.ctx, query)
	defer rows.Close()
	for rows.Next() {
		var good entities.Goods
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

func (r *GoodsRepository) CreateGoods(good *entities.Goods) (*entities.Goods, error) {
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

func (r *GoodsRepository) UpdateGoods(good *entities.Goods, id int) (*entities.Goods, error) {
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
