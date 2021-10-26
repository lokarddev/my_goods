package goods

import "gorm.io/gorm"

// GoodsRepo base db struct
type GoodsRepo struct {
	db *gorm.DB
}

// GoodsRepoInterface goods behaviour interface
type GoodsRepoInterface interface {
}

// NewGoodsRepo init func for goods repository
func NewGoodsRepo(db *gorm.DB) *GoodsRepo {
	return &GoodsRepo{db: db}
}
