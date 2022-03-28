package entity

type PgxList struct {
	PgxBaseModel
}

func (m *PgxList) ToClean() *List {
	return &List{
		BaseModel: BaseModel{
			Id:          m.Id.Int,
			Title:       m.Title.String,
			Description: m.Description.String,
		},
		Dishes: nil,
	}
}

type List struct {
	BaseModel
	Dishes []Dish
}

func (m *List) ToPgx() (*PgxList, error) {
	dish := &PgxList{}
	err := dish.Id.Set(m.Id)
	err = dish.Title.Set(m.Title)
	err = dish.Description.Set(m.Description)
	return dish, err
}
