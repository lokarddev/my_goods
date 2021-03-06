package dish_service

import (
	"my_goods/internal/entity"
	"my_goods/internal/entity/dto"
	"my_goods/internal/service"
)

type DishService struct {
	repo service.DishRepo
}

func NewDishService(repo service.DishRepo) *DishService {
	return &DishService{repo: repo}
}

func (s *DishService) GetDish(dishId, userId int32) (*dto.DishesResponse, error) {
	return s.repo.GetDish(dishId, userId)
}

func (s *DishService) GetAllDishes(userId int32) (*[]dto.DishesResponse, error) {
	return s.repo.GetAllDishes(userId)
}

func (s *DishService) CreateDish(dish *entity.Dish) (*entity.Dish, error) {
	return s.repo.CreateDish(dish)
}

func (s *DishService) UpdateDish(dish *entity.Dish, dishId, userId int32) (*dto.DishesResponse, error) {
	return s.repo.UpdateDish(dish, dishId, userId)
}

func (s *DishService) DeleteDish(dishId, userId int32) error {
	return s.repo.DeleteDish(dishId, userId)
}

func (s *DishService) AddGoods(dishId int32, goods map[int32]int32) error {
	return s.repo.AddGoodsToDish(dishId, goods)
}

func (s *DishService) RemoveGoodsFromDish(dishId int32, goodsIds []int32) error {
	return s.repo.RemoveGoodsFromDish(dishId, goodsIds)
}
