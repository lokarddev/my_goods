package dish

// Service init structure for dish service
type Service struct {
	repo Repository
}

// NewDishService init func for dish service
func NewDishService(repo Repository) *Service {
	return &Service{repo: repo}
}
