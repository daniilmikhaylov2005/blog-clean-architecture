package utils 

import (
  "time"
  "github.com/golang-jwt/jwt"
)


type TokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"userId"`
}

func CreateToken(userId int, accessToken string ,tokenLifetime time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenLifetime).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userId,
	})

	return token.SignedString([]byte(accessToken))
}
