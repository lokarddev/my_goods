package auth_service

import (
	"errors"
	"log"
	"my_goods/internal/delivery/web/auth"
	"my_goods/internal/entity/dto"
	"my_goods/internal/service"
	"time"
)

type UsersService struct {
	repo         service.UsersRepo
	tokenManager service.AuthenticationManager
}

func NewUsersService(repo service.UsersRepo) *UsersService {
	tokenManager, err := auth.NewTokenManager()
	if err != nil {
		log.Fatal(err)
	}
	return &UsersService{repo: repo, tokenManager: tokenManager}
}

func (s *UsersService) SignIn(input dto.LoginRequest) (dto.Access, error) {
	user, exists := s.repo.GetUserByName(input.Name)
	if !exists {
		return dto.Access{}, errors.New("no user with this name")
	}

	if valid := auth.CheckPassword(input.Password, user.Password); !valid {
		return dto.Access{}, errors.New("invalid password")
	}
	session, err := s.repo.CreateSession(user.Id, s.tokenManager.NewRefreshToken())
	accessToken, err := s.tokenManager.NewAccess(user.Id)

	return dto.Access{
		Access:  accessToken,
		Refresh: session.RefreshToken,
	}, err
}

func (s *UsersService) SignUp(input dto.LoginRequest) (dto.Access, error) {
	_, exists := s.repo.GetUserByName(input.Name)
	if exists {
		return dto.Access{}, errors.New("user with this name already exists")
	}

	hashedPass, err := auth.HashPassword(input.Password)
	input.Password = hashedPass

	user, err := s.repo.CreateUser(input)
	if err != nil {
		return dto.Access{}, err
	}

	session, err := s.repo.CreateSession(user.Id, s.tokenManager.NewRefreshToken())
	accessToken, err := s.tokenManager.NewAccess(user.Id)

	return dto.Access{
		Access:  accessToken,
		Refresh: session.RefreshToken,
	}, err
}

func (s *UsersService) RefreshAccess(token string) (dto.Access, error) {
	session, err := s.repo.GetSession(token)
	if err != nil {
		return dto.Access{}, err
	}

	if session.ExpiresIn < time.Now().Unix() {
		return dto.Access{}, errors.New("refresh already expired. new login required")
	}

	session, err = s.repo.CreateSession(session.UserId, s.tokenManager.NewRefreshToken())
	accessToken, err := s.tokenManager.NewAccess(session.UserId)

	return dto.Access{
		Access:  accessToken,
		Refresh: session.RefreshToken,
	}, err
}
