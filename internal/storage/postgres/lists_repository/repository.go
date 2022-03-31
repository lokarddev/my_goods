package lists_repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"my_goods/internal/entity"
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

func (r *ListRepository) GetList(id int32) (*entity.ListsResponse, error) {
	var err error

	list := &entity.List{}
	goods := &[]entity.GoodsWithAmount{}
	dishes := &[]entity.DishesResponse{}

	wg := &sync.WaitGroup{}
	wg.Add(3)

	go func(wg *sync.WaitGroup, id int32) {
		list, err = r.getListInfo(id)
		wg.Done()
	}(wg, id)

	go func(wg *sync.WaitGroup, id int32) {
		dishes, err = r.getDishesInfo(id)
		wg.Done()
	}(wg, id)

	go func(wg *sync.WaitGroup, id int32) {
		goods, err = r.getGoodsInfo(id)
		wg.Done()
	}(wg, id)

	wg.Wait()
	return &entity.ListsResponse{
		List:   *list,
		Dishes: *dishes,
		Goods:  *goods,
	}, err
}

func (r *ListRepository) getListInfo(id int32) (*entity.List, error) {
	var list entity.PgxList
	query := fmt.Sprintf("SELECT id, title, description FROM %s WHERE id=$1", postgres.ListsTable)
	rows, err := r.DB.Query(r.Ctx, query, id)
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

func (r *ListRepository) getDishesInfo(id int32) (*[]entity.DishesResponse, error) {
	var dishes []entity.DishesResponse
	var ids []int32
	query := fmt.Sprintf("SELECT dish_id FROM %s WHERE list_id=$1", postgres.ListToDishes)
	rows, err := r.DB.Query(r.Ctx, query, id)
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
		goods := &[]entity.GoodsWithAmount{}
		wg := &sync.WaitGroup{}
		wg.Add(2)
		go func(wg *sync.WaitGroup, id int32) {
			dish, err = r.GetBaseDishInfo(id)
			wg.Done()
		}(wg, v)

		go func(wg *sync.WaitGroup, id int32) {
			goods, err = r.GetBaseGoodsInfo(id)
			wg.Done()
		}(wg, v)
		wg.Wait()
		dishes = append(dishes, entity.DishesResponse{Dish: *dish, Goods: *goods})
	}

	return &dishes, err
}

func (r *ListRepository) getGoodsInfo(id int32) (*[]entity.GoodsWithAmount, error) {
	var goods []entity.GoodsWithAmount
	query := fmt.Sprintf("select goods.id, goods.title, goods.description, list_goods.amount from %s "+
		"full join %s on goods.id=list_goods.goods_id where list_goods.list_id=$1",
		postgres.ListToGoods, postgres.GoodsTable)
	rows, err := r.DB.Query(r.Ctx, query, id)
	defer rows.Close()
	for rows.Next() {
		good := entity.GoodsWithAmount{}
		if err = rows.Scan(&good.Id, &good.Title, &good.Description, &good.Amount); err != nil {
			logger.Error(err)
		}
		goods = append(goods, good)
	}
	return &goods, err
}

func (r *ListRepository) GetAllLists() (*[]entity.ListsResponse, error) {
	var response []entity.ListsResponse
	var lists []entity.List

	goods := &[]entity.GoodsWithAmount{}
	dishes := &[]entity.DishesResponse{}

	query := fmt.Sprintf("SELECT id, title FROM %s", postgres.ListsTable)
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

	wg := &sync.WaitGroup{}
	wg.Add(3)
	for _, list := range lists {
		go func(wg *sync.WaitGroup, id int32) {
			dishes, err = r.getDishesInfo(id)
			wg.Done()
		}(wg, list.Id)

		go func(wg *sync.WaitGroup, id int32) {
			goods, err = r.getGoodsInfo(id)
			wg.Done()
		}(wg, list.Id)
		response = append(response, entity.ListsResponse{
			List:   list,
			Dishes: *dishes,
			Goods:  *goods,
		})
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

func (r *ListRepository) UpdateList(list *entity.List, id int32) (*entity.ListsResponse, error) {
	var pgxList entity.PgxList
	query := fmt.Sprintf("UPDATE %s SET updated_at=now(), title=$1, description=$2 WHERE id=$3 RETURNING id, title, description", postgres.ListsTable)
	err := r.DB.BeginTxFunc(r.Ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		rows, err := r.DB.Query(r.Ctx, query, list.Title, list.Description, id)
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
	return &entity.ListsResponse{List: *pgxList.ToClean()}, err
}

func (r *ListRepository) DeleteList(id int32) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", postgres.ListsTable)
	err := r.DB.BeginTxFunc(r.Ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		_, err := r.DB.Exec(r.Ctx, query, id)
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
		query := fmt.Sprintf("INSERT INTO %s (dish_id, goods_id, amount) VALUES ($1, $2, $3)", postgres.ListToGoods)
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
