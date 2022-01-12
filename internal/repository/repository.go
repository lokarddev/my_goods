package repository

import "gorm.io/gorm"

type Repository struct {
	Goods GoodsRepo
	Dish  DishRepo
	List  ListRepo
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Goods: NewGoodsRepository(db),
		Dish:  NewDishRepository(db),
		List:  NewListRepository(db),
	}
}
