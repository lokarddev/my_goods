package services

import "my_goods/internal/repos"

// Service base service structure contains all service interfaces
type Service struct {
	GoodsServiceInterface
	DishServiceInterface
	ListServiceInterface
}

// NewService init func for all services
func NewService(repo *repos.Repository) *Service {
	return &Service{
		GoodsServiceInterface: NewGoodsService(repo.GoodsRepoInterface),
		DishServiceInterface:  NewDishService(repo.DishRepoInterface),
		ListServiceInterface:  NewListService(repo.ListRepoInterface),
	}
}
