package lists_repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"my_goods/internal/entity"
	"my_goods/internal/storage/postgres"
	"my_goods/pkg/logger"
)

type ListRepository struct {
	db  postgres.PgxPoolInterface
	ctx context.Context
}

const (
	listsTable = "lists"
)

func NewListRepository(db postgres.PgxPoolInterface) *ListRepository {
	return &ListRepository{db: db, ctx: context.Background()}
}

func (r *ListRepository) GetList(id int) (*entity.List, error) {
	var list entity.List
	query := fmt.Sprintf("SELECT id, title, description FROM %s WHERE id=$1", listsTable)
	rows, err := r.db.Query(r.ctx, query, id)
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&list); err != nil {
			logger.Error(err)
		}
	}
	if err != nil {
		logger.Error(err)
	}
	return &list, err
}

func (r *ListRepository) GetAllLists() (*[]entity.List, error) {
	var lists []entity.List
	query := fmt.Sprintf("SELECT * FROM %s", listsTable)
	rows, err := r.db.Query(r.ctx, query)
	defer rows.Close()
	for rows.Next() {
		var good entity.List
		if err = rows.Scan(&good); err != nil {
			logger.Error(err)
		}
		lists = append(lists, good)
	}
	if err != nil {
		logger.Error(err)
	}
	return &lists, err
}

func (r *ListRepository) CreateList(list *entity.List) (*entity.List, error) {
	query := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING *", listsTable)
	err := r.db.BeginTxFunc(r.ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		rows, err := r.db.Query(r.ctx, query, list.Title, list.Description)
		defer rows.Close()
		for rows.Next() {
			if err = rows.Scan(list); err != nil {
				return err
			}
		}
		return err
	})
	if err != nil {
		logger.Error(err)
	}
	return list, err
}

func (r *ListRepository) UpdateList(list *entity.List, id int) (*entity.List, error) {
	query := fmt.Sprintf("UPDATE %s SET title=$1, description=$2 WHERE id=$3 RETURNING *", listsTable)
	err := r.db.BeginTxFunc(r.ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		rows, err := r.db.Query(r.ctx, query, list.Title, list.Description, id)
		defer rows.Close()
		for rows.Next() {
			if err = rows.Scan(list); err != nil {
				return err
			}
		}
		return err
	})
	if err != nil {
		logger.Error(err)
	}
	return list, err
}

func (r *ListRepository) DeleteList(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", listsTable)
	err := r.db.BeginTxFunc(r.ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		_, err := r.db.Exec(r.ctx, query, id)
		return err
	})
	if err != nil {
		logger.Error(err)
	}
	return err
}
