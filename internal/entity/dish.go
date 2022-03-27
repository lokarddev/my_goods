package entity

type Dish struct {
	BaseModel
}

type CleanDish struct {
	CleanBaseModel
}

func (m *Dish) ToClean() *CleanDish {
	return &CleanDish{CleanBaseModel{
		Id:          m.Id.Int,
		Title:       m.Title.String,
		Description: m.Description.String,
	}}
}
