package list

// Service init structure for list service
type Service struct {
	repo Repository
}

// NewListService init func for list service
func NewListService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) getGoods() {
}

func (s *Service) getAllGoods() {
}

func (s *Service) createGoods() {
}

func (s *Service) updateGoods() {
}

func (s *Service) deleteGoods() {
}
