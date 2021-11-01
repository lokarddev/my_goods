package goods

import (
	"gorm.io/gorm"
	"my_goods/internal/entity"
)

// Repository base db struct
type Repository struct {
	db *gorm.DB
}

// NewGoodsRepo init func for goods repository
func NewGoodsRepo(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) getGoods(id int) *entity.Goods {
	good := entity.Goods{}
	r.db.First(&good, id)
	return &good
}

func (r *Repository) getAllGoods() *[]entity.Goods {
	var goods []entity.Goods
	r.db.Find(&goods)
	return &goods
}

func (r *Repository) createGoods(good *entity.Goods) *entity.Goods {
	r.db.Create(good)
	return good
}

func (r *Repository) updateGoods(good *entity.Goods) *entity.Goods {
	r.db.Model(&good).Updates(good)
	return good
}

func (r *Repository) deleteGoods(id int) {
	r.db.Delete(&entity.Goods{}, id)
}
