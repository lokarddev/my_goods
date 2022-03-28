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

func (s *ListService) GetList(id int) (*entity.List, error) {
	return s.repo.GetList(id)
}

func (s *ListService) GetAllLists() (*[]entity.List, error) {
	return s.repo.GetAllLists()
}

func (s *ListService) CreateList(list *entity.List) (*entity.List, error) {
	return s.repo.CreateList(list)
}

func (s *ListService) UpdateList(list *entity.List, id int) (*entity.List, error) {
	return s.repo.UpdateList(list, id)
}

func (s *ListService) DeleteList(id int) error {
	return s.repo.DeleteList(id)
}
