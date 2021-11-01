package list

import "my_goods/internal/entity"

// Service init structure for list service
type Service struct {
	repo Repository
}

// NewListService init func for list service
func NewListService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) getList(id int) *entity.List {
	return s.repo.getList(id)
}

func (s *Service) getAllLists() *[]entity.List {
	return s.repo.getAllLists()
}

func (s *Service) createList(list *entity.List) *entity.List {
	return s.repo.createList(list)
}

func (s *Service) updateList(list *entity.List) *entity.List {
	return s.repo.updateList(list)
}

func (s *Service) deleteList(id int) {
	s.repo.deleteList(id)
}
