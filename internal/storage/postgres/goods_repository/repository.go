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

func (r *GoodsRepository) GetGoods(id int) *entities.Goods {
	good := entities.Goods{}

	return &good
}

func (r *GoodsRepository) GetAllGoods() *[]entities.Goods {
	var goods []entities.Goods

	return &goods
}

func (r *GoodsRepository) CreateGoods(good *entities.Goods) *entities.Goods {

	return good
}

func (r *GoodsRepository) UpdateGoods(good *entities.Goods) *entities.Goods {

	return good
}

func (r *GoodsRepository) DeleteGoods(id int) {

}
