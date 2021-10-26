package list

import "gorm.io/gorm"

// ListRepo base db struct
type ListRepo struct {
	db *gorm.DB
}

// ListRepoInterface goods behaviour interface
type ListRepoInterface interface {
}

// NewListRepo init func for list repository
func NewListRepo(db *gorm.DB) *ListRepo {
	return &ListRepo{db: db}
}
