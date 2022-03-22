package lists_repository

import (
	"context"
	"my_goods/internal/entities"
	"my_goods/internal/storage/postgres"
)

type ListRepository struct {
	db  postgres.PgxPoolInterface
	ctx context.Context
}

func NewListRepository(db postgres.PgxPoolInterface) *ListRepository {
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
