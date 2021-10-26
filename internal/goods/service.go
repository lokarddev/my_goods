package goods

// GoodsServiceInterface goods service behaviour
type GoodsServiceInterface interface {
}

// GoodsService init structure for goods service
type GoodsService struct {
	repo GoodsRepoInterface
}

// NewGoodsService init func for goods service
func NewGoodsService(repo GoodsRepoInterface) *GoodsService {
	return &GoodsService{repo: repo}
}
