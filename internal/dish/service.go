package dish

import "my_goods/internal/entities"

type ServeDish interface {
	GetAllDishes() *[]entities.Dish
	GetDish(id int) *entities.Dish
	CreateDish(dish *entities.Dish) *entities.Dish
	UpdateDish(dish *entities.Dish) *entities.Dish
	DeleteDish(id int)
}

// Service init structure for dish service
type Service struct {
	repo RepoDish
}

// NewDishService init func for dish service
func NewDishService(repo RepoDish) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetDish(id int) *entities.Dish {
	return s.repo.GetDish(id)
}

func (s *Service) GetAllDishes() *[]entities.Dish {
	return s.repo.GetAllDishes()
}

func (s *Service) CreateDish(dish *entities.Dish) *entities.Dish {
	return s.repo.CreateDish(dish)
}

func (s *Service) UpdateDish(dish *entities.Dish) *entities.Dish {
	return s.repo.UpdateDish(dish)
}

func (s *Service) DeleteDish(id int) {
	s.repo.DeleteDish(id)
}
