package service

import "my_goods/internal/entity"

type ListRepo interface {
	GetList(id int) (*entity.List, error)
	GetAllLists() (*[]entity.List, error)
	CreateList(list *entity.List) (*entity.List, error)
	UpdateList(list *entity.List, id int) (*entity.List, error)
	DeleteList(id int) error
	AddGoodsToList(listId int32, goods map[int32]int32) error
	AddDishToList(listId int32, dishes map[int32]int32) error
}

type GoodsRepo interface {
	GetGoods(id int) (*entity.Goods, error)
	GetAllGoods() (*[]entity.Goods, error)
	CreateGoods(good *entity.Goods) (*entity.Goods, error)
	UpdateGoods(good *entity.Goods, id int) (*entity.Goods, error)
	DeleteGoods(id int) error
}

type DishRepo interface {
	GetAllDishes() (*[]entity.DishesResponse, error)
	GetDish(id int32) (*entity.DishesResponse, error)
	CreateDish(dish *entity.Dish) (*entity.Dish, error)
	UpdateDish(dish *entity.Dish, id int32) (*entity.DishesResponse, error)
	DeleteDish(id int32) error
	AddGoodsToDish(dishId int32, goods map[int32]int32) error
	RemoveGoodsFromDish(dishId int32, goodsIds []int32) error
}
