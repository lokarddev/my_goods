package service

import (
	"my_goods/internal/entity"
	"my_goods/internal/entity/dto"
)

type ListRepo interface {
	GetList(listId, userId int32) (*dto.ListsResponse, error)
	GetAllLists(userId int32) (*[]dto.ListsResponse, error)
	CreateList(list *entity.List) (*entity.List, error)
	UpdateList(list *entity.List, listId, userId int32) (*dto.ListsResponse, error)
	DeleteList(listId, userId int32) error
	AddGoodsToList(listId int32, goods map[int32]int32) error
	AddDishToList(listId int32, dishes []int32) error
}

type GoodsRepo interface {
	GetGoods(goodsId, userId int32) (*entity.Goods, error)
	GetAllGoods(userId int32) (*[]entity.Goods, error)
	CreateGoods(good *entity.Goods) (*entity.Goods, error)
	UpdateGoods(good *entity.Goods, goodsId, userId int32) (*entity.Goods, error)
	DeleteGoods(goodsId, userId int32) error
}

type DishRepo interface {
	GetAllDishes(userId int32) (*[]dto.DishesResponse, error)
	GetDish(dishId, userId int32) (*dto.DishesResponse, error)
	CreateDish(dish *entity.Dish) (*entity.Dish, error)
	UpdateDish(dish *entity.Dish, dishId, userId int32) (*dto.DishesResponse, error)
	DeleteDish(dishId, userId int32) error
	AddGoodsToDish(dishId int32, goods map[int32]int32) error
	RemoveGoodsFromDish(dishId int32, goodsIds []int32) error
}

type UsersRepo interface {
	GetUserByName(userName string) (entity.User, bool)
	CreateUser(input dto.LoginRequest) (entity.User, error)
	CreateSession(userId int32, refresh string) (entity.Session, error)
	GetSession(token string) (entity.Session, error)
}

// AuthenticationManager provides logic for JWT & Refresh tokens generation and parsing.
type AuthenticationManager interface {
	NewAccess(userId int32) (string, error)
	Parse(accessToken string) (string, error)
	NewRefreshToken() string
}
