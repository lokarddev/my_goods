package service

import "my_goods/internal/repository"

type Service struct {
	Dish  DishServiceInterface
	Goods GoodsServiceInterface
	List  ListServiceInterface
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Dish:  NewDishService(repo.Dish),
		Goods: NewGoodsService(repo.Goods),
		List:  NewListService(repo.List),
	}
}
