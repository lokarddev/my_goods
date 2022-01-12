package repository

import (
	"gorm.io/gorm"
	"my_goods/internal/entities"
)

type ListRepo interface {
	GetList(id int) *entities.List
	GetAllLists() *[]entities.List
	CreateList(list *entities.List) *entities.List
	UpdateList(list *entities.List) *entities.List
	DeleteList(id int)
}

type ListRepository struct {
	db *gorm.DB
}

func NewListRepository(db *gorm.DB) *ListRepository {
	return &ListRepository{db: db}
}

func (r *ListRepository) GetList(id int) *entities.List {
	list := entities.List{}
	r.db.First(&list, id)
	return &list
}

func (r *ListRepository) GetAllLists() *[]entities.List {
	var lists []entities.List
	r.db.Find(&lists)
	return &lists
}

func (r *ListRepository) CreateList(list *entities.List) *entities.List {
	r.db.Create(&list)
	return list
}

func (r *ListRepository) UpdateList(list *entities.List) *entities.List {
	r.db.Model(&list).Updates(&list)
	return list
}

func (r *ListRepository) DeleteList(id int) {
	r.db.Delete(&entities.List{}, id)
}
