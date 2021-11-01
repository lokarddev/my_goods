package dish

import "my_goods/internal/entity"

// Service init structure for dish service
type Service struct {
	repo Repository
}

// NewDishService init func for dish service
func NewDishService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) getDish(id int) *entity.Dish {
	return s.repo.getDish(id)
}

func (s *Service) getAllDishes() *[]entity.Dish {
	return s.repo.getAllDishes()
}

func (s *Service) createDish(dish *entity.Dish) *entity.Dish {
	return s.repo.createDish(dish)
}

func (s *Service) updateDish(dish *entity.Dish) *entity.Dish {
	return s.repo.updateDish(dish)
}

func (s *Service) deleteDish(id int) {
	s.repo.deleteDish(id)
}
