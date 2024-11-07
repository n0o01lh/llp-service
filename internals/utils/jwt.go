package utils

import (
	"os"

	"github.com/gofiber/fiber/v2/log"
	jwt "github.com/golang-jwt/jwt"
	"github.com/n0o01lh/llp/internals/core/domain"
)

func GenerateJWT(user *domain.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.Id,
		"email": user.Email,
		"role":  user.RoleId,
	})

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	if err != nil {
		log.Error(err)
		return "", err
	}

	return signedToken, nil
}

func ParseJWT(token string) (*jwt.Token, error) {

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	return parsedToken, err
}
