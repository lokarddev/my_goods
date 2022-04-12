package lists_service

import (
	"my_goods/internal/entity"
	"my_goods/internal/entity/dto"
	"my_goods/internal/service"
)

type ListService struct {
	repo service.ListRepo
}

func NewListService(repo service.ListRepo) *ListService {
	return &ListService{repo: repo}
}

func (s *ListService) GetList(listId, userId int32) (*dto.ListsResponse, error) {
	return s.repo.GetList(listId, userId)
}

func (s *ListService) GetAllLists(userId int32) (*[]dto.ListsResponse, error) {
	return s.repo.GetAllLists(userId)
}

func (s *ListService) CreateList(list *entity.List) (*entity.List, error) {
	return s.repo.CreateList(list)
}

func (s *ListService) UpdateList(list *entity.List, listId, userId int32) (*dto.ListsResponse, error) {
	return s.repo.UpdateList(list, listId, userId)
}

func (s *ListService) DeleteList(listId, userId int32) error {
	return s.repo.DeleteList(listId, userId)
}

func (s *ListService) AddGoodsToList(listId int32, goods map[int32]int32) error {
	return s.repo.AddGoodsToList(listId, goods)
}

func (s *ListService) AddDishToLIst(listId int32, dishes []int32) error {
	return s.repo.AddDishToList(listId, dishes)
}

func (s *ListService) GetShopping(listId, userId int32) (map[string]int32, error) {
	result := make(map[string]int32)
	list, err := s.GetList(listId, userId)
	if err != nil {
		return result, err
	}
	for _, good := range list.Goods {
		result[good.Title] += good.Amount
	}
	for _, dish := range list.Dishes {
		for _, good := range dish.Goods {
			result[good.Title] += good.Amount
		}
	}
	return result, err
}
