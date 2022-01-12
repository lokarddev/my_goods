package repository

import (
	"context"
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
	db  PgxPoolInterface
	ctx context.Context
}

func NewListRepository(db PgxPoolInterface) *ListRepository {
	return &ListRepository{db: db, ctx: context.Background()}
}

func (r *ListRepository) GetList(id int) *entities.List {
	list := entities.List{}

	return &list
}

func (r *ListRepository) GetAllLists() *[]entities.List {
	var lists []entities.List

	return &lists
}

func (r *ListRepository) CreateList(list *entities.List) *entities.List {

	return list
}

func (r *ListRepository) UpdateList(list *entities.List) *entities.List {

	return list
}

func (r *ListRepository) DeleteList(id int) {

}
