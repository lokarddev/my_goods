package delivery

import (
	"my_goods/internal/entity"
	"my_goods/internal/entity/dto"
)

type DishServiceInterface interface {
	GetAllDishes() (*[]dto.DishesResponse, error)
	GetDish(id int32) (*dto.DishesResponse, error)
	CreateDish(dish *entity.Dish) (*entity.Dish, error)
	UpdateDish(dish *entity.Dish, id int32) (*dto.DishesResponse, error)
	DeleteDish(id int32) error
	AddGoods(dishId int32, goods map[int32]int32) error
	RemoveGoodsFromDish(dishId int32, goodsIds []int32) error
}

type GoodsServiceInterface interface {
	GetGoods(id int32) (*entity.Goods, error)
	GetAllGoods() (*[]entity.Goods, error)
	CreateGoods(good *entity.Goods) (*entity.Goods, error)
	UpdateGoods(good *entity.Goods, id int32) (*entity.Goods, error)
	DeleteGoods(id int32) error
}

type ListServiceInterface interface {
	GetList(id int32) (*dto.ListsResponse, error)
	GetAllLists() (*[]dto.ListsResponse, error)
	CreateList(list *entity.List) (*entity.List, error)
	UpdateList(list *entity.List, id int32) (*dto.ListsResponse, error)
	DeleteList(id int32) error
	AddGoodsToList(listId int32, goods map[int32]int32) error
	AddDishToLIst(listId int32, dishes []int32) error
}

type UserServiceInterface interface {
	SignIn(input dto.LoginRequest) (dto.Access, error)
	SignUp(input dto.LoginRequest) (dto.Access, error)
	RefreshAccess(token string) (dto.Access, error)
}
