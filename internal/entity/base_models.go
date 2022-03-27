package entity

import "github.com/jackc/pgtype"

type BaseModel struct {
	Id          pgtype.Int4    `json:"id" db:"id"`
	Title       pgtype.Varchar `json:"title" db:"title"`
	Description pgtype.Varchar `json:"description" db:"description"`
}

type CleanBaseModel struct {
	Id          int32  `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
}

func (m *BaseModel) ToClean() *CleanBaseModel {
	return &CleanBaseModel{
		Id:          m.Id.Int,
		Title:       m.Title.String,
		Description: m.Description.String,
	}
}

type DishToGoods struct {
	DishId  pgtype.Int4 `json:"dish_id" db:"dish_id"`
	GoodsId pgtype.Int4 `json:"goods_id" db:"goods_id"`
}

type ListsToDish struct {
	DishId  pgtype.Int4 `json:"dish_id" db:"dish_id"`
	ListsId pgtype.Int4 `json:"lists_id" db:"lists_id"`
}
