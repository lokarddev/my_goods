package lists_repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"my_goods/internal/entity"
	"my_goods/internal/entity/dto"
	"my_goods/internal/storage/postgres"
	"my_goods/pkg/logger"
	"sync"
)

type ListRepository struct {
	postgres.Repository
}

func NewListRepository(db postgres.PgxPoolInterface) *ListRepository {
	return &ListRepository{Repository: postgres.Repository{
		DB:  db,
		Ctx: context.Background(),
	}}
}

func (r *ListRepository) GetList(listId, userId int32) (*dto.ListsResponse, error) {
	var err error

	list := &entity.List{}
	goods := &[]dto.GoodsWithAmount{}
	dishes := &[]dto.DishesResponse{}

	wg := &sync.WaitGroup{}
	wg.Add(3)

	go func(wg *sync.WaitGroup, id int32) {
		list, err = r.getListInfo(listId, userId)
		wg.Done()
	}(wg, listId)

	go func(wg *sync.WaitGroup, id int32) {
		dishes, err = r.getDishesInfo(listId, userId)
		wg.Done()
	}(wg, listId)

	go func(wg *sync.WaitGroup, id int32) {
		goods, err = r.getGoodsInfo(listId, userId)
		wg.Done()
	}(wg, listId)

	wg.Wait()
	return &dto.ListsResponse{
		List:   *list,
		Dishes: *dishes,
		Goods:  *goods,
	}, err
}

func (r *ListRepository) getListInfo(listId, userId int32) (*entity.List, error) {
	var list entity.PgxList
	query := fmt.Sprintf("SELECT id, title, description FROM %s WHERE id=$1 AND user_id=$2", postgres.ListsTable)
	rows, err := r.DB.Query(r.Ctx, query, listId, userId)
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

func (r *ListRepository) getDishesInfo(listId, userId int32) (*[]dto.DishesResponse, error) {
	var dishes []dto.DishesResponse
	var ids []int32
	query := fmt.Sprintf("SELECT dish_id FROM %s WHERE list_id=$1 AND user_id=$2", postgres.ListToDishes)
	rows, err := r.DB.Query(r.Ctx, query, listId, userId)
	defer rows.Close()
	for rows.Next() {
		var i int32
		if err = rows.Scan(&i); err != nil {
			logger.Error(err)
		}
		ids = append(ids, i)
	}

	for _, v := range ids {
		dish := &entity.Dish{}
		goods := &[]dto.GoodsWithAmount{}
		wg := &sync.WaitGroup{}
		wg.Add(2)
		go func(wg *sync.WaitGroup, id int32) {
			dish, err = r.GetBaseDishInfo(listId, userId)
			wg.Done()
		}(wg, v)

		go func(wg *sync.WaitGroup, id int32) {
			goods, err = r.GetBaseGoodsInfo(listId, userId)
			wg.Done()
		}(wg, v)
		wg.Wait()
		dishes = append(dishes, dto.DishesResponse{Dish: *dish, Goods: *goods})
	}

	return &dishes, err
}

func (r *ListRepository) getGoodsInfo(listId, userId int32) (*[]dto.GoodsWithAmount, error) {
	var goods []dto.GoodsWithAmount
	query := fmt.Sprintf("SELECT goods.id, goods.title, goods.description, list_goods.amount FROM %s "+
		"FULL JOIN %s ON goods.id=list_goods.goods_id WHERE list_goods.list_id=$1 AND goods.user_id=$2",
		postgres.ListToGoods, postgres.GoodsTable)
	rows, err := r.DB.Query(r.Ctx, query, listId, userId)
	defer rows.Close()
	for rows.Next() {
		good := dto.GoodsWithAmount{}
		if err = rows.Scan(&good.Id, &good.Title, &good.Description, &good.Amount); err != nil {
			logger.Error(err)
		}
		goods = append(goods, good)
	}
	return &goods, err
}

func (r *ListRepository) GetAllLists(userId int32) (*[]dto.ListsResponse, error) {
	var response []dto.ListsResponse
	var lists []entity.List

	query := fmt.Sprintf("SELECT id, title, description FROM %s", postgres.ListsTable)
	rows, err := r.DB.Query(r.Ctx, query)
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
	for _, v := range lists {
		list, err := r.GetList(v.Id, userId)
		if err != nil {
			logger.Error(err)
		}
		response = append(response, *list)
	}

	return &response, err
}

func (r *ListRepository) CreateList(list *entity.List) (*entity.List, error) {
	var pgxList entity.PgxList
	query := fmt.Sprintf("INSERT INTO %s (created_at, updated_at, title, description) VALUES (now(), now(), $1, $2) RETURNING id, title, description", postgres.ListsTable)
	err := r.DB.BeginTxFunc(r.Ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		rows, err := r.DB.Query(r.Ctx, query, list.Title, list.Description)
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

func (r *ListRepository) UpdateList(list *entity.List, listId, userId int32) (*dto.ListsResponse, error) {
	var pgxList entity.PgxList
	query := fmt.Sprintf("UPDATE %s SET updated_at=now(), title=$1, description=$2 "+
		"WHERE id=$3 AND user_id=$4 RETURNING id, title, description", postgres.ListsTable)
	err := r.DB.BeginTxFunc(r.Ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		rows, err := r.DB.Query(r.Ctx, query, list.Title, list.Description, listId, userId)
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
	return &dto.ListsResponse{List: *pgxList.ToClean()}, err
}

func (r *ListRepository) DeleteList(listId, userId int32) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 AND user_id=$2", postgres.ListsTable)
	err := r.DB.BeginTxFunc(r.Ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		_, err := r.DB.Exec(r.Ctx, query, listId, userId)
		return err
	})
	if err != nil {
		logger.Error(err)
	}
	return err
}

func (r *ListRepository) AddDishToList(listId int32, dishes []int32) error {
	b := &pgx.Batch{}
	for _, dishId := range dishes {
		query := fmt.Sprintf("INSERT INTO %s (list_id, dish_id) VALUES ($1, $2)", postgres.ListToDishes)
		b.Queue(query, listId, dishId)
	}
	err := r.DB.BeginTxFunc(r.Ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		res := tx.SendBatch(r.Ctx, b)
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
		query := fmt.Sprintf("INSERT INTO %s (list_id, goods_id, amount) VALUES ($1, $2, $3)", postgres.ListToGoods)
		b.Queue(query, listId, k, v)
	}
	err := r.DB.BeginTxFunc(r.Ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		res := tx.SendBatch(r.Ctx, b)
		err := res.Close()
		if err != nil {
			logger.Error(err)
		}
		return err
	})
	return err
}
