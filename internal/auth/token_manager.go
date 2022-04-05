package auth

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"my_goods/pkg/env"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// AuthenticationManager provides logic for JWT & Refresh tokens generation and parsing.
type AuthenticationManager interface {
	NewAccess(userId string, ttl time.Duration) (string, error)
	Parse(accessToken string) (string, error)
	NewRefreshToken() (string, error)
}

type TokenManager struct {
	signingKey string
}

func NewTokenManager() (*TokenManager, error) {
	signingKey := env.JWTSign
	if signingKey == "" {
		return nil, errors.New("empty signing key")
	}
	return &TokenManager{signingKey: signingKey}, nil
}

func (m *TokenManager) NewAccess(userId string, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(ttl).Unix(),
		Subject:   userId,
	})

	return token.SignedString([]byte(m.signingKey))
}

func (m *TokenManager) Parse(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(m.signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("error get user claims from token")
	}

	return claims["sub"].(string), nil
}

func (m *TokenManager) NewRefreshToken() (string, error) {
	u, _ := uuid.NewUUID()
	return u.String(), nil
}
