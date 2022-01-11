package goods

import (
	"gorm.io/gorm"
	"my_goods/internal/entities"
)

// Repository base db struct
type Repository struct {
	db *gorm.DB
}

// NewGoodsRepo init func for goods repository
func NewGoodsRepo(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) getGoods(id int) *entities.Goods {
	good := entities.Goods{}
	r.db.First(&good, id)
	return &good
}

func (r *Repository) getAllGoods() *[]entities.Goods {
	var goods []entities.Goods
	r.db.Find(&goods)
	return &goods
}

func (r *Repository) createGoods(good *entities.Goods) *entities.Goods {
	r.db.Create(&good)
	return good
}

func (r *Repository) updateGoods(good *entities.Goods) *entities.Goods {
	r.db.Model(&good).Updates(&good)
	return good
}

func (r *Repository) deleteGoods(id int) {
	r.db.Delete(&entities.Goods{}, id)
}
