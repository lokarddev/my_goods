package services

import "my_goods/internal/repos"

// GoodsServiceInterface goods service behaviour
type GoodsServiceInterface interface {
}

// GoodsService init structure for goods service
type GoodsService struct {
	repo repos.GoodsRepoInterface
}

// NewGoodsService init func for goods service
func NewGoodsService(repo repos.GoodsRepoInterface) *GoodsService {
	return &GoodsService{repo: repo}
}
