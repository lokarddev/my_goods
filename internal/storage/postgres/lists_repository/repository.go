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

func (r *ListRepository) GetList(id int) (*entities.List, error) {
	list := entities.List{}
	return &list, nil
}

func (r *ListRepository) GetAllLists() (*[]entities.List, error) {
	var lists []entities.List
	return &lists, nil
}

func (r *ListRepository) CreateList(list *entities.List) (*entities.List, error) {
	return list, nil
}

func (r *ListRepository) UpdateList(list *entities.List, id int) (*entities.List, error) {
	return list, nil
}

func (r *ListRepository) DeleteList(id int) error {
	return nil
}
