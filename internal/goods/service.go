package goods

// Service init structure for goods service
type Service struct {
	repo Repository
}

// NewGoodsService init func for goods service
func NewGoodsService(repo Repository) *Service {
	return &Service{repo: repo}
}
