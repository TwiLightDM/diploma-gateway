package services

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

type JWTService struct {
	Key string
}

func NewJWTService(key string) *JWTService {
	return &JWTService{
		Key: key,
	}
}

func (s *JWTService) ParseJWT(tokenString string) (map[string]any, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.Key), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", UnexpectedSigningMethodError)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token claims")
	}

	if exp, ok := claims["exp"].(float64); ok {
		if time.Now().Unix() > int64(exp) {
			return nil, fmt.Errorf("token expired: %w", LifetimeIsOverError)
		}
	}

	data := make(map[string]any)
	for key, value := range claims {
		data[key] = value
	}

	return data, nil
}
