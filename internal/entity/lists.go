package entity

type List struct {
	BaseModel
}

type CleanList struct {
	CleanBaseModel
	Dishes []CleanDish
}

func (m *List) ToClean() *CleanList {
	return &CleanList{
		CleanBaseModel: CleanBaseModel{
			Id:          m.Id.Int,
			Title:       m.Title.String,
			Description: m.Description.String,
		},
		Dishes: nil,
	}
}
