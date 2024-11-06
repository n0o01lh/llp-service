package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/n0o01lh/llp/internals/core/domain"
	"github.com/n0o01lh/llp/internals/core/ports"
)

type UserHandlers struct {
	UserService ports.UserService
}

func NewUserHandlers(userService ports.UserService) *UserHandlers {
	return &UserHandlers{UserService: userService}
}

var _ ports.UserHandlers = (*UserHandlers)(nil)

func (h *UserHandlers) Register(ctx *fiber.Ctx) error {

	user := new(domain.RegisterRequest)
	err := ctx.BodyParser(&user)

	if err != nil {
		log.Error(err)
		ctx.Status(http.StatusBadRequest)
		return err
	}

	userRegistered, err := h.UserService.Register(user)

	if err != nil {
		log.Error(err)
		ctx.Status(http.StatusInternalServerError)
		return err
	}

	ctx.Status(http.StatusCreated)
	ctx.JSON(userRegistered)

	return nil
}

func (h *UserHandlers) Login(ctx *fiber.Ctx) error {

	loginRequest := new(domain.LoginRequest)
	err := ctx.BodyParser(&loginRequest)

	if err != nil {
		log.Error(err)
		ctx.Status(http.StatusBadRequest)
		return err
	}

	userLogged, err := h.UserService.Login(loginRequest)

	if err != nil {
		log.Error(err)
		ctx.Status(http.StatusInternalServerError)
		ctx.JSON(&domain.ErrorResponse{
			Error: err.Error(),
		})
		return nil
	}

	ctx.Status(http.StatusOK)
	ctx.JSON(userLogged)

	return nil
}
