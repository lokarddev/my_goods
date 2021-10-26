package goods

import "gorm.io/gorm"

// Repository base db struct
type Repository struct {
	db *gorm.DB
}

// NewGoodsRepo init func for goods repository
func NewGoodsRepo(db *gorm.DB) *Repository {
	return &Repository{db: db}
}
