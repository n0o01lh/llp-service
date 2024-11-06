package ports

import (
	"github.com/gofiber/fiber/v2"
	"github.com/n0o01lh/llp/internals/core/domain"
)

type UserService interface {
	Register(user *domain.RegisterRequest) (*domain.RegisterResponse, error)
	Login(loginRequest *domain.LoginRequest) (*domain.LoginResponse, error)
}

type UserRepository interface {
	Register(user *domain.RegisterRequest) (*domain.RegisterResponse, error)
	Login(loginRequest *domain.LoginRequest) (*domain.User, error)
}

type UserHandlers interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
}
