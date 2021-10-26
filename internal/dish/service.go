package dish

// DishServiceInterface dish service behaviour
type DishServiceInterface interface {
}

// DishService init structure for dish service
type DishService struct {
	repo DishRepoInterface
}

// NewDishService init func for dish service
func NewDishService(repo DishRepoInterface) *DishService {
	return &DishService{repo: repo}
}
