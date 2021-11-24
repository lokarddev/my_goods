package auth

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"my_goods/pkg/environ"
	"my_goods/pkg/logger"
	"time"
)

var (
	signingKey = "someverysecretkey"
)

type ServeAuth interface {
	ParseToken(token string) (int, error)
	CreateUser(input Auth) (int, error)
	GenerateToken(input Auth) (string, error)
}

type Service struct {
	repo RepoAuth
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repo RepoAuth) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(input Auth) (int, error) {
	hash := randHash(input.Pass)
	input.Pass = hash
	return s.repo.CreateUser(input)
}

func (s *Service) ParseToken(token string) (int, error) {
	access, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		logger.Error(err)
		return 0, err
	}
	claims, ok := access.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims invalid")
	}
	return claims.UserId, nil
}

func (s *Service) GenerateToken(input Auth) (string, error) {
	login, pass := input.Login, input.Pass
	userId, err := s.repo.GetUser(login, randHash(pass))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + int64(time.Duration(environ.Ttl)),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: userId,
	})
	return token.SignedString([]byte(signingKey))
}

func randHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(environ.Salt)))
}
