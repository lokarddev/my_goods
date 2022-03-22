package delivery

import "my_goods/internal/entities"

type DishServiceInterface interface {
	GetAllDishes() (*[]entities.Dish, error)
	GetDish(id int) (*entities.Dish, error)
	CreateDish(dish *entities.Dish) (*entities.Dish, error)
	UpdateDish(dish *entities.Dish, id int) (*entities.Dish, error)
	DeleteDish(id int) error
}

type GoodsServiceInterface interface {
	GetGoods(id int) (*entities.Goods, error)
	GetAllGoods() (*[]entities.Goods, error)
	CreateGoods(good *entities.Goods) (*entities.Goods, error)
	UpdateGoods(good *entities.Goods, id int) (*entities.Goods, error)
	DeleteGoods(id int) error
}

type ListServiceInterface interface {
	GetList(id int) (*entities.List, error)
	GetAllLists() (*[]entities.List, error)
	CreateList(list *entities.List) (*entities.List, error)
	UpdateList(list *entities.List, id int) (*entities.List, error)
	DeleteList(id int) error
}
