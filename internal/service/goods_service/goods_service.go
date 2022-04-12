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

func (s *GoodsService) GetGoods(goodsId, userId int32) (*entity.Goods, error) {
	return s.repo.GetGoods(goodsId, userId)
}

func (s *GoodsService) GetAllGoods(userId int32) (*[]entity.Goods, error) {
	return s.repo.GetAllGoods(userId)
}

func (s *GoodsService) CreateGoods(good *entity.Goods) (*entity.Goods, error) {
	return s.repo.CreateGoods(good)
}

func (s *GoodsService) UpdateGoods(good *entity.Goods, goodsId, userId int32) (*entity.Goods, error) {
	return s.repo.UpdateGoods(good, goodsId, userId)
}

func (s *GoodsService) DeleteGoods(goodsId, userId int32) error {
	return s.repo.DeleteGoods(goodsId, userId)
}
