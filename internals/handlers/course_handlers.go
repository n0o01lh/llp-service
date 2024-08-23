package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/n0o01lh/llp/internals/core/domain"
	"github.com/n0o01lh/llp/internals/core/ports"
)

type CourseHandlers struct {
	courseService ports.CourseService
}

func NewsCourseHandlers(courseService ports.CourseService) *CourseHandlers {
	return &CourseHandlers{
		courseService: courseService,
	}
}

var _ ports.CourseHandlers = (*CourseHandlers)(nil)

func (h *CourseHandlers) Create(ctx *fiber.Ctx) error {

	course := new(domain.Course)

	if err := ctx.BodyParser(&course); err != nil {
		log.Error(err)
		ctx.SendStatus(400)
		return err
	}

	courseCreated, err := h.courseService.Create(course)

	if err != nil {
		log.Error(err)
		return err
	}

	log.Info("course created", course)

	response, _ := json.Marshal(&courseCreated)

	ctx.SendString(string(response))

	return nil
}

func (h *CourseHandlers) ListAll(ctx *fiber.Ctx) error {
	courseList, err := h.courseService.ListAll()

	if err != nil {
		log.Error(err)
		ctx.SendStatus(http.StatusInternalServerError)
		return err
	}

	response, _ := json.Marshal(courseList)

	ctx.Send(response)
	return nil
}

func (h *CourseHandlers) FindOne(ctx *fiber.Ctx) error {
	id := ctx.QueryInt("id")

	course, err := h.courseService.FindOne(uint(id))

	if err != nil {
		log.Error(err)
		ctx.SendStatus(http.StatusInternalServerError)
		return err
	}

	log.Info("course founded", course)

	response, _ := json.Marshal(course)

	ctx.Send(response)

	return nil
}

func (h *CourseHandlers) Update(ctx *fiber.Ctx) error {

	id, err := ctx.ParamsInt("id")
	course := new(domain.Course)

	if err != nil {
		log.Error(err)
		ctx.SendStatus(http.StatusInternalServerError)
		return err
	}

	if err := ctx.BodyParser(&course); err != nil {
		log.Error(err)
		ctx.SendStatus(http.StatusBadRequest)
		return err
	}

	result, err := h.courseService.Update(uint(id), course)

	if err != nil {
		log.Error(err)
		ctx.SendStatus(http.StatusInternalServerError)
		return err
	}

	log.Info("course Updated", result)

	response, _ := json.Marshal(result)
	ctx.Send(response)

	return nil
}

func (h *CourseHandlers) Delete(ctx *fiber.Ctx) error {

	id, err := ctx.ParamsInt("id")

	if err != nil {
		log.Error(err)
		ctx.SendStatus(http.StatusInternalServerError)
		return err
	}

	err = h.courseService.Delete(uint(id))

	if err != nil {
		log.Error(err)
		ctx.SendStatus(http.StatusInternalServerError)
		return err
	}

	ctx.SendStatus(http.StatusOK)
	return nil
}