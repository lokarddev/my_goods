package postgres

import (
	"context"
	"fmt"
	"my_goods/internal/entity"
	"my_goods/internal/entity/dto"
	"my_goods/pkg/logger"
)

const (
	DishTable      = "dishes"
	DishGoodsTable = "dish_goods"
	GoodsTable     = "goods"
	ListsTable     = "lists"
	ListToDishes   = "list_dishes"
	ListToGoods    = "list_goods"
	Session        = "refresh_session"
	Users          = "users"
)

type Repository struct {
	DB  PgxPoolInterface
	Ctx context.Context
}

func (r *Repository) GetBaseGoodsInfo(dishId int32, userId int32) (*[]dto.GoodsWithAmount, error) {
	var goods []dto.GoodsWithAmount
	query := fmt.Sprintf("SELECT goods.Id, goods.title, goods.description, dish_goods.amount FROM %s "+
		"FULL JOIN %s ON dish_goods.goods_id = goods.id WHERE dish_id=$1 AND user_id=$2", DishGoodsTable, GoodsTable)
	rows, err := r.DB.Query(r.Ctx, query, dishId, userId)
	defer rows.Close()
	for rows.Next() {
		goodInfo := dto.GoodsWithAmount{}
		if err = rows.Scan(&goodInfo.Id, &goodInfo.Title, &goodInfo.Description, &goodInfo.Amount); err != nil {
			logger.Error(err)
		}
		goods = append(goods, goodInfo)
	}
	return &goods, err
}

func (r *Repository) GetBaseDishInfo(dishId int32, userId int32) (*entity.Dish, error) {
	var pgxDish entity.PgxDish
	query := fmt.Sprintf("SELECT id, title, description FROM %s WHERE id=$1 AND user_id=$2", DishTable)
	rows, err := r.DB.Query(r.Ctx, query, dishId, userId)
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
