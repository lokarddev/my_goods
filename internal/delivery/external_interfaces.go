package delivery

import "my_goods/internal/entities"

type DishServiceInterface interface {
	GetAllDishes() *[]entities.Dish
	GetDish(id int) (*entities.Dish, error)
	CreateDish(dish *entities.Dish) *entities.Dish
	UpdateDish(dish *entities.Dish) *entities.Dish
	DeleteDish(id int)
}

type GoodsServiceInterface interface {
	GetGoods(id int) *entities.Goods
	GetAllGoods() *[]entities.Goods
	CreateGoods(good *entities.Goods) *entities.Goods
	UpdateGoods(good *entities.Goods) *entities.Goods
	DeleteGoods(id int)
}

type ListServiceInterface interface {
	GetList(id int) *entities.List
	GetAllLists() *[]entities.List
	CreateList(list *entities.List) *entities.List
	UpdateList(list *entities.List) *entities.List
	DeleteList(id int)
}
