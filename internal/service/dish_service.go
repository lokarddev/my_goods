package service

import (
	"my_goods/internal/entities"
	"my_goods/internal/repository"
)

type DishServiceInterface interface {
	GetAllDishes() *[]entities.Dish
	GetDish(id int) (*entities.Dish, error)
	CreateDish(dish *entities.Dish) *entities.Dish
	UpdateDish(dish *entities.Dish) *entities.Dish
	DeleteDish(id int)
}

type DishService struct {
	repo repository.DishRepo
}

func NewDishService(repo repository.DishRepo) *DishService {
	return &DishService{repo: repo}
}

func (s *DishService) GetDish(id int) (*entities.Dish, error) {
	return s.repo.GetDish(id)
}

func (s *DishService) GetAllDishes() *[]entities.Dish {
	return s.repo.GetAllDishes()
}

func (s *DishService) CreateDish(dish *entities.Dish) *entities.Dish {
	return s.repo.CreateDish(dish)
}

func (s *DishService) UpdateDish(dish *entities.Dish) *entities.Dish {
	return s.repo.UpdateDish(dish)
}

func (s *DishService) DeleteDish(id int) {
	s.repo.DeleteDish(id)
}
