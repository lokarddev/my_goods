package auth

import (
	"errors"
	"gorm.io/gorm"
	"my_goods/internal/entity"
)

type RepoAuth interface {
	CreateUser(input Auth) (int, error)
	GetUser(login, pass string) (int, error)
}

type Repository struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetUser(login, pass string) (int, error) {
	var user entity.User
	r.db.Where("login = ? AND pass = ?", login, pass).Find(&user)
	if user.ID == 0 {
		return 0, errors.New("no user with this credentials")
	}
	return int(user.ID), nil
}

func (r *Repository) CreateUser(input Auth) (int, error) {
	user := &entity.User{Login: input.Login, Pass: input.Pass}
	r.db.Create(&user)
	if user.ID == 0 {
		return 0, errors.New("something goes wrong")
	}
	return int(user.ID), nil
}
