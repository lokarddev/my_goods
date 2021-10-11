package services

import "my_goods/internal/repos"

// ListServiceInterface list service behaviour
type ListServiceInterface interface {
}

// ListService init structure for list service
type ListService struct {
	repo repos.ListRepoInterface
}

// NewListService init func for list service
func NewListService(repo repos.ListRepoInterface) *ListService {
	return &ListService{repo: repo}
}
