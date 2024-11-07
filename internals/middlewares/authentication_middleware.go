package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt"
	"github.com/n0o01lh/llp/internals/core/domain"
	"github.com/n0o01lh/llp/internals/utils"
)

func AuthenticationMiddleware(ctx *fiber.Ctx) error {

	bearertoken := strings.Split(ctx.Get("Authorization"), " ")

	if len(bearertoken) <= 1 {
		return ctx.Status(fiber.StatusUnauthorized).SendString("Invalid authorization token")
	}

	token := bearertoken[1]

	parsedToken, err := utils.ParseJWT(token)

	if err != nil || !parsedToken.Valid {
		log.Error(err)
		return ctx.Status(fiber.StatusUnauthorized).SendString("Invalid authorization token")
	}

	claims := parsedToken.Claims.(jwt.MapClaims)

	//User from token
	user := &domain.User{
		Id:     uint(claims["id"].(float64)),
		Email:  claims["email"].(string),
		RoleId: int(claims["role"].(float64)),
	}

	ctx.Locals("user", user)

	return ctx.Next()
}
