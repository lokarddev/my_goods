package repository

import (
	"context"
	"my_goods/internal/entities"
)

type GoodsRepo interface {
	GetGoods(id int) *entities.Goods
	GetAllGoods() *[]entities.Goods
	CreateGoods(good *entities.Goods) *entities.Goods
	UpdateGoods(good *entities.Goods) *entities.Goods
	DeleteGoods(id int)
}

type GoodsRepository struct {
	db  PgxPoolInterface
	ctx context.Context
}

func NewGoodsRepository(db PgxPoolInterface) *GoodsRepository {
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
