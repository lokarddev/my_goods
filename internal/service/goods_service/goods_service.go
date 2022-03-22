package goods_service

import (
	"my_goods/internal/entities"
	"my_goods/internal/service"
)

type GoodsService struct {
	repo service.GoodsRepo
}

func NewGoodsService(repo service.GoodsRepo) *GoodsService {
	return &GoodsService{repo: repo}
}

func (s *GoodsService) GetGoods(id int) (*entities.Goods, error) {
	return s.repo.GetGoods(id)
}

func (s *GoodsService) GetAllGoods() (*[]entities.Goods, error) {
	return s.repo.GetAllGoods()
}

func (s *GoodsService) CreateGoods(good *entities.Goods) (*entities.Goods, error) {
	return s.repo.CreateGoods(good)
}

func (s *GoodsService) UpdateGoods(good *entities.Goods, id int) (*entities.Goods, error) {
	return s.repo.UpdateGoods(good, id)
}

func (s *GoodsService) DeleteGoods(id int) error {
	return s.repo.DeleteGoods(id)
}
