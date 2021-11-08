package dish

import "my_goods/internal/entity"

type ServeDish interface {
	GetAllDishes() *[]entity.Dish
	GetDish(id int) *entity.Dish
	CreateDish(dish *entity.Dish) *entity.Dish
	UpdateDish(dish *entity.Dish) *entity.Dish
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

func (s *Service) GetDish(id int) *entity.Dish {
	return s.repo.GetDish(id)
}

func (s *Service) GetAllDishes() *[]entity.Dish {
	return s.repo.GetAllDishes()
}

func (s *Service) CreateDish(dish *entity.Dish) *entity.Dish {
	return s.repo.CreateDish(dish)
}

func (s *Service) UpdateDish(dish *entity.Dish) *entity.Dish {
	return s.repo.UpdateDish(dish)
}

func (s *Service) DeleteDish(id int) {
	s.repo.DeleteDish(id)
}
