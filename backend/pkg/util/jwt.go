package util

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(secret string, userID string, role string, expiration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  userID,
		"role": role,
		"exp":  time.Now().Add(expiration).Unix(),
		"iat":  time.Now().Unix(),
	})

	return token.SignedString([]byte(secret))
}

func ValidateJWT(secret string, tokenString string) (string, string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		
		return []byte(secret), nil
	})

	if err != nil {
		return "", "", err
	}

	if !token.Valid {
		return "", "", fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", fmt.Errorf("invalid token claims")
	}

	userID, ok := claims["sub"].(string)
	if !ok {
		return "", "", fmt.Errorf("invalid user id claim")
	}

	role, ok := claims["role"].(string)
	if !ok {
		return "", "", fmt.Errorf("invalid role claim")
	}

	return userID, role, nil
}