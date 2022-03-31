package goods_service

import (
	"my_goods/internal/entity"
	"my_goods/internal/service"
)

type GoodsService struct {
	repo service.GoodsRepo
}

func NewGoodsService(repo service.GoodsRepo) *GoodsService {
	return &GoodsService{repo: repo}
}

func (s *GoodsService) GetGoods(id int32) (*entity.Goods, error) {
	return s.repo.GetGoods(id)
}

func (s *GoodsService) GetAllGoods() (*[]entity.Goods, error) {
	return s.repo.GetAllGoods()
}

func (s *GoodsService) CreateGoods(good *entity.Goods) (*entity.Goods, error) {
	return s.repo.CreateGoods(good)
}

func (s *GoodsService) UpdateGoods(good *entity.Goods, id int32) (*entity.Goods, error) {
	return s.repo.UpdateGoods(good, id)
}

func (s *GoodsService) DeleteGoods(id int32) error {
	return s.repo.DeleteGoods(id)
}
