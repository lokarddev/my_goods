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
