package dish

import (
	"gorm.io/gorm"
	"my_goods/internal/entity"
)

// Repository base db struct
type Repository struct {
	db *gorm.DB
}

// NewDishRepo init func for dish repository
func NewDishRepo(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) getDish(id int) *entity.Dish {
	dish := entity.Dish{}
	r.db.First(&dish, id)
	return &dish
}

func (r *Repository) getAllDishes() *[]entity.Dish {
	var dish []entity.Dish
	r.db.Find(&dish)
	return &dish
}

func (r *Repository) createDish(dish *entity.Dish) *entity.Dish {
	r.db.Create(dish)
	return dish
}

func (r *Repository) updateDish(dish *entity.Dish) *entity.Dish {
	r.db.Model(&dish).Updates(dish)
	return dish
}

func (r *Repository) deleteDish(id int) {
	r.db.Delete(&entity.Dish{}, id)
}
