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

func NewListRepository(db postgres.PgxPoolInterface) *ListRepository {
	return &ListRepository{db: db, ctx: context.Background()}
}

func (r *ListRepository) GetList(id int) (*entity.List, error) {
	var list entity.PgxList
	query := fmt.Sprintf("SELECT id, title, description FROM %s WHERE id=$1", postgres.ListsTable)
	rows, err := r.db.Query(r.ctx, query, id)
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&list.Id, &list.Title, &list.Description); err != nil {
			logger.Error(err)
		}
	}
	if err != nil {
		logger.Error(err)
	}
	return list.ToClean(), err
}

func (r *ListRepository) GetAllLists() (*[]entity.List, error) {
	var lists []entity.List
	query := fmt.Sprintf("SELECT * FROM %s", postgres.ListsTable)
	rows, err := r.db.Query(r.ctx, query)
	defer rows.Close()
	for rows.Next() {
		var list entity.PgxList
		if err = rows.Scan(&list.Id, &list.Title, &list.Description); err != nil {
			logger.Error(err)
		}
		lists = append(lists, *list.ToClean())
	}
	if err != nil {
		logger.Error(err)
	}
	return &lists, err
}

func (r *ListRepository) CreateList(list *entity.List) (*entity.List, error) {
	var pgxList entity.PgxList
	query := fmt.Sprintf("INSERT INTO %s (created_at, updated_at, title, description) VALUES (now(), now(), $1, $2) RETURNING id, title, description", postgres.ListsTable)
	err := r.db.BeginTxFunc(r.ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		rows, err := r.db.Query(r.ctx, query, list.Title, list.Description)
		defer rows.Close()
		for rows.Next() {
			if err = rows.Scan(&pgxList.Id, &pgxList.Title, &pgxList.Description); err != nil {
				return err
			}
		}
		return err
	})
	if err != nil {
		logger.Error(err)
	}
	return pgxList.ToClean(), err
}

func (r *ListRepository) UpdateList(list *entity.List, id int) (*entity.List, error) {
	var pgxList entity.PgxList
	query := fmt.Sprintf("UPDATE %s SET updated_at=now(), title=$1, description=$2 WHERE id=$3 RETURNING id, title, description", postgres.ListsTable)
	err := r.db.BeginTxFunc(r.ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		rows, err := r.db.Query(r.ctx, query, list.Title, list.Description, id)
		defer rows.Close()
		for rows.Next() {
			if err = rows.Scan(&pgxList.Id, &pgxList.Title, &pgxList.Description); err != nil {
				return err
			}
		}
		return err
	})
	if err != nil {
		logger.Error(err)
	}
	return pgxList.ToClean(), err
}

func (r *ListRepository) DeleteList(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", postgres.ListsTable)
	err := r.db.BeginTxFunc(r.ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		_, err := r.db.Exec(r.ctx, query, id)
		return err
	})
	if err != nil {
		logger.Error(err)
	}
	return err
}

func (r *ListRepository) AddDishToList(listId int32, dishes map[int32]int32) error {
	b := &pgx.Batch{}
	for k, v := range dishes {
		query := fmt.Sprintf("INSERT INTO %s (dish_id, goods_id, amount) VALUES ($1, $2, $3)", postgres.ListToDishes)
		b.Queue(query, listId, k, v)
	}
	err := r.db.BeginTxFunc(r.ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		res := tx.SendBatch(r.ctx, b)
		err := res.Close()
		if err != nil {
			logger.Error(err)
		}
		return err
	})
	return err
}

func (r *ListRepository) AddGoodsToList(listId int32, goods map[int32]int32) error {
	b := &pgx.Batch{}
	for k, v := range goods {
		query := fmt.Sprintf("INSERT INTO %s (dish_id, goods_id, amount) VALUES ($1, $2, $3)", postgres.ListToGoods)
		b.Queue(query, listId, k, v)
	}
	err := r.db.BeginTxFunc(r.ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		res := tx.SendBatch(r.ctx, b)
		err := res.Close()
		if err != nil {
			logger.Error(err)
		}
		return err
	})
	return err
}
