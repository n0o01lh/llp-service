package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/n0o01lh/llp/internals/core/domain"
	"github.com/n0o01lh/llp/internals/core/ports"
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

	if err := ctx.BodyParser(&resource); err != nil {
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

	response, _ := json.Marshal(&resourceCreated)

	ctx.SendString(string(response))

	return nil
}

func (h *ResourceHandlers) ListAll(ctx *fiber.Ctx) error {
	resourceList, err := h.resourceService.ListAll()

	if err != nil {
		log.Error(err)
		ctx.SendStatus(http.StatusInternalServerError)
		return err
	}

	response, _ := json.Marshal(resourceList)

	ctx.Send(response)
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

	response, _ := json.Marshal(resource)

	ctx.Send(response)

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

	if err := ctx.BodyParser(&resource); err != nil {
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

	response, _ := json.Marshal(result)
	ctx.Send(response)

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
