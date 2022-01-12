package service

import (
	"my_goods/internal/entities"
	"my_goods/internal/repository"
)

type ListServiceInterface interface {
	GetList(id int) *entities.List
	GetAllLists() *[]entities.List
	CreateList(list *entities.List) *entities.List
	UpdateList(list *entities.List) *entities.List
	DeleteList(id int)
}

type ListService struct {
	repo repository.ListRepo
}

func NewListService(repo repository.ListRepo) *ListService {
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
