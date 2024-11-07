package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/n0o01lh/llp/internals/core/domain"
	"github.com/n0o01lh/llp/internals/shared/constants"
)

func AuthorizationTeacherMiddleware(ctx *fiber.Ctx) error {

	user := ctx.Locals("user").(*domain.User)

	if user.RoleId != constants.ROLE_TEACHER && user.RoleId != constants.ROLE_ADMIN {
		return ctx.Status(fiber.StatusForbidden).SendString("Permission denied")
	}

	return ctx.Next()
}
