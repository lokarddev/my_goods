package service

import (
	"my_goods/internal/entities"
	"my_goods/internal/repository"
)

type GoodsServiceInterface interface {
	GetGoods(id int) *entities.Goods
	GetAllGoods() *[]entities.Goods
	CreateGoods(good *entities.Goods) *entities.Goods
	UpdateGoods(good *entities.Goods) *entities.Goods
	DeleteGoods(id int)
}

type GoodsService struct {
	repo repository.GoodsRepo
}

func NewGoodsService(repo repository.GoodsRepo) *GoodsService {
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
