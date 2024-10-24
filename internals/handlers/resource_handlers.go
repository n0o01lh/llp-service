package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/n0o01lh/llp/internals/core/domain"
	"github.com/n0o01lh/llp/internals/core/ports"
	"github.com/n0o01lh/llp/internals/utils"
)

type ResourceHandlers struct {
	resourceService ports.ResourceService
}

func NewResourceHandlers(resourceService ports.ResourceService) *ResourceHandlers {

	return &ResourceHandlers{
		resourceService: resourceService,
	}
}

var _ ports.ResourceHandlers = (*ResourceHandlers)(nil)

func (h *ResourceHandlers) Create(ctx *fiber.Ctx) error {

	resource := new(domain.Resource)
	err := ctx.BodyParser(&resource)

	validationErrors := utils.Validate(resource)

	if len(validationErrors) > 0 {
		errMsgs := utils.GetErrorsMessages(validationErrors)

		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: strings.Join(errMsgs, " and "),
		}
	}

	if err != nil {
		log.Error(err)
		ctx.SendStatus(400)
		return err
	}

	resourceCreated, err := h.resourceService.Create(resource)

	if err != nil {
		log.Error(err)
		return err
	}

	log.Info("Resource created", resource)

	ctx.JSON(resourceCreated)
	ctx.Status(http.StatusOK)

	return nil
}

func (h *ResourceHandlers) ListAll(ctx *fiber.Ctx) error {
	resourceList, err := h.resourceService.ListAll()

	if err != nil {
		log.Error(err)
		ctx.SendStatus(http.StatusInternalServerError)
		return err
	}

	ctx.JSON(resourceList)
	ctx.Status(http.StatusOK)

	return nil
}

func (h *ResourceHandlers) ListAllByTeacherId(ctx *fiber.Ctx) error {
	teacherId := ctx.QueryInt("id")

	resourceList, err := h.resourceService.ListAllByTeacherId(uint(teacherId))

	if err != nil {
		log.Error(err)
		ctx.SendStatus(http.StatusInternalServerError)
		return err
	}

	ctx.JSON(resourceList)
	ctx.Status(http.StatusOK)

	return nil
}

func (h *ResourceHandlers) FindOne(ctx *fiber.Ctx) error {
	id := ctx.QueryInt("id")

	resource, err := h.resourceService.FindOne(uint(id))

	if err != nil {
		log.Error(err)
		ctx.SendStatus(http.StatusInternalServerError)
		return err
	}

	log.Info("Resource founded", resource)

	ctx.JSON(resource)
	ctx.Status(http.StatusOK)

	return nil
}

func (h *ResourceHandlers) Update(ctx *fiber.Ctx) error {

	id, err := ctx.ParamsInt("id")
	resource := new(domain.Resource)

	if err != nil {
		log.Error(err)
		ctx.SendStatus(http.StatusInternalServerError)
		return err
	}
	err = ctx.BodyParser(&resource)

	if err != nil {
		log.Error(err)
		ctx.SendStatus(http.StatusBadRequest)
		return err
	}

	result, err := h.resourceService.Update(uint(id), resource)

	if err != nil {
		log.Error(err)
		ctx.SendStatus(http.StatusInternalServerError)
		return err
	}

	log.Info("Resource Updated", result)

	ctx.JSON(result)
	ctx.Status(http.StatusOK)

	return nil
}

func (h *ResourceHandlers) Delete(ctx *fiber.Ctx) error {

	id, err := ctx.ParamsInt("id")

	if err != nil {
		log.Error(err)
		ctx.SendStatus(http.StatusInternalServerError)
		return err
	}

	err = h.resourceService.Delete(uint(id))

	if err != nil {
		log.Error(err)
		ctx.SendStatus(http.StatusInternalServerError)
		return err
	}

	ctx.SendStatus(http.StatusOK)
	return nil
}

func (h *ResourceHandlers) Search(ctx *fiber.Ctx) error {
	criteria := ctx.Query("title")

	if criteria == "" {
		return errors.New("Criteria is empty")
	}

	resources, err := h.resourceService.Search(criteria)

	if err != nil {
		log.Error(err)
		ctx.Status(http.StatusInternalServerError)
		return err
	}

	ctx.JSON(resources)
	ctx.Status(http.StatusOK)

	return nil
}

func (h *ResourceHandlers) SalesHistory(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")

	if err != nil {
		log.Error(err)
		ctx.Status(http.StatusBadRequest)
		return err
	}

	salesHistory, err := h.resourceService.SalesHistory(uint(id))

	if err != nil {
		log.Error(err)
		ctx.Status(http.StatusInternalServerError)
		return err
	}

	ctx.JSON(salesHistory)
	ctx.Status(http.StatusOK)

	return nil
}

func (h *ResourceHandlers) SalesHistoryByTeacher(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")

	if err != nil {
		log.Error(err)
		ctx.Status(http.StatusBadRequest)
		return err
	}

	salesHistory, err := h.resourceService.SalesHistoryByTeacher(uint(id))

	if err != nil {
		log.Error(err)
		ctx.Status(http.StatusInternalServerError)
		return err
	}

	ctx.JSON(salesHistory)
	ctx.Status(http.StatusOK)

	return nil
}
