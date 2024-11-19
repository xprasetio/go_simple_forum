package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(id int64, username, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 12).Unix(),
	})
	key := []byte(secretKey)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(tokenString, secretKey string) (int64, string, error) {
	key := []byte(secretKey)
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return 0, "", err
	}
	if !token.Valid {
		return 0, "", errors.New("invalid token")
	}
	return int64(claims["id"].(float64)), claims["username"].(string), nil
}