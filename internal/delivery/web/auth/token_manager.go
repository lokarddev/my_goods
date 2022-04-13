package auth

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"my_goods/pkg/env"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenManager struct {
	signingKey string
	ttl        time.Duration
}

func NewTokenManager() (*TokenManager, error) {
	signingKey := env.JWTSign
	ttl := time.Duration(env.JWTExpiration) * time.Minute
	if signingKey == "" {
		return nil, errors.New("empty signing key")
	}
	return &TokenManager{signingKey: signingKey, ttl: ttl}, nil
}

func (m *TokenManager) NewAccess(userId int32) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(m.ttl).Unix(),
		Subject:   string(userId),
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

func (m *TokenManager) NewRefreshToken() string {
	u, _ := uuid.NewUUID()
	return u.String()
}
