package dish

import "gorm.io/gorm"

// Repository base db struct
type Repository struct {
	db *gorm.DB
}

// NewDishRepo init func for dish repository
func NewDishRepo(db *gorm.DB) *Repository {
	return &Repository{db: db}
}
