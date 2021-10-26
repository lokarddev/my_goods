package list

import "gorm.io/gorm"

// Repository base db struct
type Repository struct {
	db *gorm.DB
}

// NewListRepo init func for list repository
func NewListRepo(db *gorm.DB) *Repository {
	return &Repository{db: db}
}
