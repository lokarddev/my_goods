package entity

type PgxDish struct {
	PgxBaseModel
}

func (m *PgxDish) ToClean() *Dish {
	return &Dish{BaseModel{
		Id:          m.Id.Int,
		Title:       m.Title.String,
		Description: m.Description.String,
	}}
}

type Dish struct {
	BaseModel
}

func (m *Dish) ToPgx() (*PgxDish, error) {
	dish := &PgxDish{}
	err := dish.Id.Set(m.Id)
	err = dish.Title.Set(m.Title)
	err = dish.Description.Set(m.Description)
	return dish, err
}
