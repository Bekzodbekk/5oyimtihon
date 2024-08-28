package token

import (
	"fmt"
	"time"
	"user-service/internal/pkg/load"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(conf load.Config, id string) (string, error) {
	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(conf.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(conf load.Config, tokenString string) (string, error) {
	claims := &jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return conf.SecretKey, nil
	})
	if err != nil {
		return "", err
	}

	if token.Valid {
		id := (*claims)["id"].(string)
		return id, nil
	}

	return "", fmt.Errorf("Invalid token")
}
