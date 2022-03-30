package entity

type AddToListRequest struct {
	ListId int32           `json:"dish_id"`
	Ids    map[int32]int32 `json:"ids"`
}

type AddToDishRequest struct {
	DishId int32           `json:"dish_id"`
	Ids    map[int32]int32 `json:"ids"`
}

type RemoveFromDishRequest struct {
	DishId int32   `json:"dish_id"`
	Ids    []int32 `json:"ids"`
}
