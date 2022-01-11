package goods

import "my_goods/internal/entities"

// Service init structure for goods service
type Service struct {
	repo Repository
}

// NewGoodsService init func for goods service
func NewGoodsService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) getGoods(id int) *entities.Goods {
	return s.repo.getGoods(id)
}

func (s *Service) getAllGoods() *[]entities.Goods {
	return s.repo.getAllGoods()
}

func (s *Service) createGoods(good *entities.Goods) *entities.Goods {
	return s.repo.createGoods(good)
}

func (s *Service) updateGoods(good *entities.Goods) *entities.Goods {
	return s.repo.updateGoods(good)
}

func (s *Service) deleteGoods(id int) {
	s.repo.deleteGoods(id)
}
