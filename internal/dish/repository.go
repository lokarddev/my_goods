package dish

import (
	"gorm.io/gorm"
	"my_goods/internal/entity"
)

type RepoDish interface {
	GetAllDishes() *[]entity.Dish
	GetDish(id int) *entity.Dish
	CreateDish(dish *entity.Dish) *entity.Dish
	UpdateDish(dish *entity.Dish) *entity.Dish
	DeleteDish(id int)
}

// Repository base db struct
type Repository struct {
	db *gorm.DB
}

// NewDishRepo init func for dish repository
func NewDishRepo(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetDish(id int) *entity.Dish {
	dish := entity.Dish{}
	r.db.First(&dish, id)
	return &dish
}

func (r *Repository) GetAllDishes() *[]entity.Dish {
	var dish []entity.Dish
	r.db.Find(&dish)
	return &dish
}

func (r *Repository) CreateDish(dish *entity.Dish) *entity.Dish {
	r.db.Create(&dish)
	return dish
}

func (r *Repository) UpdateDish(dish *entity.Dish) *entity.Dish {
	r.db.Model(&dish).Updates(&dish)
	return dish
}

func (r *Repository) DeleteDish(id int) {
	r.db.Delete(&entity.Dish{}, id)
}
