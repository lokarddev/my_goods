package lists_service

import (
	"my_goods/internal/entities"
	"my_goods/internal/service"
)

type ListService struct {
	repo service.ListRepo
}

func NewListService(repo service.ListRepo) *ListService {
	return &ListService{repo: repo}
}

func (s *ListService) GetList(id int) *entities.List {
	return s.repo.GetList(id)
}

func (s *ListService) GetAllLists() *[]entities.List {
	return s.repo.GetAllLists()
}

func (s *ListService) CreateList(list *entities.List) *entities.List {
	return s.repo.CreateList(list)
}

func (s *ListService) UpdateList(list *entities.List) *entities.List {
	return s.repo.UpdateList(list)
}

func (s *ListService) DeleteList(id int) {
	s.repo.DeleteList(id)
}
