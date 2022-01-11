package list

import (
	"gorm.io/gorm"
	"my_goods/internal/entities"
)

// Repository base db struct
type Repository struct {
	db *gorm.DB
}

// NewListRepo init func for list repository
func NewListRepo(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) getList(id int) *entities.List {
	list := entities.List{}
	r.db.First(&list, id)
	return &list
}

func (r *Repository) getAllLists() *[]entities.List {
	var lists []entities.List
	r.db.Find(&lists)
	return &lists
}

func (r *Repository) createList(list *entities.List) *entities.List {
	r.db.Create(&list)
	return list
}

func (r *Repository) updateList(list *entities.List) *entities.List {
	r.db.Model(&list).Updates(&list)
	return list
}

func (r *Repository) deleteList(id int) {
	r.db.Delete(&entities.List{}, id)
}
