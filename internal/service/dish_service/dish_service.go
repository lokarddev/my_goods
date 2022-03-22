package dish_service

import (
	"my_goods/internal/entities"
	"my_goods/internal/service"
)

type DishService struct {
	repo service.DishRepo
}

func NewDishService(repo service.DishRepo) *DishService {
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
