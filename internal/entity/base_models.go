package entity

import "github.com/jackc/pgtype"

type PgxBaseModel struct {
	Id          pgtype.Int4    `json:"id" db:"id"`
	Title       pgtype.Varchar `json:"title" db:"title"`
	Description pgtype.Varchar `json:"description" db:"description"`
}

type BaseModel struct {
	Id          int32  `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
}

type DishToGoods struct {
	DishId  pgtype.Int4 `json:"dish_id" db:"dish_id"`
	GoodsId pgtype.Int4 `json:"goods_id" db:"goods_id"`
}

type ListsToDish struct {
	DishId  pgtype.Int4 `json:"dish_id" db:"dish_id"`
	ListsId pgtype.Int4 `json:"lists_id" db:"lists_id"`
}
