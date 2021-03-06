package dto

type AddGoodsListRequest struct {
	ListId int32           `json:"list_id"`
	Ids    map[int32]int32 `json:"ids"`
}

type AddDishListRequest struct {
	ListId int32   `json:"list_id"`
	Ids    []int32 `json:"ids"`
}

type AddToDishRequest struct {
	DishId int32           `json:"dish_id"`
	Ids    map[int32]int32 `json:"ids"`
}

type RemoveFromDishRequest struct {
	DishId int32   `json:"dish_id"`
	Ids    []int32 `json:"ids"`
}

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type RefreshRequest struct {
	Refresh string `json:"refresh"`
}
