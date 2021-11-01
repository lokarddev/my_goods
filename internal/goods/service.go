package goods

import "my_goods/internal/entity"

// Service init structure for goods service
type Service struct {
	repo Repository
}

// NewGoodsService init func for goods service
func NewGoodsService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) getGoods(id int) *entity.Goods {
	return s.repo.getGoods(id)
}

func (s *Service) getAllGoods() *[]entity.Goods {
	return s.repo.getAllGoods()
}

func (s *Service) createGoods(good *entity.Goods) *entity.Goods {
	return s.repo.createGoods(good)
}

func (s *Service) updateGoods(good *entity.Goods) *entity.Goods {
	return s.repo.updateGoods(good)
}

func (s *Service) deleteGoods(id int) {
	s.repo.deleteGoods(id)
}
