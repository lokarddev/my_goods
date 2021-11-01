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

func (r *Repository) getGoods(id int) {
}

func (r *Repository) getAllGoods() {
}

func (r *Repository) createGoods() {
}

func (r *Repository) updateGoods() {
}

func (r *Repository) deleteGoods() {
}
