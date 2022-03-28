package dish_service

import (
	"my_goods/internal/entity"
	"my_goods/internal/service"
)

type DishService struct {
	repo service.DishRepo
}

func NewDishService(repo service.DishRepo) *DishService {
	return &DishService{repo: repo}
}

func (s *DishService) GetDish(id int) (*entity.Dish, error) {
	return s.repo.GetDish(id)
}

func (s *DishService) GetAllDishes() (*[]entity.Dish, error) {
	return s.repo.GetAllDishes()
}

func (s *DishService) CreateDish(dish *entity.Dish) (*entity.Dish, error) {
	return s.repo.CreateDish(dish)
}

func (s *DishService) UpdateDish(dish *entity.Dish, id int) (*entity.Dish, error) {
	return s.repo.UpdateDish(dish, id)
}

func (s *DishService) DeleteDish(id int) error {
	return s.repo.DeleteDish(id)
}
