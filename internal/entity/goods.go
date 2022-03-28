package entity

type PgxGoods struct {
	PgxBaseModel
}

func (m *PgxGoods) ToClean() *Goods {
	return &Goods{
		BaseModel{
			Id:          m.Id.Int,
			Title:       m.Title.String,
			Description: m.Description.String,
		},
	}
}

type Goods struct {
	BaseModel
}

func (m *Goods) ToPgx() (*PgxGoods, error) {
	dish := &PgxGoods{}
	err := dish.Id.Set(m.Id)
	err = dish.Title.Set(m.Title)
	err = dish.Description.Set(m.Description)
	return dish, err
}
