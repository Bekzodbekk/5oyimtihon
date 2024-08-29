package token

import (
	"fmt"
	"time"
	"user-service/internal/pkg/load"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(conf load.Config, id string, email string) (string, error) {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Expiration time in Unix format
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(conf.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(conf load.Config, tokenString string) (string, string, error) {
	claims := &jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(conf.SecretKey), nil
	})
	if err != nil {
		return "", "", err
	}

	if token.Valid {
		id, ok := (*claims)["id"].(string)
		if !ok {
			return "", "", fmt.Errorf("Invalid token claims for ID")
		}

		email, ok := (*claims)["email"].(string)
		if !ok {
			return "", "", fmt.Errorf("Invalid token claims for Email")
		}

		return id, email, nil
	}

	return "", "", fmt.Errorf("Invalid token")
}
