package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/n0o01lh/llp/internals/core/ports"
)

type ResourceCourseHandlers struct {
	resourceCourseService ports.ResourceCourseService
}

func NewResourceCourseHandlers(rcService ports.ResourceCourseService) *ResourceCourseHandlers {
	return &ResourceCourseHandlers{
		resourceCourseService: rcService,
	}
}

var _ ports.ResourceCourseHandlers = (*ResourceCourseHandlers)(nil)

func (h *ResourceCourseHandlers) AddResourceToCourse(ctx *fiber.Ctx) error {

	var requestBody map[string]interface{}
	error := json.Unmarshal(ctx.Body(), &requestBody)

	if error != nil {
		log.Error("Error parsing request body ", error)
		return error
	}

	log.Info("request body ", requestBody)
	// get the resource id
	resourceId := requestBody["resource_id"].(float64)
	// get the course id
	courseId := requestBody["course_id"].(float64)

	courseUpdated, error := h.resourceCourseService.AddResourceToCourse(uint(resourceId), uint(courseId))

	if error != nil {
		log.Error(error)
		return error
	}

	response, _ := json.Marshal(courseUpdated)

	ctx.Response().SetBody(response)
	ctx.Response().SetStatusCode(http.StatusOK)

	return nil

}

func (h *ResourceCourseHandlers) AsignCourseToResources(ctx *fiber.Ctx) error {
	return nil
}
