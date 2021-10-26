package dish

import "gorm.io/gorm"

// DishRepo base db struct
type DishRepo struct {
	db *gorm.DB
}

// DishRepoInterface goods behaviour interface
type DishRepoInterface interface {
}

// NewDishRepo init func for dish repository
func NewDishRepo(db *gorm.DB) *DishRepo {
	return &DishRepo{db: db}
}
