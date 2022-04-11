package entity

import "github.com/jackc/pgtype"

type PgxGoods struct {
	PgxBaseModel
	MeasureId pgtype.Int4 `json:"measure_id" db:"measure_id"`
}

func (m *PgxGoods) ToClean() *Goods {
	return &Goods{
		BaseModel{
			Id:          m.Id.Int,
			Title:       m.Title.String,
			Description: m.Description.String,
		},
		m.MeasureId.Int,
	}
}

type Goods struct {
	BaseModel
	MeasureId int32 `json:"measure_id"`
}

func (m *Goods) ToPgx() (*PgxGoods, error) {
	dish := &PgxGoods{}
	err := dish.Id.Set(m.Id)
	err = dish.Title.Set(m.Title)
	err = dish.Description.Set(m.Description)
	err = dish.MeasureId.Set(m.MeasureId)
	return dish, err
}
