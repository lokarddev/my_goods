package postgres

import (
	"context"
	"fmt"
	"my_goods/internal/entity"
	"my_goods/pkg/logger"
)

const (
	DishTable      = "dishes"
	DishGoodsTable = "dish_goods"
	GoodsTable     = "goods"
	ListsTable     = "lists"
	ListToDishes   = "list_dishes"
	ListToGoods    = "list_goods"
)

type Repository struct {
	DB  PgxPoolInterface
	Ctx context.Context
}

func (r *Repository) GetBaseGoodsInfo(id int32) (*[]entity.GoodsWithAmount, error) {
	var goods []entity.GoodsWithAmount
	query := fmt.Sprintf("SELECT goods.Id, goods.title, goods.description, dish_goods.amount FROM %s "+
		"FULL JOIN %s ON dish_goods.goods_id = goods.id WHERE dish_id=$1", DishGoodsTable, GoodsTable)
	rows, err := r.DB.Query(r.Ctx, query, id)
	defer rows.Close()
	for rows.Next() {
		goodInfo := entity.GoodsWithAmount{}
		if err = rows.Scan(&goodInfo.Id, &goodInfo.Title, &goodInfo.Description, &goodInfo.Amount); err != nil {
			logger.Error(err)
		}
		goods = append(goods, goodInfo)
	}
	return &goods, err
}

func (r *Repository) GetBaseDishInfo(id int32) (*entity.Dish, error) {
	var pgxDish entity.PgxDish
	query := fmt.Sprintf("SELECT id, title, description FROM %s WHERE id=$1", DishTable)
	rows, err := r.DB.Query(r.Ctx, query, id)
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&pgxDish.Id, &pgxDish.Title, &pgxDish.Description); err != nil {
			logger.Error(err)
		}
	}
	if err != nil {
		logger.Error(err)
	}
	return pgxDish.ToClean(), err
}
