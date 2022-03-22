package goods_repository

import (
	"context"
	"my_goods/internal/entities"
	"my_goods/internal/storage/postgres"
)

type GoodsRepository struct {
	db  postgres.PgxPoolInterface
	ctx context.Context
}

func NewGoodsRepository(db postgres.PgxPoolInterface) *GoodsRepository {
	return &GoodsRepository{db: db, ctx: context.Background()}
}

func (r *GoodsRepository) GetGoods(id int) (*entities.Goods, error) {
	good := entities.Goods{}
	return &good, nil
}

func (r *GoodsRepository) GetAllGoods() (*[]entities.Goods, error) {
	var goods []entities.Goods
	return &goods, nil
}

func (r *GoodsRepository) CreateGoods(good *entities.Goods) (*entities.Goods, error) {
	return good, nil
}

func (r *GoodsRepository) UpdateGoods(good *entities.Goods, id int) (*entities.Goods, error) {
	return good, nil
}

func (r *GoodsRepository) DeleteGoods(id int) error {
	return nil
}
