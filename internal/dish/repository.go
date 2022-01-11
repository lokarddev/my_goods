package dish

import (
	"gorm.io/gorm"
	"my_goods/internal/entities"
)

type RepoDish interface {
	GetAllDishes() *[]entities.Dish
	GetDish(id int) *entities.Dish
	CreateDish(dish *entities.Dish) *entities.Dish
	UpdateDish(dish *entities.Dish) *entities.Dish
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

func (r *Repository) GetDish(id int) *entities.Dish {
	dish := entities.Dish{}
	r.db.First(&dish, id)
	return &dish
}

func (r *Repository) GetAllDishes() *[]entities.Dish {
	var dish []entities.Dish
	r.db.Find(&dish)
	return &dish
}

func (r *Repository) CreateDish(dish *entities.Dish) *entities.Dish {
	r.db.Create(&dish)
	return dish
}

func (r *Repository) UpdateDish(dish *entities.Dish) *entities.Dish {
	r.db.Model(&dish).Updates(&dish)
	return dish
}

func (r *Repository) DeleteDish(id int) {
	r.db.Delete(&entities.Dish{}, id)
}
