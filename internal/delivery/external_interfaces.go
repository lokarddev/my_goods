package delivery

import (
	"my_goods/internal/entity"
	"my_goods/internal/entity/dto"
)

type DishServiceInterface interface {
	GetAllDishes(userId int32) (*[]dto.DishesResponse, error)
	GetDish(dishId, userId int32) (*dto.DishesResponse, error)
	CreateDish(dish *entity.Dish) (*entity.Dish, error)
	UpdateDish(dish *entity.Dish, dishId, userId int32) (*dto.DishesResponse, error)
	DeleteDish(dishId, userId int32) error
	AddGoods(dishId int32, goods map[int32]int32) error
	RemoveGoodsFromDish(dishId int32, goodsIds []int32) error
}

type GoodsServiceInterface interface {
	GetGoods(goodsId, userId int32) (*entity.Goods, error)
	GetAllGoods(userId int32) (*[]entity.Goods, error)
	CreateGoods(good *entity.Goods) (*entity.Goods, error)
	UpdateGoods(good *entity.Goods, goodsId, userId int32) (*entity.Goods, error)
	DeleteGoods(goodsId, userId int32) error
}

type ListServiceInterface interface {
	GetList(listId, userId int32) (*dto.ListsResponse, error)
	GetAllLists(userId int32) (*[]dto.ListsResponse, error)
	CreateList(list *entity.List) (*entity.List, error)
	UpdateList(list *entity.List, listId, userId int32) (*dto.ListsResponse, error)
	DeleteList(listId, userId int32) error
	AddGoodsToList(listId int32, goods map[int32]int32) error
	AddDishToLIst(listId int32, dishes []int32) error
	GetShopping(listId, userId int32) (map[string]int32, error)
}

type UserServiceInterface interface {
	SignIn(input dto.LoginRequest) (dto.Access, error)
	SignUp(input dto.LoginRequest) (dto.Access, error)
	RefreshAccess(token string) (dto.Access, error)
}
