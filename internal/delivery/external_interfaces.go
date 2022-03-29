package delivery

import "my_goods/internal/entity"

type DishServiceInterface interface {
	GetAllDishes() (*[]entity.Dish, error)
	GetDish(id int) (*entity.Dish, error)
	CreateDish(dish *entity.Dish) (*entity.Dish, error)
	UpdateDish(dish *entity.Dish, id int) (*entity.Dish, error)
	DeleteDish(id int) error
	AddGoods(dishId int32, goods map[int32]int32) error
}

type GoodsServiceInterface interface {
	GetGoods(id int) (*entity.Goods, error)
	GetAllGoods() (*[]entity.Goods, error)
	CreateGoods(good *entity.Goods) (*entity.Goods, error)
	UpdateGoods(good *entity.Goods, id int) (*entity.Goods, error)
	DeleteGoods(id int) error
}

type ListServiceInterface interface {
	GetList(id int) (*entity.List, error)
	GetAllLists() (*[]entity.List, error)
	CreateList(list *entity.List) (*entity.List, error)
	UpdateList(list *entity.List, id int) (*entity.List, error)
	DeleteList(id int) error
}
