package services

import "my_goods/internal/repos"

// DishServiceInterface dish service behaviour
type DishServiceInterface interface {
}

// DishService init structure for dish service
type DishService struct {
	repo repos.DishRepoInterface
}

// NewDishService init func for dish service
func NewDishService(repo repos.DishRepoInterface) *DishService {
	return &DishService{repo: repo}
}
