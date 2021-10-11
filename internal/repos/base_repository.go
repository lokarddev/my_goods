package repos

import "gorm.io/gorm"

// Repository base repository init struct, contains all behaviour interfaces
type Repository struct {
	GoodsRepoInterface
	DishRepoInterface
	ListRepoInterface
}

// NewRepository init func for base repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		GoodsRepoInterface: NewDishRepo(db),
		DishRepoInterface:  NewDishRepo(db),
		ListRepoInterface:  NewDishRepo(db),
	}
}
