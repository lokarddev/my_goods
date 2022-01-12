package repository

import (
	"gorm.io/gorm"
	"my_goods/internal/entities"
)

type DishRepo interface {
	GetAllDishes() *[]entities.Dish
	GetDish(id int) *entities.Dish
	CreateDish(dish *entities.Dish) *entities.Dish
	UpdateDish(dish *entities.Dish) *entities.Dish
	DeleteDish(id int)
}

type DishRepository struct {
	db *gorm.DB
}

func NewDishRepository(db *gorm.DB) *DishRepository {
	return &DishRepository{db: db}
}

func (r *DishRepository) GetDish(id int) *entities.Dish {
	dish := entities.Dish{}
	r.db.First(&dish, id)
	return &dish
}

func (r *DishRepository) GetAllDishes() *[]entities.Dish {
	var dish []entities.Dish
	r.db.Find(&dish)
	return &dish
}

func (r *DishRepository) CreateDish(dish *entities.Dish) *entities.Dish {
	r.db.Create(&dish)
	return dish
}

func (r *DishRepository) UpdateDish(dish *entities.Dish) *entities.Dish {
	r.db.Model(&dish).Updates(&dish)
	return dish
}

func (r *DishRepository) DeleteDish(id int) {
	r.db.Delete(&entities.Dish{}, id)
}
