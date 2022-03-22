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

func (s *GoodsService) GetGoods(id int) *entities.Goods {
	return s.repo.GetGoods(id)
}

func (s *GoodsService) GetAllGoods() *[]entities.Goods {
	return s.repo.GetAllGoods()
}

func (s *GoodsService) CreateGoods(good *entities.Goods) *entities.Goods {
	return s.repo.CreateGoods(good)
}

func (s *GoodsService) UpdateGoods(good *entities.Goods) *entities.Goods {
	return s.repo.UpdateGoods(good)
}

func (s *GoodsService) DeleteGoods(id int) {
	s.repo.DeleteGoods(id)
}
