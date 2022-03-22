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

func (s *ListService) GetList(id int) (*entities.List, error) {
	return s.repo.GetList(id)
}

func (s *ListService) GetAllLists() (*[]entities.List, error) {
	return s.repo.GetAllLists()
}

func (s *ListService) CreateList(list *entities.List) (*entities.List, error) {
	return s.repo.CreateList(list)
}

func (s *ListService) UpdateList(list *entities.List, id int) (*entities.List, error) {
	return s.repo.UpdateList(list, id)
}

func (s *ListService) DeleteList(id int) error {
	return s.repo.DeleteList(id)
}
