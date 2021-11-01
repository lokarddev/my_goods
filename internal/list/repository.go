package list

import (
	"gorm.io/gorm"
	"my_goods/internal/entity"
)

// Repository base db struct
type Repository struct {
	db *gorm.DB
}

// NewListRepo init func for list repository
func NewListRepo(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) getList(id int) *entity.List {
	list := entity.List{}
	r.db.First(&list, id)
	return &list
}

func (r *Repository) getAllLists() *[]entity.List {
	var lists []entity.List
	r.db.Find(&lists)
	return &lists
}

func (r *Repository) createList(list *entity.List) *entity.List {
	r.db.Create(&list)
	return list
}

func (r *Repository) updateList(list *entity.List) *entity.List {
	r.db.Model(&list).Updates(&list)
	return list
}

func (r *Repository) deleteList(id int) {
	r.db.Delete(&entity.List{}, id)
}
