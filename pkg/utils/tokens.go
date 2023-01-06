package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"userId"`
}

func CreateToken(userId int, accessToken string, tokenLifetime time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenLifetime).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userId,
	})

	return token.SignedString([]byte(accessToken))
}

func ParseToken(inputToken string, accessToken string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(inputToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(accessToken), nil
	})

	if err != nil {
		return &TokenClaims{}, err
	}

	claims, ok := token.Claims.(*TokenClaims)

	if !ok {
		return &TokenClaims{}, errors.New("token claims are not of type")
	}

	return claims, nil
}
