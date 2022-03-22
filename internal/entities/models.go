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
	DishId  pgtype.Int4
	GoodsId pgtype.Int4
}

type ListsToDish struct {
	DishId  pgtype.Int4
	ListsId pgtype.Int4
}

type BaseModel struct {
	Title       pgtype.Varchar `json:"title" db:"title"`
	Description pgtype.Varchar `json:"description" db:"description"`
}
