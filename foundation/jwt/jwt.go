package jwt

import (
	"elektron-canteen/api/data/user"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

func GenerateJWT(user user.User) (string, error) {
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
