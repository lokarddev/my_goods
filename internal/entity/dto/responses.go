package dto

import "my_goods/internal/entity"

type DishesResponse struct {
	Dish  entity.Dish
	Goods []GoodsWithAmount
}

type ListsResponse struct {
	List   entity.List
	Dishes []DishesResponse
	Goods  []GoodsWithAmount
}

type GoodsWithAmount struct {
	entity.Goods
	Amount int32 `json:"amount" db:"amount"`
}

type Access struct {
	Access  string
	Refresh string
}
