package list

// ListServiceInterface list service behaviour
type ListServiceInterface interface {
}

// ListService init structure for list service
type ListService struct {
	repo ListRepoInterface
}

// NewListService init func for list service
func NewListService(repo ListRepoInterface) *ListService {
	return &ListService{repo: repo}
}
