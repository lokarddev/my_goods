package service

import "my_goods/internal/entities"

type ListRepo interface {
	GetList(id int) *entities.List
	GetAllLists() *[]entities.List
	CreateList(list *entities.List) *entities.List
	UpdateList(list *entities.List) *entities.List
	DeleteList(id int)
}

type GoodsRepo interface {
	GetGoods(id int) *entities.Goods
	GetAllGoods() *[]entities.Goods
	CreateGoods(good *entities.Goods) *entities.Goods
	UpdateGoods(good *entities.Goods) *entities.Goods
	DeleteGoods(id int)
}

type DishRepo interface {
	GetAllDishes() *[]entities.Dish
	GetDish(id int) (*entities.Dish, error)
	CreateDish(dish *entities.Dish) *entities.Dish
	UpdateDish(dish *entities.Dish) *entities.Dish
	DeleteDish(id int)
}
