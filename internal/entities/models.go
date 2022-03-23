package entities

import "github.com/jackc/pgtype"

type Dish struct {
	Id pgtype.Int4 `json:"id" db:"id"`
	BaseModel
}

type Goods struct {
	Id pgtype.Int4 `json:"id" db:"id"`
	BaseModel
}

type List struct {
	Dishes []Dish
	BaseModel
}

type DishToGoods struct {
	DishId  pgtype.Int4 `json:"dish_id" db:"dish_id"`
	GoodsId pgtype.Int4 `json:"goods_id" db:"goods_id"`
}

type ListsToDish struct {
	DishId  pgtype.Int4 `json:"dish_id" db:"dish_id"`
	ListsId pgtype.Int4 `json:"lists_id" db:"lists_id"`
}

type BaseModel struct {
	Title       pgtype.Varchar `json:"title" db:"title"`
	Description pgtype.Varchar `json:"description" db:"description"`
}
