package lists_service

import (
	"my_goods/internal/entity"
	"my_goods/internal/service"
)

type ListService struct {
	repo service.ListRepo
}

func NewListService(repo service.ListRepo) *ListService {
	return &ListService{repo: repo}
}

func (s *ListService) GetList(id int32) (*entity.ListsResponse, error) {
	return s.repo.GetList(id)
}

func (s *ListService) GetAllLists() (*[]entity.ListsResponse, error) {
	return s.repo.GetAllLists()
}

func (s *ListService) CreateList(list *entity.List) (*entity.List, error) {
	return s.repo.CreateList(list)
}

func (s *ListService) UpdateList(list *entity.List, id int32) (*entity.ListsResponse, error) {
	return s.repo.UpdateList(list, id)
}

func (s *ListService) DeleteList(id int32) error {
	return s.repo.DeleteList(id)
}

func (s *ListService) AddGoodsToList(listId int32, goods map[int32]int32) error {
	return s.repo.AddGoodsToList(listId, goods)
}

func (s *ListService) AddDishToLIst(listId int32, dishes []int32) error {
	return s.repo.AddDishToList(listId, dishes)
}
