package entity

type DishesResponse struct {
	Dish  Dish
	Goods []GoodsWithAmount
}

type ListsResponse struct {
	List   List
	Dishes []Dish
	Goods  []GoodsWithAmount
}

type GoodsWithAmount struct {
	Goods
	Amount int32 `json:"amount" db:"amount"`
}
