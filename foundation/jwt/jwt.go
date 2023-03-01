package jwtutil

import (
	"elektron-canteen/api/data/user"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

func Generate(user user.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = user.ID.Hex()
	claims["exp"] = time.Now().Add(time.Hour * 30).Unix()

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, err
}

func DecodeIntoClaims(token string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	return claims, err
}
