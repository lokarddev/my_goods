package repository

import (
	"gorm.io/gorm"
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
	db *gorm.DB
}

func NewGoodsRepository(db *gorm.DB) *GoodsRepository {
	return &GoodsRepository{db: db}
}

func (r *GoodsRepository) GetGoods(id int) *entities.Goods {
	good := entities.Goods{}
	r.db.First(&good, id)
	return &good
}

func (r *GoodsRepository) GetAllGoods() *[]entities.Goods {
	var goods []entities.Goods
	r.db.Find(&goods)
	return &goods
}

func (r *GoodsRepository) CreateGoods(good *entities.Goods) *entities.Goods {
	r.db.Create(&good)
	return good
}

func (r *GoodsRepository) UpdateGoods(good *entities.Goods) *entities.Goods {
	r.db.Model(&good).Updates(&good)
	return good
}

func (r *GoodsRepository) DeleteGoods(id int) {
	r.db.Delete(&entities.Goods{}, id)
}
