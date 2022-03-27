package entity

type Goods struct {
	BaseModel
}

type CleanGoods struct {
	CleanBaseModel
}

func (m *Goods) ToClean() *CleanGoods {
	return &CleanGoods{
		CleanBaseModel{
			Id:          m.Id.Int,
			Title:       m.Title.String,
			Description: m.Description.String,
		},
	}
}
